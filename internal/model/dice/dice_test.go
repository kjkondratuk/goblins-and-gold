package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDice_Roll(t *testing.T) {
	rslt, err := DefaultRoller.Roll("1d6")
	if err != nil {
		t.Errorf("Single die roll should not fail")
	}

	assert.Equal(t, 1, len(rslt))
	t.Logf("Roll result: %d", rslt[0])
	assert.Less(t, rslt[0], 7, "Dice roll result should not exceed sides")
}
