package state

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEncounter_SingleEnemy(t *testing.T) {
	e := NewEncounter(EncounterDefinition{
		Type:        "Combat",
		Description: "A group of enemies approach",
		Enemies: []actors.MonsterParams{
			{
				CombatantParams: actors.CombatantParams{
					Name:      "combatant1",
					AC:        0,
					HP:        0,
					BaseStats: stats.BaseStats{},
					Inventory: nil,
					Attacks:   nil,
				},
			},
		},
	})

	assert.NotNil(t, e)
	assert.Equal(t, 1, len(e.Enemies()))
	assert.Equal(t, "combatant1", e.Enemies()[0].Name())
	assert.Equal(t, TypeCombat, e.(*encounter)._type)
}

func TestNewEncounter_MultiEnemy(t *testing.T) {
	e := NewEncounter(EncounterDefinition{
		Type:        "Combat",
		Description: "A group of enemies approach",
		Enemies: []actors.MonsterParams{
			{
				CombatantParams: actors.CombatantParams{
					Name:      "combatant1",
					AC:        0,
					HP:        0,
					BaseStats: stats.BaseStats{},
					Inventory: nil,
					Attacks:   nil,
				},
			}, {
				CombatantParams: actors.CombatantParams{
					Name:      "combatant2",
					AC:        0,
					HP:        0,
					BaseStats: stats.BaseStats{},
					Inventory: nil,
					Attacks:   nil,
				},
			},
		},
	})

	assert.NotNil(t, e)
	assert.Equal(t, 2, len(e.Enemies()))
	assert.Equal(t, "combatant1", e.Enemies()[0].Name())
	assert.Equal(t, "combatant2", e.Enemies()[1].Name())
	assert.Equal(t, TypeCombat, e.(*encounter)._type)
}

func TestNewEncounter_EmptyEnemy(t *testing.T) {
	e := NewEncounter(EncounterDefinition{
		Type:        "Combat",
		Description: "A group of enemies approach",
		Enemies:     []actors.MonsterParams{},
	})

	assert.NotNil(t, e)
	assert.Equal(t, 0, len(e.Enemies()))
	assert.Equal(t, TypeCombat, e.(*encounter)._type)
}

func TestNewEncounter_NoEnemy(t *testing.T) {
	e := NewEncounter(EncounterDefinition{
		Type:        "Combat",
		Description: "A group of enemies approach",
	})

	assert.NotNil(t, e)
	assert.NotNil(t, e.Enemies())
	assert.Equal(t, 0, len(e.Enemies()))
	assert.Equal(t, TypeCombat, e.(*encounter)._type)
}

//func TestEncounter_Run_OverBeforeStarted(t *testing.T) {
//	p := actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
//		Name:      "",
//		AC:        0,
//		HP:        0,
//		BaseStats: stats.BaseStats{},
//		Inventory: nil,
//		Attacks:   nil,
//	}})
//	e := NewEncounter(EncounterDefinition{
//		Type:        "Combat",
//		Description: "A group of enemies approach",
//		Enemies:     []actors.MonsterParams{},
//	})
//
//	out := e.Run(p)
//
//	assert.NotNil(t, out)
//}

// TODO : make Run() less stateful so we can apply individual actions and evaluate the outcomes
//func TestEncounter_Run_OverBeforeStarted(t *testing.T) {
//	p := actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
//		Name:      "",
//		AC:        0,
//		HP:        0,
//		BaseStats: stats.BaseStats{},
//		Inventory: nil,
//		Attacks:   nil,
//	}})
//	e := NewEncounter(EncounterDefinition{
//		Type:        "Combat",
//		Description: "A group of enemies approach",
//		Enemies:     []actors.MonsterParams{},
//	})
//
//	out := e.Run(p)
//
//	assert.NotNil(t, out)
//}
