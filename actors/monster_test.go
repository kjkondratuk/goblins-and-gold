package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/model/attack"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMonster(t *testing.T) {
	t.Run("should create monster without options", func(t *testing.T) {
		m := NewMonster(MonsterParams{
			CombatantParams: CombatantParams{
				Name:      "",
				AC:        0,
				HP:        0,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			},
			Description: "",
		})

		assert.NotNil(t, m)
	})

	t.Run("should create monster with options", func(t *testing.T) {
		m := NewMonster(MonsterParams{
			CombatantParams: CombatantParams{
				Name:      "",
				AC:        0,
				HP:        0,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			},
			Description: "",
		}, WithMonsterDice(dice.NewStaticDice()))

		assert.NotNil(t, m)
	})
}

func Test_monster_Attack(t *testing.T) {
	t.Run("should be able to hit a player with an attack", func(t *testing.T) {
		m := NewMonster(MonsterParams{
			CombatantParams: CombatantParams{
				Name:      "goblin",
				BaseStats: stats.BaseStats{},
				Attacks: attack.AttackSet{
					"handaxe": attack.Attack{
						Damage: []attack.Damage{
							{
								"1d4",
								0,
								"slashing",
							},
						},
					},
				},
			},
			Description: "",
		}, WithMonsterDice(dice.NewStaticDice().AddRolls(
			dice.StaticRoll{R: 0, B: true},
			dice.StaticRoll{R: 11, B: true},
			dice.StaticRoll{R: 5, B: true},
		)))

		p := NewPlayer(PlayerParams{
			CombatantParams: CombatantParams{
				Name:      "player",
				AC:        10,
				HP:        10,
				BaseStats: stats.BaseStats{},
			},
		})

		m.Attack(p, RandomAttackSelector{})
		assert.Equal(t, 5, p.Health())
	})

	t.Run("should be able to miss a player with an attack", func(t *testing.T) {
		m := NewMonster(MonsterParams{
			CombatantParams: CombatantParams{
				Name:      "goblin",
				BaseStats: stats.BaseStats{},
				Attacks: attack.AttackSet{
					"handaxe": attack.Attack{
						Damage: []attack.Damage{
							{
								"1d4",
								0,
								"slashing",
							},
						},
					},
				},
			},
			Description: "",
		}, WithMonsterDice(dice.NewStaticDice().AddRolls(
			dice.StaticRoll{R: 0, B: true},
			dice.StaticRoll{R: 7, B: true},
		)))

		p := NewPlayer(PlayerParams{
			CombatantParams: CombatantParams{
				Name:      "player",
				AC:        10,
				HP:        10,
				BaseStats: stats.BaseStats{},
			},
		})

		m.Attack(p, RandomAttackSelector{})
		assert.Equal(t, 10, p.Health())
	})

	t.Run("should not deal damage if the damage roll specifier is invalid", func(t *testing.T) {
		m := NewMonster(MonsterParams{
			CombatantParams: CombatantParams{
				Name:      "goblin",
				BaseStats: stats.BaseStats{},
				Attacks: attack.AttackSet{
					"handaxe": attack.Attack{
						Damage: []attack.Damage{
							{
								"",
								0,
								"slashing",
							},
						},
					},
				},
			},
			Description: "",
		}, WithMonsterDice(dice.NewStaticDice().AddRolls(
			dice.StaticRoll{R: 0, B: true},
			dice.StaticRoll{R: 11, B: true},
			dice.StaticRoll{R: 0, B: false}),
		))

		p := NewPlayer(PlayerParams{
			CombatantParams: CombatantParams{
				Name:      "player",
				AC:        10,
				HP:        10,
				BaseStats: stats.BaseStats{},
			},
		})

		m.Attack(p, RandomAttackSelector{})
		assert.Equal(t, 10, p.Health())
	})
}

func Test_monster_Describe(t *testing.T) {
	type fields struct {
		combatant combatant
		_desc     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"should return the description of the monster",
			fields{
				combatant: combatant{},
				_desc:     "some description",
			},
			"some description",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &monster{
				combatant: tt.fields.combatant,
				_desc:     tt.fields._desc,
			}
			if got := m.Describe(); got != tt.want {
				t.Errorf("Describe() = %v, want %v", got, tt.want)
			}
		})
	}
}
