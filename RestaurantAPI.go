package main

import (
	"restaurantAPI/config"
	"restaurantAPI/http"
)

func main() {

	config.Load()

	/*
		Fixtures.LoadRecipes(services.Container.GetRecipesRepository())
		Fixtures.LoadRestaurants(services.Container.GetRestaurantsRepository())
	*/

	panic(http.Serve(config.HostAndPort()))

}
