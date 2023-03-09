package state

import (
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/model/challenge"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")

	InteractionTypeCancel = interaction.Type("Cancel")
	InteractionTypeOpen   = interaction.Type("Open")
	InteractionTypeLoot   = interaction.Type("Loot")
	InteractionTypeUnlock = interaction.Type("Unlock")
)

type Container struct {
	Type                  Type                      `yaml:"type"`
	Locked                *challenge.SkillChallenge `yaml:"locked"`
	Description           string                    `yaml:"description"`
	SupportedInteractions []interaction.Type        `yaml:"interactions"`
	Items                 []item.Item               `yaml:"items"`
}

func (c *Container) RemoveItem(i int) []item.Item {
	return append(c.Items[:i], c.Items[i+1:]...)
}

func (c Container) Describe() string {
	return c.Description
}
