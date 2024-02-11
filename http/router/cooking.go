package router

import (
	"github.com/gin-gonic/gin"
	"restaurantAPI/http/controllers"
)

func CookingRoutes(r *gin.RouterGroup) {
	controller := controllers.NewCookingController()
	r.POST("/transform", controller.Transform)
}
