package container

import (
	"github.com/kjkondratuk/goblins-and-gold/model/challenge"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")
)

type Type string

func (t Type) Describe() string {
	return string(t)
}

type Container struct {
	Type                  Type                      `yaml:"type"`
	Locked                *challenge.SkillChallenge `yaml:"locked"`
	Description           string                    `yaml:"description"`
	SupportedInteractions []Type                    `yaml:"interactions"`
	Items                 []item.Item               `yaml:"items"`
}

func (c *Container) RemoveItem(i int) []item.Item {
	return append(c.Items[:i], c.Items[i+1:]...)
}

func (c Container) Describe() string {
	return c.Description
}
