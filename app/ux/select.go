package ux

import (
	"github.com/manifoldco/promptui"
)

type Described interface {
	Describe() string
}

type selector struct {
	cancel string
	label  string
}

type Select interface {
	Run(items []Described) (int, string, error)
}

func NewSelector(c string, l string) Select {
	return &selector{
		cancel: c,
		label:  l,
	}
}

func (c *selector) Run(items []Described) (int, string, error) {
	var options []string
	options = append(options, c.cancel)
	for _, i := range items {
		options = append(options, i.Describe())
	}
	p := promptui.Select{Label: c.label, Items: options}
	i, v, err := p.Run()
	if i > 0 {
		// use i-1 because we added a cancel option
		return i - 1, v, err
	}
	return 0, "", nil
}
