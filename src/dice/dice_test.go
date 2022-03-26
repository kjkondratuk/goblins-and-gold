package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoller_Roll(t *testing.T) {
	rslt, err := DefaultRoller.Roll("1d6")
	if err != nil {
		t.Errorf("Single die roll should not fail")
	}

	assert.Equal(t, 1, len(rslt))
	t.Logf("Roll result: %d", rslt[0])
	assert.Less(t, rslt[0], 7, "Dice roll result should not exceed sides")

	rslt, e1 := DefaultRoller.Roll("26")
	assert.Error(t, e1, "Should be an error when there is no number/size delimiter.")

	rslt, e2 := DefaultRoller.Roll("d12")
	assert.Error(t, e2, "Should be an error when there are no number of dice specified.")

	rslt, e3 := DefaultRoller.Roll("12d")
	assert.Error(t, e3, "Should be an error when there are no sides specified.")

	rslt, e4 := DefaultRoller.Roll("gdpdr")
	assert.Error(t, e4, "Should be an error when there are multiple delimiters.")
}
