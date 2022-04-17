package dice

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
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
			total += d._seed.Intn(s)
		}

		return total, true
	}

	return 0, false
}
