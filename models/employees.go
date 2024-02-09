package models

import "restaurantAPI/lib/database"

type Employee struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Stats Stats  `json:"stats" bson:"stats"`
}

func (r Employee) SetID(id database.ID) {
	r.ID = string(id)
}
func (r Employee) GetID() database.ID {
	return database.ID(r.ID)
}

func (r Employee) GetEntityName() database.CollectionName {
	return database.Employees
}

type Stats struct {
	Energy    int     `json:"energy" bson:"energy"`
	EnergyMax int     `json:"energy_max" bson:"energy_max"`
	Skills    []Skill `json:"skills" bson:"skills"`
}

type Skill string

const (
	Roasting Skill = "roasting"
	Baking   Skill = "baking"
	Grilling Skill = "grilling"
	Chopping Skill = "chopping"
	Dishing  Skill = "dishing"
	Washing  Skill = "washing"
)
