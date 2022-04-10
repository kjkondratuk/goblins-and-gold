package container

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/challenge"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")
)

type Type string

type Container struct {
	Type                  Type                      `yaml:"type"`
	Locked                *challenge.SkillChallenge `yaml:"locked"`
	Description           string                    `yaml:"description"`
	SupportedInteractions []interaction.Type        `yaml:"interactions"`
	Items                 []item.Item               `yaml:"items"`
}

func (c Container) Interactions() map[interaction.Type]interaction.Func {
	return map[interaction.Type]interaction.Func{
		"Open": func(c *context.Context) interaction.Result {
			return interaction.Result{}
		},
	}
}

func (c Container) Describe() string {
	return c.Description
}
