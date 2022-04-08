package src

import (
	"github.com/kjkondratuk/goblins-and-gold/src/player"
)

type PlayerAction interface {
	Do(actor player.Player, action func(actee Interactable) (InteractionResult, error))
}
