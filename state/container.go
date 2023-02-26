package state

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/model/challenge"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"math"
)

const (
	Chest = Type("Chest")
	Body  = Type("Body")

	InteractionTypeCancel = interaction.Type("Cancel")
	InteractionTypeOpen   = interaction.Type("Open")
	InteractionTypeLoot   = interaction.Type("Loot")
	InteractionTypeUnlock = interaction.Type("Unlock")
)

type Container struct {
	Type                  Type                      `yaml:"type"`
	Locked                *challenge.SkillChallenge `yaml:"locked"`
	Description           string                    `yaml:"description"`
	SupportedInteractions []interaction.Type        `yaml:"interactions"`
	Items                 []item.Item               `yaml:"items"`
}

// Interactions : returns a map of interaction types to interaction functions that enumerates
func (c *Container) interactions() map[interaction.Type]InteractionFunc {
	return map[interaction.Type]InteractionFunc{
		InteractionTypeCancel: c.cancel,
		InteractionTypeOpen:   c.open,
		InteractionTypeLoot:   c.loot,
		InteractionTypeUnlock: c.unlock,
	}
}

func (c *Container) Do(s State, t interaction.Type) (interaction.Result, error) {
	if action, ok := c.interactions()[t]; ok {
		return action(s)
	}
	return interaction.Result{}, errors.New(fmt.Sprintf("%s is not a valid action type for a container", t))
}

func (c *Container) removeItem(i int) []item.Item {
	return append(c.Items[:i], c.Items[i+1:]...)
}

func (c *Container) open(s State) (interaction.Result, error) {
	if c.Locked == nil {
		msg := fmt.Sprintf("%s: %s!\n", InteractionTypeOpen+"ed", c.Description)

		msg = msg + fmt.Sprint("Contents:\n")
		for _, a := range c.Items {
			msg = msg + fmt.Sprintf(" - %s\n", a.Describe())
		}

		return interaction.Result{
			Type:    interaction.RtSuccess,
			Message: msg,
		}, nil
	} else {
		return interaction.Result{
			Type:    interaction.RtFailure,
			Message: "This container needs to be unlocked first",
		}, nil
	}
}

func (c *Container) cancel(s State) (interaction.Result, error) {
	return interaction.Result{
		Type:    interaction.RtSuccess,
		Message: "Cancelled",
	}, nil
}

func (c *Container) loot(s State) (interaction.Result, error) {
	if c.Locked == nil {
		if len(c.Items) > 0 {
			d := make([]ux.Described, len(c.Items))
			for i, x := range c.Items {
				d[i] = x
			}

			resultIdx, _, err := s.Prompter().Select("Items", append([]string{"Cancel"}, ux.DescribeToList(d)...))
			if err != nil {
				return interaction.Result{}, err
			}
			if resultIdx <= 0 {
				return interaction.Result{}, nil
			}
			// TODO : getting an error here when looting the small chest and then cancelling
			it := c.Items[resultIdx-1]
			c.Items = c.removeItem(resultIdx - 1)

			return interaction.Result{
				Type:          interaction.RtSuccess,
				Message:       fmt.Sprintf("You successfully looted: %s\n", it.Description),
				AcquiredItems: []item.Item{it},
			}, nil
		} else {
			return interaction.Result{
				Type:    interaction.RtFailure,
				Message: "No items to loot.  The container is empty.",
			}, nil
		}
	} else {
		return interaction.Result{
			Type:    interaction.RtFailure,
			Message: "This container needs to be opened first",
		}, nil
	}
}

func (c *Container) unlock(s State) (interaction.Result, error) {
	if c.Locked == nil {
		return interaction.Result{
			Type:    interaction.RtFailure,
			Message: "The container is already unlocked\n",
		}, nil
	}

	// TODO : should probably validate the context before we do it.
	p := s.Player()
	// Get skill check type: c.Locked.Type
	// Get the player's modifier for the specified skill
	value, _ := p.BaseStats().ModifierByName(string(c.Locked.Type))
	// Roll the check
	roll := p.Roll("1d20")

	// if the roll total beats the DC, remove the lock
	if c.Locked.DC == math.MaxUint32 {
		return interaction.Result{
			Type:    interaction.RtFailure,
			Message: "You struggle with the lock again, but it fails to open",
		}, nil
	} else if value+roll > c.Locked.DC {
		// Unset locked skill check to unlock the container
		c.Locked = nil
	} else {
		// You get one chance, just one.  Good luck!
		c.Locked.DC = math.MaxUint32
		return interaction.Result{
			Type:    interaction.RtFailure,
			Message: fmt.Sprintf("You rolled a %d.  Unlock attempt failed.", roll),
		}, nil
	}

	return interaction.Result{
		Type:    interaction.RtSuccess,
		Message: fmt.Sprintf("You rolled a %d!  Conatiner successfully unlocked!\n", roll),
	}, nil
}

func (c Container) Describe() string {
	return c.Description
}
