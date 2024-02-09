package builders

import (
	"restaurantAPI/models"
	"restaurantAPI/models/utils"
)

type IngredientBuilder struct {
	Ingredient models.Ingredient
}

func RawIngredientBuild(name utils.IngredientType) *IngredientBuilder {
	return &IngredientBuilder{Ingredient: models.RawIngredient{string(name)}}
}

func IngredientBuild(i models.Ingredient) *IngredientBuilder {
	return &IngredientBuilder{Ingredient: i}
}

func IngredientsBuild(i []models.Ingredient) *IngredientBuilder {

	if len(i) > 1 {
		return &IngredientBuilder{Ingredient: models.CombinedIngredients{Ingredients: i}}
	}
	return IngredientBuild(i[0])
}

func (ib *IngredientBuilder) Add(i models.Ingredient) *IngredientBuilder {
	// if ib is a combined ingredient, add i to the list
	if c, ok := ib.Ingredient.(models.CombinedIngredients); ok {
		c.Ingredients = append(c.Ingredients, i)
		ib.Ingredient = c
	} else {
		// if ib is not a combined ingredient, create a new combined ingredient
		ib.Ingredient = models.CombinedIngredients{Ingredients: []models.Ingredient{ib.Ingredient, i}}
	}
	return ib
}

func (ib *IngredientBuilder) AddRaw(name utils.IngredientType) *IngredientBuilder {
	return ib.Add(models.RawIngredient{string(name)})
}

func (ib *IngredientBuilder) Transform(t utils.TransformType) *IngredientBuilder {
	ib.Ingredient = models.TransformedIngredient{ib.Ingredient, t}
	return ib
}

func (ib *IngredientBuilder) Build() models.ConcreteIngredient {
	return models.ConcreteIngredient{ib.Ingredient}
}

func (ib *IngredientBuilder) Chop() *IngredientBuilder {
	return ib.Transform(utils.Chopped)
}

func (ib *IngredientBuilder) Fry() *IngredientBuilder {
	return ib.Transform(utils.Fried)
}

func (ib *IngredientBuilder) Dish() *IngredientBuilder {
	return ib.Transform(utils.Dished)
}

func (ib *IngredientBuilder) Boil() *IngredientBuilder {
	return ib.Transform(utils.Boiled)
}
