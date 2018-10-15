package gql

import (
	"context"

	"Growth/core/adapter"
	"Growth/core/entity"
	"Growth/core/usecase"
)

type Mutation struct {
	EventStore adapter.EventStore
}

func (m *Mutation) NewEvent(ctx context.Context) (*Event, error) {
	event, err := usecase.PushEvent(entity.Event{}, m.EventStore)
	return &Event{
		event: event,
	}, err
}
