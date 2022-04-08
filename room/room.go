package room

import "github.com/kjkondratuk/goblins-and-gold/path"

type Room struct {
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Paths       []path.Path `yaml:"paths"`
}
