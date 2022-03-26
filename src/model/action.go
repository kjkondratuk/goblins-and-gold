package model

import "github.com/kjkondratuk/goblins-and-gold/src/model/player"

type PlayerAction interface {
	Do(actor player.Player, action func(actee Interactable) (InteractionResult, error))
}
