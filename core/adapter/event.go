package adapter

import (
	"fmt"

	"Growth/core/entity"
)

// EventStore is used for saving and fetching events.
type EventStore interface {
	Save(event entity.Event) (entity.ID, error)
	FetchByID(id entity.ID) (event entity.Event, err error)
}

type ErrEventNotFound struct {
	ID entity.ID
}

func (e *ErrEventNotFound) Error() string {
	return fmt.Sprintf("event:%d not found", e.ID)
}

type ErrOutOfCapacity struct {
	Capacity int
}

func (e *ErrOutOfCapacity) Error() string {
	return fmt.Sprintf("event store out of capacity, max: %d", e.Capacity)
}
