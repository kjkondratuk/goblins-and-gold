package encounter

import (
	"container/ring"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"reflect"
	"testing"
)

func TestNewCombatSequencer(t *testing.T) {
	//testPlayer := actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
	//	"some name",
	//	15,
	//	10,
	//	stats.BaseStats{},
	//	[]item.Item{},
	//	attack.AttackSet{},
	//}})
	//
	//testMonsters := []actors.Monster{
	//	actors.NewMonster(actors.MonsterParams{
	//		CombatantParams: actors.CombatantParams{
	//			Name:      "",
	//			AC:        0,
	//			HP:        0,
	//			BaseStats: stats.BaseStats{},
	//			Inventory: nil,
	//			Attacks:   nil,
	//		},
	//	}),
	//}

	type args struct {
		p actors.Player
		e Encounter
	}
	tests := []struct {
		name string
		args args
		want Sequencer
	}{
		{
			"should create an empty sequencer when inputs are nil",
			args{
				p: nil,
				e: nil,
			},
			&sequencer{
				_turnOrder: ring.New(0),
				_player:    nil,
				_fighters:  nil,
				_turn:      0,
			},
			// TODO : finish this test
		}, /*{
			name: "should create a sequencer with player and enemies",
			args: args{
				p: testPlayer,
				e: &encounter{
					_type:        "",
					_description: "",
					_enemies:     testMonsters,
				},
			},
			want: &sequencer{
				_turnOrder: ring.New(1),
				_player:    testPlayer,
				//_fighters:  testMonsters,
				_turn: 0,
			},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCombatSequencer(tt.args.p, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCombatSequencer() = %v, want %v", got, tt.want)
			}
		})
	}
}
