package container

import (
	"context"
	"errors"
	"fmt"

	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/challenge"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/pterm/pterm"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")

	I_Open = interaction.Type("Open")
	I_Loot = interaction.Type("Loot")
)

type Type string

type Container struct {
	Type                  Type                      `yaml:"type"`
	Locked                *challenge.SkillChallenge `yaml:"locked"`
	Description           string                    `yaml:"description"`
	SupportedInteractions []interaction.Type        `yaml:"interactions"`
	Items                 []item.Item               `yaml:"items"`
}

// Interactions : returns a map of interaction types to interaction functions that enumerates
func (c *Container) interactions() map[interaction.Type]interaction.Func {
	return map[interaction.Type]interaction.Func{
		I_Open: c.open,
		I_Loot: c.loot,
	}
}

func (c *Container) Do(ctx context.Context, t interaction.Type) (interaction.Result, error) {
	if action, ok := c.interactions()[t]; ok {
		return action(ctx)
	}
	return interaction.Result{}, errors.New(fmt.Sprintf("%s is not a valid action type for a container", t))
}

func (c *Container) removeItem(i int) []item.Item {
	return append(c.Items[:i], c.Items[i+1:]...)
}

func (c *Container) open(ctx context.Context) (interaction.Result, error) {
	if c.Locked == nil {
		pterm.Success.Printf("%s: %s!\n", I_Open+"ed", c.Description)

		pterm.Info.Println("Contents:")
		for _, a := range c.Items {
			pterm.Info.Printf(" - %s\n", a.Describe())
		}

		return interaction.Result{}, nil
	} else {
		pterm.Error.Println("This container needs to be unlocked first")
		return interaction.Result{
			Type: interaction.RT_Failure,
		}, nil
	}
}

func (c *Container) loot(ctx context.Context) (interaction.Result, error) {
	if c.Locked == nil {
		d := make([]ux.Described, len(c.Items))
		for i, x := range c.Items {
			d[i] = x
		}

		result, err := ux.NewSelector("Cancel", "Items", func(ctx context.Context, idx int, val string, err error) (interface{}, error) {
			pterm.Success.Printf("Selected item: %s\n", val)
			it := c.Items[idx-1]
			c.Items = c.removeItem(idx - 1)

			return interaction.Result{
				Type:          interaction.RT_Success,
				AcquiredItems: []item.Item{it},
			}, nil
		}).Run(ctx, d)

		return result.(interaction.Result), err
	} else {
		pterm.Error.Println("This container needs to be opened first")
		return interaction.Result{
			Type: interaction.RT_Failure,
		}, nil
	}
}

func (c *Container) unlock(ctx context.Context) (interaction.Result, error) {
	if c.Locked == nil {
		return interaction.Result{Type: interaction.RT_Failure}, nil
	}

}

func (c Container) Describe() string {
	return c.Description
}
