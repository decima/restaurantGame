package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restaurantAPI/http/utils"
	"restaurantAPI/models"
	"restaurantAPI/models/constants"
	"restaurantAPI/models/ingredients"
	"restaurantAPI/models/transformers/builders"
	"restaurantAPI/services"
)

type CookingController struct {
}

func NewCookingController() *CookingController {
	return &CookingController{}
}

func (cc *CookingController) Transform(c *gin.Context) {
	repository := *services.Container.GetRestaurantsRepository()
	var t transformationDTO
	if err := c.BindJSON(&t); err != nil {
		utils.BadRequest(c)
		return
	}
	restaurant, _ := repository.Find("R1")

	var found *int
	for i, item := range restaurant.Kitchen.Inventory {
		if item.Ingredient.String() == t.Ingredient {
			found = &i

			break
		}
	}
	if found == nil {
		utils.NotFound(c, fmt.Sprintf("Ingredient %v not found", t.Ingredient))
		return
	}
	if restaurant.Kitchen.Inventory[*found].Quantity < 1 {
		utils.NotFound(c, fmt.Sprintf("No mor ingredient %v", t.Ingredient))
		return
	}
	restaurant.Kitchen.Inventory[*found].Quantity--
	newIngredient := builders.IngredientBuild(restaurant.Kitchen.Inventory[*found].Ingredient).
		Transform(constants.TransformType(t.Transformation)).
		Build()

	if restaurant.Kitchen.Inventory[*found].Quantity == 0 {
		restaurant.Kitchen.Inventory = append(restaurant.Kitchen.Inventory[:*found], restaurant.Kitchen.Inventory[*found+1:]...)
	}

	foundNew := false
	for i, item := range restaurant.Kitchen.Inventory {
		if item.Ingredient.String() == newIngredient.String() {
			foundNew = true
			restaurant.Kitchen.Inventory[i].Quantity++
			break
		}
	}

	if !foundNew {
		restaurant.Kitchen.Inventory = append(restaurant.Kitchen.Inventory, models.InventoryItem{
			Ingredient: newIngredient,
			Quantity:   1,
		})
	}
	repository.Update(restaurant)

	utils.Ok(c, transformationResponse{
		Ingredient: newIngredient,
		Quantity:   1,
		Restaurant: *restaurant,
	})

}

type transformationDTO struct {
	Ingredient     string `json:"ingredient"`
	Transformation string `json:"transformation"`
}

type transformationResponse struct {
	Ingredient ingredients.Ingredient `json:"ingredient"`
	Quantity   int                    `json:"quantity"`
	Restaurant models.Restaurant      `json:"restaurant"`
}
