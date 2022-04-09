package room

import (
	"github.com/kjkondratuk/goblins-and-gold/interactable"
	"github.com/kjkondratuk/goblins-and-gold/path"
)

type Room struct {
	Name         string                      `yaml:"name"`
	Description  string                      `yaml:"description"`
	Paths        []path.Path                 `yaml:"paths"`
	Iteractables []interactable.Interactable `yaml:"interactables"`
}
