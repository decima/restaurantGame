package Fixtures

import (
	"fmt"
	"restaurantAPI/models"
	"restaurantAPI/models/constants"
	"restaurantAPI/models/transformers/builders"
	"restaurantAPI/services"
)

func LoadRecipes(collection *services.RecipesRepository) {
	(*collection).Truncate()
	Frying(collection)
}

func Frying(collection *services.RecipesRepository) {

	recipes := []models.Recipe{
		{
			ID:   "A1",
			Name: "French fries",
			Output: builders.RawIngredientBuild(constants.Potato).
				Chop().
				Fry().
				Build(),
		},
		{
			ID:   "A2",
			Name: "Fried fish",
			Output: builders.RawIngredientBuild(constants.Fish).
				Chop().
				Fry().
				Build(),
		},
		{
			ID:   "A3",
			Name: "Fish and chips",
			Output: builders.RawIngredientBuild(constants.Fish).
				Chop().
				Fry().
				Add(builders.RawIngredientBuild(constants.Potato).
					Chop().
					Fry().
					Build(),
				).
				Build(),
		},
	}

	var errors []error
	for _, recipe := range recipes {
		errors = append(errors, (*collection).Insert(&recipe))
	}
	fmt.Println(errors)
}
