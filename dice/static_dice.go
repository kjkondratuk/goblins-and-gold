package dice

type staticDice struct {
	_values []int
	_curr   int
}

func NewStaticDice() *staticDice {
	return &staticDice{
		_values: make([]int, 0),
		_curr:   0,
	}
}

func NewStaticDiceWithValues(values []int) Dice {
	return &staticDice{
		_values: values,
		_curr:   0,
	}
}

func (d *staticDice) AddRolls(i ...int) *staticDice {
	d._values = append(d._values, i...)
	return d
}

func (d *staticDice) Roll(rollExp string) (int, bool) {
	return d._values[d._curr], true
}
