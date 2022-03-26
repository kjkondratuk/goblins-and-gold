// Package dice provides a service for rolling outcomes for either players or enemies.
package dice

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Roller : An interface to abstract the rolling of dice.  If specific rolls are needed for enemies or
// something of that nature, alternative implementations might be created.
type Roller interface {
	// Roll : Rolls a number of dice based on a dice expression of the format:
	//    <number of dice>d<number of sides>
	//
	// Returns an error when the expression is improperly formatted
	Roll(exp string) (int, error)
}

type roller struct{}

var (
	// DefaultRoller : The default roller.  Currently there is no other way to create one, so we just expose this one.
	DefaultRoller = &roller{}

	emptyResult = make([]int, 0)
	once        sync.Once
)

func (r *roller) Roll(exp string) ([]int, error) {
	// Lazily seed our RNG source
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	if strings.Contains(exp, "d") {
		pts := strings.Split(exp, "d")
		if len(pts) == 2 {
			num, err := strconv.Atoi(pts[0])
			if err != nil {
				return emptyResult, errors.New("First half of expression must be a number")
			}
			sides, err := strconv.Atoi(pts[1])
			if err != nil {
				return emptyResult, errors.New("Second half of expression must be a number")
			}
			return rollAll(num, sides), nil
		} else {
			return emptyResult, errors.New("Roll expressions must contain two numbers")
		}
	} else {
		return make([]int, 0), errors.New("Roll expressions must contain 'd'")
	}
}

func rollAll(num int, sides int) []int {
	r := make([]int, num)
	for i := 0; i < num; i++ {
		r[i] = roll(sides)
	}
	return r
}

func roll(sides int) int {
	return rand.Intn(sides + 1)
}
