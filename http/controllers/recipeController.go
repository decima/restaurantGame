package controllers

import (
	"github.com/gin-gonic/gin"
	"restaurantAPI/lib/database"
	"restaurantAPI/services"
)

type RecipeController struct {
}

func NewRecipeController() *RecipeController {
	return &RecipeController{}
}

func (rc *RecipeController) GetRecipe(c *gin.Context) {

	repository := *services.Container.GetRecipesRepository()
	recipe, err := repository.Find(database.ID(c.Param("id")))

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if recipe == nil {

		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	c.JSON(200, recipe)
}
