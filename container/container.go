package container

import (
	"context"
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/challenge"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"math"
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

		var val interaction.Result
		if result == nil {
			val = interaction.Result{}
		} else {
			val = result.(interaction.Result)
		}
		return val, err
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

	p := ctx.Value(interaction.PlayerDataKey).(actors.Player)
	// Get skill check type: c.Locked.Type
	// Get the player's modifier for the specified skill
	value, _ := p.BaseStats().ModifierByName(string(c.Locked.Type))
	// Roll the check
	roll := p.Roll("1d20")

	// if the roll total beats the DC, remove the lock
	if c.Locked.DC == math.MaxUint32 {
		return interaction.Result{
			Type:    interaction.RT_Failure,
			Message: "You struggle with the lock again, but it fails to open",
		}, nil
	} else if value+roll > c.Locked.DC {
		// Unset locked skill check to unlock the container
		c.Locked = nil
	} else {
		// You get one chance, just one.  Good luck!
		c.Locked.DC = math.MaxUint32
		return interaction.Result{
			Type:    interaction.RT_Failure,
			Message: fmt.Sprintf("You rolled a %d.  Unlock attempt failed.", roll),
		}, nil
	}

	return interaction.Result{
		Type:    interaction.RT_Success,
		Message: fmt.Sprintf("You rolled a %d!  Conatiner successfully unlocked!\n", roll),
	}, nil
}

func (c Container) Describe() string {
	return c.Description
}
