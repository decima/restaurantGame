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

func GetRestaurantOrAbort(c *gin.Context) *models.Restaurant {
	restaurant := GetRestaurant(c)
	if restaurant == nil {
		NotFound(c)
	}
	return restaurant
}

func SetRestaurant(c *gin.Context, restaurant *models.Restaurant) {
	c.Set("restaurant", restaurant)
}

func SetCook(c *gin.Context, cook *models.CrewMember) {
	c.Set("cook", cook)
}

func GetCook(c *gin.Context) *models.CrewMember {
	cook, ok := c.Get("cook")
	if !ok {
		return nil
	}
	return cook.(*models.CrewMember)
}

func IsOwner(c *gin.Context) bool {
	return GetCook(c).IsOwner()
}
