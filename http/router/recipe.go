package router

import (
	"github.com/gin-gonic/gin"
	"restaurantAPI/http/controllers"
)

func RecipeRoutes(r *gin.RouterGroup) {
	controller := controllers.NewRecipeController()
	r.GET("/:id", controller.GetRecipe)
}
