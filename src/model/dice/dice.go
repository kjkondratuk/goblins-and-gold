package dice

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

type Roller interface {
	Roll(exp string) (int, error)
}

type roller struct{}

var (
	DefaultRoller = &roller{}
	emptyResult   = make([]int, 0)
)

func (r *roller) Roll(exp string) ([]int, error) {
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
		// r = append(r, roll(sides))
		r[i] = roll(sides)
	}
	return r
}

func roll(sides int) int {
	return rand.Intn(sides + 1)
}
