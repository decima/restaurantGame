package ingredients

import (
	"fmt"
	"restaurantAPI/models/constants"
	"strings"
)

type Combined struct {
	Ingredients []Ingredient
}

func (c Combined) GetListOfTransformations() []constants.TransformType {

	var transformations []constants.TransformType
	for _, i := range c.Ingredients {
		transformations = append(transformations, i.GetListOfTransformations()...)
	}
	return transformations
}

func (c Combined) GetRawIngredients() []constants.IngredientType {
	var rawIngredients []constants.IngredientType
	for _, i := range c.Ingredients {
		rawIngredients = append(rawIngredients, i.GetRawIngredients()...)
	}
	return rawIngredients
}

func (c Combined) IsTransformed(t constants.TransformType) bool {
	for _, i := range c.Ingredients {
		if i.IsTransformed(t) {
			return true
		}
	}
	return false
}

func (c Combined) IsRawIngredient(ingredientType constants.IngredientType) bool {
	for _, i := range c.Ingredients {
		if i.IsRawIngredient(ingredientType) {
			return true
		}
	}
	return false
}

func (c Combined) String() string {
	var ingredients []string
	for _, i := range c.Ingredients {
		ingredients = append(ingredients, i.String())
	}
	return fmt.Sprintf("%s", strings.Join(ingredients, ","))
}
