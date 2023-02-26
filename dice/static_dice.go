package dice

type StaticRoll struct {
	R int
	B bool
}

type staticDice struct {
	_values []StaticRoll
	_curr   int
}

func NewStaticDice() *staticDice {
	return &staticDice{
		_values: make([]StaticRoll, 0),
		_curr:   0,
	}
}

func NewStaticDiceWithValues(values []StaticRoll) Dice {
	return &staticDice{
		_values: values,
		_curr:   0,
	}
}

func (d *staticDice) AddRolls(i ...StaticRoll) *staticDice {
	d._values = append(d._values, i...)
	return d
}

func (d *staticDice) Roll(rollExp string) (int, bool) {
	r := d._values[d._curr].R
	b := d._values[d._curr].B
	d._curr++
	return r, b
}
