package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"restaurantAPI/lib/faker"
	"restaurantAPI/models/constants"
	"restaurantAPI/models/generators"
)

type Crew []*CrewMember

type CrewRole string

func (c CrewRole) UnmarshalBSONValue(b bsontype.Type, bytes []byte) error {
	v, _, ok := bsoncore.ReadValue(bytes, b)
	if !ok {
		return fmt.Errorf("cannot read value of type %s", b.String())
	}

	i, ok := v.StringValueOK()
	if !ok {
		return fmt.Errorf("cannot convert value to string")
	}
	switch i {
	case "owner":
		c = "Owner"
	case "cook":
		c = "Cook"
	case "bot":
		c = "Bot"
	default:
		c = ""
	}
	return nil
}

const (
	Owner string = "owner"
	Cook  string = "cook"
	Bot   string = "bot"
)

type CrewMember struct {
	ID     string            `json:"id" bson:"_id"`
	Role   string            `json:"role" bson:"role"`
	Name   string            `json:"name" bson:"name"`
	Skills []constants.Skill `json:"skills" bson:"skills"`
}

func (c *CrewMember) IsOwner() bool {
	return c.Role == Owner
}

func (c *CrewMember) IsCook() bool {
	return c.Role == Cook
}

func (c *CrewMember) IsBot() bool {
	return c.Role == Bot
}

func NewCrewMate(name string, role string) *CrewMember {
	return &CrewMember{
		Name:   name,
		Role:   string(role),
		Skills: constants.AllSkills(),
	}
}

func NewBotCrewMate(skills []constants.Skill) CrewMember {
	return CrewMember{
		Name:   faker.PersonName(),
		Role:   Bot,
		Skills: skills,
	}
}

func (c *Crew) HireMember(crewMember *CrewMember) error {

	crewMember.ID = generators.ShortID.Generate(func(probe string) bool {
		_, ok := c.GetMember(probe)
		return !ok
	}, "C_")
	*c = append(*c, crewMember)
	// print all crew members

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
