package container

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/model/container"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"math"
)

type containerController struct {
	interactionMap map[interaction.Type]InteractionFunc
}

//go:generate mockery --name ContainerController
type ContainerController interface {
	Do(s state.State, c *container.Container, t interaction.Type) (interaction.Result, error)
}

type InteractionFunc func(s state.State, c *container.Container) (interaction.Result, error)

func NewContainerController() ContainerController {
	c := &containerController{}

	c.interactionMap = map[interaction.Type]InteractionFunc{
		interaction.InteractionTypeCancel: c.cancel,
		interaction.InteractionTypeOpen:   c.open,
		interaction.InteractionTypeLoot:   c.loot,
		interaction.InteractionTypeUnlock: c.unlock,
	}

	return c
}

func (cc *containerController) Do(s state.State, c *container.Container, t interaction.Type) (interaction.Result, error) {
	if action, ok := cc.interactionMap[t]; ok {
		return action(s, c)
	}
	return interaction.Result{}, errors.New(fmt.Sprintf("%s is not a valid action type for a container", t))
}

func (cc *containerController) open(s state.State, c *container.Container) (interaction.Result, error) {
	if c.Locked == nil {
		msg := fmt.Sprintf("%s: %s!\n", interaction.InteractionTypeOpen+"ed", c.Description)

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

func (cc *containerController) cancel(s state.State, c *container.Container) (interaction.Result, error) {
	return interaction.Result{
		Type:    interaction.RtSuccess,
		Message: "Cancelled",
	}, nil
}

func (cc *containerController) loot(s state.State, c *container.Container) (interaction.Result, error) {
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
			c.Items = c.RemoveItem(resultIdx - 1)

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

func (cc *containerController) unlock(s state.State, c *container.Container) (interaction.Result, error) {
	if c.Locked == nil {
		return interaction.Result{
			Type:    interaction.RtFailure,
			Message: "The container is already unlocked\n",
		}, nil
	}

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
		Message: fmt.Sprintf("You rolled a %d!  Container successfully unlocked!\n", roll),
	}, nil
}
