package room

import (
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/world/path"
)

type Room struct {
	Name        string                `yaml:"name"`
	Description string                `yaml:"description"`
	Paths       []path.Path           `yaml:"paths"`
	Containers  []container.Container `yaml:"containers"`
}
