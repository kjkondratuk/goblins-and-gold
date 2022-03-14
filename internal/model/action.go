package model

import "goblins-and-gold/internal/model/player"

type PlayerAction interface {
	Do(actor player.Player, action func(actee Interactable) (InteractionResult, error))
}
