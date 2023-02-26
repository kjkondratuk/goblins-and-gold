package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/model/attack"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"reflect"
	"testing"
)

func Test_combatant_AC(t *testing.T) {
	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"should return AC",
			fields{_ac: 17},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.AC(); got != tt.want {
				t.Errorf("AC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combatant_BaseStats(t *testing.T) {
	bs := stats.BaseStats{
		Lvl: 10,
		Str: 12,
		Dex: 17,
		Con: 16,
		Int: 11,
		Wis: 14,
		Cha: 13,
	}

	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	tests := []struct {
		name   string
		fields fields
		want   stats.BaseStats
	}{
		{
			"should return base stats",
			fields{_baseStats: bs},
			bs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.BaseStats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combatant_Dmg(t *testing.T) {
	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	type args struct {
		hp int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"hitpoints cannot added by passing negative hp",
			fields{_hp: 5},
			args{hp: -5},
			false,
		}, {
			"damage can be applied in excess of the hitpoint total",
			fields{_hp: 5},
			args{hp: 15},
			true,
		}, {
			"damage can be applied that exactly kills a combatant",
			fields{_hp: 5},
			args{hp: 5},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.Dmg(tt.args.hp); got != tt.want {
				t.Errorf("Dmg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combatant_Health(t *testing.T) {
	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"returns health value",
			fields{_hp: 12},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.Health(); got != tt.want {
				t.Errorf("Health() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combatant_Name(t *testing.T) {
	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"should return the combatants name",
			fields{_name: "Goblin Warlord"},
			"Goblin Warlord",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combatant_Roll(t *testing.T) {
	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	type args struct {
		rollExp string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"rolls the combatant's die",
			fields{_dice: dice.NewStaticDice().AddRolls(
				dice.StaticRoll{
					R: 5,
					B: true,
				})},
			args{rollExp: "1d4"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.Roll(tt.args.rollExp); got != tt.want {
				t.Errorf("Roll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combatant_Unconscious(t *testing.T) {
	type fields struct {
		_name      string
		_dice      dice.Dice
		_ac        int
		_hp        int
		_baseStats stats.BaseStats
		_inventory []item.Item
		_attacks   attack.AttackSet
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"should report when combatant is unconscious",
			fields{_hp: 0},
			true,
		}, {
			"should report when combatant is not unconscious",
			fields{_hp: 5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combatant{
				_name:      tt.fields._name,
				_dice:      tt.fields._dice,
				_ac:        tt.fields._ac,
				_hp:        tt.fields._hp,
				_baseStats: tt.fields._baseStats,
				_inventory: tt.fields._inventory,
				_attacks:   tt.fields._attacks,
			}
			if got := c.Unconscious(); got != tt.want {
				t.Errorf("Unconscious() = %v, want %v", got, tt.want)
			}
		})
	}
}
