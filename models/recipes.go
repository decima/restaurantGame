package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"restaurantAPI/lib/database"
)

type Recipe struct {
	ID     string             `json:"id" bson:"_id" `
	Name   string             `json:"name" bson:"name"`
	Output ConcreteIngredient `json:"output" bson:"output"`
}

func (r Recipe) SetID(id database.ID) {
	r.ID = string(id)
}

func (r Recipe) GetID() database.ID {
	return database.ID(r.ID)
}

func (r Recipe) MarshalJSON() ([]byte, error) {
	//marshal struct to json
	m := map[string]interface{}{
		"id":     r.ID,
		"name":   r.Name,
		"output": r.Output,
	}
	return json.Marshal(m)
}

func (r *Recipe) UnmarshalJSON(bytes []byte) error {
	//unmarshal json to struct
	m := map[string]interface{}{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}
	id := m["id"].(primitive.ObjectID)
	r.ID = id.Hex()
	r.Name = m["name"].(string)
	r.Output = ConcreteIngredient{Ingredient: IngredientParser(m["output"].(string))}
	return nil
}

var _ database.Entity = Recipe{}
