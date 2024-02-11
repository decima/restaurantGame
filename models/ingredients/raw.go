package ingredients

import (
	"restaurantAPI/models/constants"
)

type Raw struct {
	Name string
}

func (r Raw) GetListOfTransformations() []constants.TransformType {
	return []constants.TransformType{}
}

func (r Raw) GetRawIngredients() []constants.IngredientType {
	return []constants.IngredientType{constants.IngredientType(r.Name)}
}

func (r Raw) IsTransformed(t constants.TransformType) bool {
	return false
}

func (r Raw) IsRawIngredient(ingredientType constants.IngredientType) bool {
	return r.Name == string(ingredientType)
}

func (r Raw) String() string { return r.Name }
