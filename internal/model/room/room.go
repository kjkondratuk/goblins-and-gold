package room

import "github.com/google/uuid"

type room struct {
	_id          uuid.UUID
	_description string
}

type Option func(r room) room

func NewRoom(opts ...Option) Room {
	id := uuid.New()
	r := room{
		_id: id,
	}

	for _, opt := range opts {
		r = opt(r)
	}

	return &r
}

func WithDescription(d string) Option {
	return func(r room) room {
		r._description = d
		return r
	}
}

type Room interface {
	ID() uuid.UUID
	Description() string
}

func (r *room) ID() uuid.UUID {
	return r._id
}

func (r *room) Description() string {
	return r._description
}
