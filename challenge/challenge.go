package challenge

import (
	"github.com/kjkondratuk/goblins-and-gold/attack"
	"github.com/kjkondratuk/goblins-and-gold/item"
)

const (
	Strength     = "Strength"
	Dexterity    = "Dexterity"
	Constitution = "Constitution"
	Intelligence = "Intelligence"
	Wisdom       = "Wisdom"
	Charisma     = "Charisma"

	Acrobatics     = "Acrobatics"
	AnimalHandling = "Animal Handling"
	Arcana         = "Arcana"
	Athletics      = "Athletics"
	Deception      = "Deception"
	History        = "History"
	Insight        = "Insight"
	Intimidation   = "Intimidation"
	Investigation  = "Investiagation"
	Medicine       = "Medicine"
	Nature         = "Nature"
	Perception     = "Perception"
	Performance    = "Performance"
	Persuasion     = "Persuasion"
	Religion       = "Religion"
	SlightOfHand   = "Slight of Hand"
	Stealth        = "Stealth"
	Survival       = "Survival"

	PassivePerception    = "Passive Perception"
	PassiveInvestigation = "Passive Investigation"
	PassiveInsight       = "Passive Insight"
)

type SkillType string

type SkillChallenge struct {
	DC          int       `yaml:"dc"`
	Type        SkillType `yaml:"type"`
	CritSuccess Outcome   `yaml:"crit_success"`
	Success     Outcome   `yaml:"success"`
	Fail        Outcome   `yaml:"fail"`
	CritFail    Outcome   `yaml:"crit_fail"`
}

type Outcome struct {
	Items       []item.Item   `yaml:"items"`
	Damage      attack.Damage `yaml:"attack"`
	Description string        `yaml:"description"`
}
