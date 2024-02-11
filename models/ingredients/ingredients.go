package ingredients

import (
	"restaurantAPI/models/constants"
)

type Ingredient interface {
	String() string
	IsRawIngredient(ingredientType constants.IngredientType) bool
	IsTransformed(t constants.TransformType) bool
	GetRawIngredients() []constants.IngredientType
	GetListOfTransformations() []constants.TransformType
}

var _ Ingredient = Raw{}
var _ Ingredient = Transformed{}
var _ Ingredient = Combined{}
var _ Ingredient = Final{}
