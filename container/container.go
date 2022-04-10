package container

import (
	"github.com/kjkondratuk/goblins-and-gold/challenge"
	"github.com/kjkondratuk/goblins-and-gold/item"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")
)

type Type string

type Container struct {
	Type        Type                      `yaml:"type"`
	Locked      *challenge.SkillChallenge `yaml:"locked"`
	Description string                    `yaml:"description"`
	Items       []item.Item               `yaml:"items"`
}
