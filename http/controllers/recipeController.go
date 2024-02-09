package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restaurantAPI/config"
	"restaurantAPI/lib/database"
	"restaurantAPI/models"
)

type RecipeController struct {
	collection database.Collection[models.Recipe]
}

func NewRecipeController() *RecipeController {
	collection, _ := config.Collection[models.Recipe]("recipes")
	return &RecipeController{
		collection: collection,
	}
}

func (rc RecipeController) GetRecipe(c *gin.Context) {

	recipe, err := rc.collection.Find(database.ID(c.Param("id")))

	fmt.Println(recipe.Output)
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
