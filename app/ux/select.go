package ux

import (
	"context"
	"github.com/manifoldco/promptui"
)

type Described interface {
	Describe() string
}

type SelectAction func(ctx context.Context, idx int, val string, err error) error

type selector struct {
	cancel  string
	label   string
	handler SelectAction
}

type Select interface {
	Run(ctx context.Context, items []Described) error
}

func NewSelector(c string, l string, h SelectAction) Select {
	return &selector{
		cancel:  c,
		label:   l,
		handler: h,
	}
}

func (c *selector) Run(ctx context.Context, items []Described) error {
	var options []string
	options = append(options, c.cancel)
	for _, i := range items {
		options = append(options, i.Describe())
	}
	p := promptui.Select{Label: c.label, Items: options}
	i, v, err := p.Run()
	if i > 0 {
		return c.handler(ctx, i, v, err)
	}
	return nil
}
