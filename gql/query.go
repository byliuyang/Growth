package gql

import (
	"context"

	"Growth/core/adapter"
	"Growth/core/entity"
	"Growth/core/usecase"
)

type Query struct {
	EventStore adapter.EventStore
}

type eventsArgs struct {
	ID int32
}

func (q *Query) Event(ctx context.Context, args eventsArgs) (*Event, error) {
	event, err := usecase.FetchEventByID(entity.ID(args.ID), q.EventStore)
	return &Event{
		event: event,
	}, err
}
