package player

import "goblins-and-gold/internal/model/stats"

type Player struct {
	BaseStats    stats.BaseStats
	DerivedStats stats.DerivedStats
	VisionType   VisionType
	Lvl          int
	Spd          int
	Hp           int
	//_proficiencies []Proficiency
	//_equipment     []Equipment
	//_inventory     []Item
}

type VisionType int

const (
	VT_STANDARD = iota
	VT_DARKVISION
	VT_ADV_DARKVISION
	VT_LIGHT_SENSITIVE

	M_STR = "Str"
	M_DEX = "Dex"
	M_CON = "Con"
	M_INT = "Int"
	M_WIS = "Wis"
	M_CHA = "Cha"
)

var (
	longToAbbrevModMap = map[string]string{
		"Strength":     "Str",
		"Dexterity":    "Dex",
		"Consitution":  "Con",
		"Intelligence": "Int",
		"Wisdom":       "Wis",
		"Charisma":     "Cha",
	}
	abbrevToLongModMap = map[string]string{
		"Str": "Strength",
		"Dex": "Dexterity",
		"Con": "Constitution",
		"Int": "Intelligence",
		"Wis": "Wisdom",
		"Cha": "Charisma",
	}
)

type Option func(Player) Player

func NewPlayer(opts ...Option) Player {
	base := stats.BaseStats{
		Str: 10,
		Dex: 10,
		Con: 10,
		Int: 10,
		Wis: 10,
		Cha: 10,
	}
	p := Player{
		BaseStats:    base,
		DerivedStats: stats.NewDerivedStats(base),
		VisionType:   0,
		Lvl:          1,
		Spd:          30,
		Hp:           1,
		//_proficiencies: nil,
		//_equipment:     nil,
		//_inventory:     nil,
	}

	for _, opt := range opts {
		p = opt(p)
	}

	return p
}

func WithHp(hp int) Option {
	return func(p Player) Player {
		p.Hp = hp
		return p
	}
}
