package container

import (
	"context"
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/challenge"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/player"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")

	InteractionTypeOpen   = interaction.Type("Open")
	InteractionTypeLoot   = interaction.Type("Loot")
	InteractionTypeUnlock = interaction.Type("Unlock")
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
		InteractionTypeOpen:   c.open,
		InteractionTypeLoot:   c.loot,
		InteractionTypeUnlock: c.unlock,
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
		s := fmt.Sprintf("%s: %s!\n", InteractionTypeOpen+"ed", c.Description)

		s = s + fmt.Sprint("Contents:\n")
		for _, a := range c.Items {
			s = s + fmt.Sprintf(" - %s\n", a.Describe())
		}

		return interaction.Result{
			Type:    interaction.RT_Success,
			Message: s,
		}, nil
	} else {
		return interaction.Result{
			Type:    interaction.RT_Failure,
			Message: "This container needs to be unlocked first",
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
			it := c.Items[idx-1]
			c.Items = c.removeItem(idx - 1)

			return interaction.Result{
				Type:          interaction.RT_Success,
				Message:       fmt.Sprintf("You successfully looted: %s\n", it.Description),
				AcquiredItems: []item.Item{it},
			}, nil
		}).Run(ctx, d)

		return result.(interaction.Result), err
	} else {
		return interaction.Result{
			Type:    interaction.RT_Failure,
			Message: "This container needs to be opened first",
		}, nil
	}
}

func (c *Container) unlock(ctx context.Context) (interaction.Result, error) {
	if c.Locked == nil {
		return interaction.Result{
			Type:    interaction.RT_Failure,
			Message: "The container is already unlocked\n",
		}, nil
	}

	p := ctx.Value(interaction.PlayerDataKey).(player.Player)
	// Get skill check type: c.Locked.Type
	//value, _ := p.BaseStats().GetByName(string(c.Locked.Type))
	// Get the player's modifier for the specified skill
	// Roll the check
	// Add together the player's bonus and the roll
	// Determine whether the check passed or failed
	p.Roll("")

	// Unset locked skill check to unlock the container
	c.Locked = nil

	return interaction.Result{
		Type:    interaction.RT_Success,
		Message: "Conatiner successfully unlocked!\n",
	}, nil
}

func (c Container) Describe() string {
	return c.Description
}
