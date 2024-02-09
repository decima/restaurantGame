package Fixtures

import (
	"fmt"
	"restaurantAPI/lib/database"
	"restaurantAPI/models"
	"restaurantAPI/models/builders"
	"restaurantAPI/models/utils"
)

func LoadRecipes(collection database.Collection[models.Recipe]) {
	collection.Truncate()
	Frying(collection)
}

func Frying(collection database.Collection[models.Recipe]) {

	recipes := []models.Recipe{
		{
			ID:   "A1",
			Name: "French fries",
			Output: builders.RawIngredientBuild(utils.Potato).
				Chop().
				Fry().
				Build(),
		},
		{
			ID:   "A2",
			Name: "Fried fish",
			Output: builders.RawIngredientBuild(utils.Fish).
				Chop().
				Fry().
				Build(),
		},
		{
			ID:   "A3",
			Name: "Fish and chips",
			Output: builders.RawIngredientBuild(utils.Fish).
				Chop().
				Fry().
				Add(builders.RawIngredientBuild(utils.Potato).
					Chop().
					Fry().
					Build(),
				).
				Build(),
		},
	}

	var errors []error
	for _, recipe := range recipes {
		errors = append(errors, collection.Insert(recipe))
	}
	fmt.Println(errors)
}
