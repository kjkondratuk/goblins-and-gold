package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/model/loadout"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	t.Run("should create a new player without options", func(t *testing.T) {
		p := NewPlayer(PlayerParams{CombatantParams{
			Name:      "some player",
			AC:        0,
			HP:        0,
			BaseStats: stats.BaseStats{},
			Inventory: nil,
			Attacks:   nil,
		}, loadout.Loadout{}})

		assert.NotNil(t, p)
		assert.Equal(t, "some player", p.Name())
	})

	t.Run("should create a new player with options", func(t *testing.T) {
		d := dice.NewStaticDice().AddRolls(dice.StaticRoll{
			R: 5,
			B: false,
		})

		p := NewPlayer(PlayerParams{CombatantParams{
			Name:      "some player",
			AC:        0,
			HP:        0,
			BaseStats: stats.BaseStats{},
			Inventory: nil,
			Attacks:   nil,
		}, loadout.Loadout{}}, WithPlayerDice(d))

		assert.NotNil(t, p)
		assert.Equal(t, "some player", p.Name())

		// Make sure the dice roller was applied
		r := p.Roll("1d4")
		assert.Equal(t, 5, r)
	})
}

func Test_combatant_Acquire(t *testing.T) {
	t.Run("should acquire all items when they have no items", func(t *testing.T) {
		itemList := []item.Item{
			{
				"Armor",
				nil,
				"Shiny plate armor",
				1,
				"suit",
			},
		}
		p := NewPlayer(PlayerParams{CombatantParams{
			Inventory: nil,
		}, loadout.Loadout{}})

		p.Acquire(itemList...)

		assert.Equal(t, itemList, p.Summary().Inventory)
	})

	t.Run("should acquire items without modifying existing", func(t *testing.T) {
		acquiredItem := item.Item{
			Type:        "Armor",
			Description: "Shiny plate armor",
			Quantity:    1,
			Unit:        "suit",
		}
		existingItem := item.Item{
			Type:        "existing item",
			Description: "some description",
			Quantity:    1,
			Unit:        "unit",
		}
		itemList := []item.Item{
			acquiredItem,
		}
		p := NewPlayer(PlayerParams{CombatantParams{
			Inventory: []item.Item{
				existingItem,
			},
		}, loadout.Loadout{}})

		p.Acquire(itemList...)

		assert.Equal(t, []item.Item{existingItem, acquiredItem}, p.Summary().Inventory)
	})
}

func Test_player_Attack(t *testing.T) {
	type fields struct {
		combatant combatant
	}
	type args struct {
		c Combatant
		s AttackSelector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				combatant: tt.fields.combatant,
			}
			p.Attack(tt.args.c, tt.args.s)
		})
	}
}

func Test_player_Summary(t *testing.T) {
	type fields struct {
		combatant combatant
	}
	tests := []struct {
		name   string
		fields fields
		want   PlayerParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				combatant: tt.fields.combatant,
			}
			assert.Equalf(t, tt.want, p.Summary(), "Summary()")
		})
	}
}
