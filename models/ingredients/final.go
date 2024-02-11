package ingredients

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"restaurantAPI/models/constants"
)

type Final struct {
	Ingredient Ingredient
}

func (c Final) GetListOfTransformations() []constants.TransformType {
	return c.Ingredient.GetListOfTransformations()
}

func (c Final) GetRawIngredients() []constants.IngredientType {
	return c.Ingredient.GetRawIngredients()
}

func (c Final) IsTransformed(t constants.TransformType) bool {
	return c.Ingredient.IsTransformed(t)
}

func (c Final) IsRawIngredient(ingredientType constants.IngredientType) bool {
	return c.Ingredient.IsRawIngredient(ingredientType)
}

func (c Final) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Final) UnmarshalBSONValue(b bsontype.Type, bytes []byte) error {
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

func (c Final) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.String())
}

func (c Final) String() string {
	return c.Ingredient.String()
}
