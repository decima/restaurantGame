package models

import (
	"errors"
	"restaurantAPI/lib/database"
	"restaurantAPI/models/constants"
	"restaurantAPI/models/ingredients"
	"time"
)

type Restaurant struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     *string   `json:"email,omitempty" bson:"email,omitempty"`
	Money     int       `json:"money" bson:"money"`
	Kitchen   Kitchen   `json:"kitchen" bson:"kitchen"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewRestaurant(name string, email *string) *Restaurant {

	return &Restaurant{
		Name:      name,
		CreatedAt: time.Now(),
		Email:     email,
		Money:     10000,
		Kitchen: Kitchen{
			EquipmentMaxSize: 4,
			CrewMaxSize:      2,
			Equipment: map[string]Equipment{
				"EQ1": {Name: constants.CuttingStation, RequiredSkills: []constants.Skill{constants.Chopping, constants.Slicing, constants.Dicing}},
				"EQ2": {Name: constants.Fryer, RequiredSkills: []constants.Skill{constants.Frying}},
			},
			Crew: Crew{},
			Inventory: Inventory{
				{Ingredient: ingredients.Final{ingredients.Raw{string(constants.Potato)}}, Quantity: 10},
				{Ingredient: ingredients.Final{ingredients.Raw{string(constants.Fish)}}, Quantity: 10},
			},
		},
	}
}

func (r *Restaurant) GetID() database.ID {
	return database.ID(r.ID)
}

func (r *Restaurant) SetID(id database.ID) {
	r.ID = string(id)
}

func (r *Restaurant) GetEntityName() database.CollectionName {
	return database.Restaurants
}

type Inventory []InventoryItem

type InventoryItem struct {
	Ingredient ingredients.Final `json:"ingredient"`
	Quantity   int               `json:"quantity" bson:"quantity"`
}

type Kitchen struct {
	EquipmentMaxSize int                  `json:"equipment_max_size" bson:"equipment_max_size"`
	CrewMaxSize      int                  `json:"crew_max_size" bson:"crew_max_size"`
	Equipment        map[string]Equipment `json:"equipment" bson:"equipment"`
	Crew             Crew                 `json:"crew" bson:"crew"`
	Inventory        Inventory            `json:"inventory" bson:"inventory"`
}

type Equipment struct {
	Name           constants.EquipmentType `json:"name" bson:"name"`
	RequiredSkills []constants.Skill       `json:"required_skills" bson:"required_skills"`
}

var _ database.Entity = (*Restaurant)(nil)

func (r *Restaurant) GetListOfEquipmentID() []string {
	var list []string
	for k, _ := range r.Kitchen.Equipment {
		list = append(list, k)
	}
	return list
}

func (r *Restaurant) GetListOfEquipmentType() []string {
	var list []string
	for _, e := range r.Kitchen.Equipment {
		list = append(list, string(e.Name))
	}
	return list
}

func (r *Restaurant) GetListOfEmployeeID() []string {
	var list []string
	for _, v := range r.Kitchen.Crew {
		list = append(list, v.ID)
	}
	return list
}

func (r *Restaurant) GetListOfEmployeeName() []string {
	var list []string
	for _, v := range r.Kitchen.Crew {
		list = append(list, v.Name)
	}
	return list
}

func (r *Restaurant) GetListOfInventory() []string {
	var list []string
	for _, i := range r.Kitchen.Inventory {
		list = append(list, i.Ingredient.String())
	}
	return list
}

func (r *Restaurant) HireEmployee(employee *CrewMember) error {
	if ok, err := r.CanHire(); !ok {
		return err
	}
	return r.Kitchen.Crew.HireMember(employee)
}

func (r *Restaurant) CanHire() (bool, error) {
	if r.Kitchen.CrewMaxSize <= len(r.Kitchen.Crew) {
		return false, ErrKitchenIsFull
	}
	return true, nil
}

var ErrKitchenIsFull = errors.New("kitchen is full")
