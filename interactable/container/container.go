package container

import "github.com/kjkondratuk/goblins-and-gold/item"

type Container struct {
	Description string      `yaml:"description"`
	Items       []item.Item `yaml:"items"`
}
