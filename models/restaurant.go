package models

import (
	"restaurantAPI/lib/database"
	"time"
)

type Restaurant struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Inventory Inventory `json:"inventory" bson:"inventory"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func (r Restaurant) GetID() database.ID {
	return database.ID(r.ID)
}

func (r Restaurant) SetID(id database.ID) {
	r.ID = string(id)
}

func (r Restaurant) GetEntityName() database.CollectionName {
	return database.Restaurants
}

type Inventory struct {
	Items []InventoryItem `json:"items" bson:"items"`
	Money int             `json:"money" bson:"money"`
}

type InventoryItem struct {
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

type Kitchen struct {
	Size      int         `json:"size" bson:"size"`
	Equipment []Equipment `json:"equipment" bson:"equipment"`
}

type Equipment struct {
	Name           EquipmentType `json:"name" bson:"name"`
	RequiredSkills []Skill       `json:"required_skills" bson:"required_skills"`
}

type EquipmentType string

const (
	Stove          EquipmentType = "stove"
	Oven           EquipmentType = "oven"
	Grill          EquipmentType = "grill"
	Sink           EquipmentType = "sink"
	Dishwasher     EquipmentType = "dishwasher"
	CuttingStation EquipmentType = "cutting_station"
	FryingStation  EquipmentType = "frying_station"
	DishingStation EquipmentType = "dishing_station"
	GarbageBin     EquipmentType = "garbage_bin"
)

var _ database.Entity = Restaurant{}
