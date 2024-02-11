package builders

import (
	"restaurantAPI/models/constants"
	"restaurantAPI/models/ingredients"
)

type IngredientBuilder struct {
	Ingredient ingredients.Ingredient
}

func RawIngredientBuild(name constants.IngredientType) *IngredientBuilder {
	return &IngredientBuilder{Ingredient: ingredients.Raw{Name: string(name)}}
}

func IngredientBuild(i ingredients.Ingredient) *IngredientBuilder {
	return &IngredientBuilder{Ingredient: i}
}

func IngredientsBuild(i []ingredients.Ingredient) *IngredientBuilder {

	if len(i) > 1 {
		return &IngredientBuilder{Ingredient: ingredients.Combined{Ingredients: i}}
	}
	return IngredientBuild(i[0])
}

func (ib *IngredientBuilder) Add(i ingredients.Ingredient) *IngredientBuilder {
	// if ib is a combined ingredient, add i to the list
	if c, ok := ib.Ingredient.(ingredients.Combined); ok {
		c.Ingredients = append(c.Ingredients, i)
		ib.Ingredient = c
	} else {
		// if ib is not a combined ingredient, create a new combined ingredient
		ib.Ingredient = ingredients.Combined{Ingredients: []ingredients.Ingredient{ib.Ingredient, i}}
	}
	return ib
}

func (ib *IngredientBuilder) AddRaw(name constants.IngredientType) *IngredientBuilder {
	return ib.Add(ingredients.Raw{Name: string(name)})
}

func (ib *IngredientBuilder) Transform(t constants.TransformType) *IngredientBuilder {
	ib.Ingredient = ingredients.Transformed{Ingredient: ib.Ingredient, Transform: t}
	return ib
}

func (ib *IngredientBuilder) Build() ingredients.Final {
	return ingredients.Final{Ingredient: ib.Ingredient}
}

func (ib *IngredientBuilder) Chop() *IngredientBuilder {
	return ib.Transform(constants.Chopped)
}

func (ib *IngredientBuilder) Fry() *IngredientBuilder {
	return ib.Transform(constants.Fried)
}

func (ib *IngredientBuilder) Dish() *IngredientBuilder {
	return ib.Transform(constants.Dished)
}

func (ib *IngredientBuilder) Boil() *IngredientBuilder {
	return ib.Transform(constants.Boiled)
}
