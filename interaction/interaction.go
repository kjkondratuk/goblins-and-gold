package interaction

import "context"

type Type string

func (t *Type) Describe() string {
	return string(*t)
}

type Func func(c *context.Context) Result

type Result struct {
}
