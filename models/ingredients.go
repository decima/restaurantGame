package models

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"restaurantAPI/models/utils"
	"strings"
)

type Ingredient interface {
	String() string
	GetSlice() []interface{}
	DepthString(i int) string
}

var _ Ingredient = RawIngredient{}
var _ Ingredient = TransformedIngredient{}
var _ Ingredient = CombinedIngredients{}
var _ Ingredient = ConcreteIngredient{}

type RawIngredient struct {
	Name string
}

func (r RawIngredient) DepthString(i int) string {

	return "\n" + strings.Repeat(" ", i) + r.String()
}

type TransformedIngredient struct {
	Ingredient Ingredient
	Transform  utils.TransformType
}

func (t TransformedIngredient) DepthString(i int) string {
	return "\n" + strings.Repeat(" ", i) + string(t.Transform) + t.Ingredient.DepthString(i+4)
}

type CombinedIngredients struct {
	Ingredients []Ingredient
}

func (c CombinedIngredients) DepthString(i int) string {
	var result string
	for _, ingredient := range c.Ingredients {
		result += ingredient.DepthString(i + 4)
	}
	return result
}

type ConcreteIngredient struct {
	Ingredient Ingredient
}

func (c ConcreteIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c ConcreteIngredient) DepthString(i int) string {
	return c.Ingredient.DepthString(i)
}

func (c *ConcreteIngredient) UnmarshalBSONValue(b bsontype.Type, bytes []byte) error {
	if b != bson.TypeString {
		return fmt.Errorf("invalid type %s", b)
	}
	s, _, ok := bsoncore.ReadString(bytes)
	if !ok {
		return fmt.Errorf("Go errorf yourself")
	}
	parsed := IngredientParser(s)
	c.Ingredient = parsed
	return nil
}

func (c ConcreteIngredient) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.String())
}

func (c ConcreteIngredient) String() string {
	return c.Ingredient.String()
}

func (c ConcreteIngredient) GetSlice() []interface{} {
	return c.Ingredient.GetSlice()
}

func (r RawIngredient) GetSlice() []any {
	return []any{r.Name}
}

func (c CombinedIngredients) GetSlice() []any {
	slice := []interface{}{}
	for _, i := range c.Ingredients {
		slice = append(slice, i.GetSlice())
	}
	return slice
}

func (t TransformedIngredient) String() string {
	return fmt.Sprintf("%s(%s)", string(t.Transform), t.Ingredient.String())
}

func (t TransformedIngredient) GetSlice() []interface{} {
	return append(t.Ingredient.GetSlice(), t.Transform)
}

func (r RawIngredient) String() string { return r.Name }
func (c CombinedIngredients) String() string {
	var ingredients []string
	for _, i := range c.Ingredients {
		ingredients = append(ingredients, i.String())
	}
	return fmt.Sprintf("%s", strings.Join(ingredients, ","))
}
