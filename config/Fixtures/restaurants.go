package Fixtures

import (
	"restaurantAPI/models"
	"restaurantAPI/services"
)

func LoadRestaurants(collection *services.RestaurantsRepository) {
	(*collection).Truncate()

	email := "none@foodgame.api"

	demo := models.NewRestaurant("Demo Restaurant", &email)

	services.Container.GetRestaurantService().NewHumanCook(demo, nil)

	restaurants := []*models.Restaurant{demo}

	for _, r := range restaurants {
		(*collection).Insert(r)
	}
}
