package item

import "github.com/kjkondratuk/goblins-and-gold/model/attack"

type EquipInfo struct {
	Slot    string           `yaml:"slot"`
	ACBonus *int             `yaml:"ac_bonus,omitempty"`
	Attacks attack.AttackSet `yaml:"attacks,omitempty"`
}
