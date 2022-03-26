package player

import (
	stats2 "github.com/kjkondratuk/goblins-and-gold/src/stats"
)

type player struct {
	_baseStats    stats2.BaseStats
	_derivedStats stats2.DerivedStats
	_visionType   VisionType
	_spd          int

	// Hitpoints, in all their incarnations
	_hp     int
	_currHp int
	_tempHp int
	//_proficiencies []Proficiency
	//_equipment     []Equipment
	//_inventory     []Item
}

type Player interface {
	Hp() int
	BaseStats() stats2.BaseStats
	Damage(dmg int) bool
	Heal(heal int) bool
}

type VisionType int

const (
	MAX_LEVEL = 20

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

type Option func(player) player

func NewPlayer(opts ...Option) Player {
	base := stats2.NewBaseStats(
		stats2.WithLvl(1),
		stats2.WithStr(10),
		stats2.WithDex(10),
		stats2.WithCon(10),
		stats2.WithInt(10),
		stats2.WithInt(10),
		stats2.WithCha(10),
	)
	p := player{
		_baseStats:    base,
		_derivedStats: stats2.NewDerivedStats(base),
		_visionType:   0,
		//_lvl:          1,
		_spd:    30,
		_hp:     1,
		_currHp: 1,
		_tempHp: 0,
		//_proficiencies: nil,
		//_equipment:     nil,
		//_inventory:     nil,
	}

	for _, opt := range opts {
		p = opt(p)
	}

	return &p
}

func WithHp(hp int) Option {
	return func(p player) player {
		p._hp = hp
		p._currHp = hp
		return p
	}
}

func WithBaseStats(stats stats2.BaseStats) Option {
	return func(p player) player {
		p._baseStats = stats
		return p
	}
}

func (p *player) Heal(heal int) bool {
	p._currHp += heal

	if p._currHp > p._hp {
		p._currHp = p._hp
		return false
	}

	return true
}

func (p *player) Damage(dmg int) bool {
	p._currHp -= dmg

	if p._currHp < 0 {
		p._currHp = 0
	}

	return p._currHp > 0
}

func (p *player) Hp() int {
	return p._currHp
}

func (p *player) BaseStats() stats2.BaseStats {
	return p._baseStats
}
