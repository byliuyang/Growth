package resolver

import (
	"Growth/core/adapter"
	"Growth/core/entity"
	"Growth/core/usecase"
	"context"
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

type Event struct {
	event entity.Event
}

func (e *Event) Id() int32 {
	return int32(e.event.ID)
}

