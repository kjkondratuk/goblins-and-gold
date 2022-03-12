package model

type PlayerAction interface {
	Do(actor Player, action func(actee Interactable) (InteractionResult, error))
}
