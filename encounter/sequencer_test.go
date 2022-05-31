package encounter

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"reflect"
	"testing"
)

func TestNewCombatSequencer(t *testing.T) {
	type args struct {
		p actors.Player
		e Encounter
	}
	tests := []struct {
		name string
		args args
		want Sequencer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCombatSequencer(tt.args.p, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCombatSequencer() = %v, want %v", got, tt.want)
			}
		})
	}
}
