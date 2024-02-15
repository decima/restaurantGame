package models

import (
	"restaurantAPI/lib/faker"
	"restaurantAPI/models/constants"
	"restaurantAPI/models/generators"
)

type Crew []*CrewMember

type CrewMember struct {
	ID     string            `json:"id" bson:"_id"`
	Name   string            `json:"name" bson:"name"`
	Bot    bool              `json:"bot" bson:"bot"`
	Skills []constants.Skill `json:"skills" bson:"skills"`
}

func NewCrewMate(name string) *CrewMember {
	return &CrewMember{
		Name:   name,
		Bot:    false,
		Skills: constants.AllSkills(),
	}
}

func NewBotCrewMate(skills []constants.Skill) CrewMember {
	return CrewMember{
		Name:   faker.PersonName(),
		Bot:    true,
		Skills: skills,
	}
}

func (c *Crew) HireMember(crewMember *CrewMember) error {

	crewMember.ID = generators.ShortID.Generate(func(probe string) bool {
		_, ok := c.GetMember(probe)
		return !ok
	}, "C_")

	*c = append(*c, crewMember)
	return nil
}

func (c *Crew) FireMember(id string) error {
	for i, e := range *c {
		if e.ID == id {
			*c = append((*c)[:i], (*c)[i+1:]...)
			return nil
		}
	}
	return nil
}

func (c *Crew) GetMember(id string) (*CrewMember, bool) {
	for _, e := range *c {
		if e.ID == id {
			return e, true
		}
	}
	return nil, false
}
