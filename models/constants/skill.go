package constants

type Skill string

const (
	Roasting Skill = "roasting"
	Baking   Skill = "baking"
	Grilling Skill = "grilling"
	Chopping Skill = "chopping"
	Slicing  Skill = "slicing"
	Dicing   Skill = "dicing"
	Dishing  Skill = "dishing"
	Frying   Skill = "frying"

	//Washing Skill = "washing" //this skill is not used in the current version of the app
)

var allSkills = []Skill{Roasting, Baking, Grilling, Chopping, Dishing, Frying, Slicing, Dicing}

func AllSkills() []Skill {
	return allSkills
}
