package dice

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

const (
	D4   = "1d4"
	D8   = "1d8"
	D10  = "1d10"
	D12  = "1d12"
	D20  = "1d20"
	D100 = "1d100"
)

type dice struct {
	_seed *rand.Rand
}

type Dice interface {
	Roll(rollExp string) (int, bool)
}

func NewDice(seed int64) Dice {
	return &dice{
		_seed: rand.New(rand.NewSource(seed)),
	}
}

func (d *dice) Roll(rollExp string) (int, bool) {
	if match, _ := regexp.Match("^\\d+d\\d+$", []byte(rollExp)); match {
		parts := strings.Split(rollExp, "d")
		n, _ := strconv.Atoi(parts[0])
		s, _ := strconv.Atoi(parts[1])
		total := 0
		for i := 0; i < n; i++ {
			v := d._seed.Intn(s)
			total += v
		}

		return total, true
	}

	return 0, false
}
