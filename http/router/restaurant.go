package router

import (
	"github.com/gin-gonic/gin"
	"restaurantAPI/http/controllers"
	"restaurantAPI/http/middlewares"
)

func RestaurantRoutes(r *gin.RouterGroup) {
	controller := controllers.NewRestaurantController()
	r.POST("", controller.NewRestaurant)
	authenticated := r.Group("/my").Use(middlewares.AuthenticatorMiddleware())
	authenticated.GET("", controller.MyRestaurant)
	authenticated.DELETE("", controller.Close)
}
