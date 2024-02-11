package ingredients

import (
	"fmt"
	"restaurantAPI/models/constants"
)

type Transformed struct {
	Ingredient Ingredient
	Transform  constants.TransformType
}

func (t Transformed) GetListOfTransformations() []constants.TransformType {
	return append(t.Ingredient.GetListOfTransformations(), t.Transform)
}

func (t Transformed) GetRawIngredients() []constants.IngredientType {
	return t.Ingredient.GetRawIngredients()
}

func (t Transformed) IsTransformed(transformation constants.TransformType) bool {
	return t.Transform == transformation || t.Ingredient.IsTransformed(transformation)
}

func (t Transformed) IsRawIngredient(ingredientType constants.IngredientType) bool {
	return t.Ingredient.IsRawIngredient(ingredientType)
}

func (t Transformed) String() string {
	return fmt.Sprintf("%s(%s)", string(t.Transform), t.Ingredient.String())
}
