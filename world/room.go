package world

import (
	"github.com/kjkondratuk/goblins-and-gold/container"
)

type Room struct {
	Name        string                 `yaml:"name"`
	Description string                 `yaml:"description"`
	Paths       []*Path                `yaml:"paths"`
	Containers  []*container.Container `yaml:"containers"`
}
