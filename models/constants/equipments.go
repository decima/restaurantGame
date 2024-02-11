package constants

type Equipment struct {
	Name                   EquipmentType   `json:"name"`
	AvailableSkills        []Skill         `json:"available_skills"`
	AllowedTransformations []TransformType `json:"allowed_transformations"`
	Price                  int             `json:"price"`
}

var availableEquipments = map[EquipmentType]Equipment{
	FireHob: {
		Name:            FireHob,
		AvailableSkills: []Skill{Roasting, Grilling},
		Price:           1000,
	},
	Oven: {
		Name:            Oven,
		AvailableSkills: []Skill{Baking, Roasting},
		Price:           1500,
	},
	CuttingStation: {
		Name:            CuttingStation,
		AvailableSkills: []Skill{Chopping, Slicing, Dicing},
		Price:           3000,
	},
	Fryer: {
		Name:            Fryer,
		AvailableSkills: []Skill{Frying},
		Price:           2000,
	},
	Grill: {
		Name:            Grill,
		AvailableSkills: []Skill{Grilling},
		Price:           2500,
	},
	GarbageBin: {
		Name:            GarbageBin,
		AvailableSkills: []Skill{},
		Price:           100,
	},
}

func AvailableEquipments() map[EquipmentType]Equipment {
	return availableEquipments
}
