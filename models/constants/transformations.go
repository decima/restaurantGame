package constants

type Transformation struct {
	Name                 TransformType `json:"name"`
	RequiredSkills       []Skill       `json:"required_skills"`
	HumanWorkingTime     int           `json:"human_working_time"`
	AutomaticWorkingTime int           `json:"automatic_working_time"`
	AllowedExtraTime     int           `json:"allowed_extra_time"`
}

var availableTransformations = map[TransformType]Transformation{
	Chopped: {
		Name:                 Chopped,
		RequiredSkills:       []Skill{Chopping, Slicing, Dicing},
		HumanWorkingTime:     30,
		AutomaticWorkingTime: 0,
		AllowedExtraTime:     0,
	},
	Fried: {
		Name:                 Fried,
		RequiredSkills:       []Skill{Frying},
		HumanWorkingTime:     20,
		AutomaticWorkingTime: 300,
		AllowedExtraTime:     60,
	},
	Baked: {
		Name:                 Baked,
		RequiredSkills:       []Skill{Baking},
		HumanWorkingTime:     60,
		AutomaticWorkingTime: 1200,
		AllowedExtraTime:     300,
	},
	Grilled: {
		Name:                 Grilled,
		RequiredSkills:       []Skill{Grilling},
		HumanWorkingTime:     60,
		AutomaticWorkingTime: 1200,
		AllowedExtraTime:     150,
	},
	Sliced: {
		Name:                 Sliced,
		RequiredSkills:       []Skill{Slicing},
		HumanWorkingTime:     45,
		AutomaticWorkingTime: 0,
	},
	Toasted: {
		Name:                 Toasted,
		RequiredSkills:       []Skill{Roasting},
		HumanWorkingTime:     30,
		AutomaticWorkingTime: 600,
		AllowedExtraTime:     120,
	},
	Burned: {
		Name:                 Burned,
		RequiredSkills:       []Skill{},
		HumanWorkingTime:     0,
		AutomaticWorkingTime: 0,
		AllowedExtraTime:     0,
	},
}

func AvailableTransformations() map[TransformType]Transformation {
	return availableTransformations
}
