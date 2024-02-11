package utils

import (
	"github.com/gin-gonic/gin"
	"restaurantAPI/models"
)

func GetRestaurant(c *gin.Context) *models.Restaurant {
	restaurant, ok := c.Get("restaurant")
	if !ok {
		return nil
	}
	return restaurant.(*models.Restaurant)
}

func SetRestaurant(c *gin.Context, restaurant *models.Restaurant) {
	c.Set("restaurant", restaurant)
}

func SetCook(c *gin.Context, cook *models.CrewMember) {
	c.Set("cook", cook)
}
