package main

import (
	"restaurantAPI/config"
	"restaurantAPI/config/Fixtures"
	"restaurantAPI/http"
	"restaurantAPI/models"
)

func main() {

	config.Load()

	collection, _ := config.Collection[models.Recipe]("recipes")
	Fixtures.LoadRecipes(collection)

	panic(http.Serve(config.HostAndPort()))

}
