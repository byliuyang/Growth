package testadapter

import (
	"fmt"

	"Growth/core/adapter"
	"Growth/core/entity"
)

// FakeEventStore is a full implementation of adapter.EventStore
type FakeEventStore struct {
	Capacity int
	events   []entity.Event
}

// Save appends the given event to a slice.
func (es *FakeEventStore) Save(event entity.Event) (entity.ID, error) {
	es.init()

	if es.idleCapacity() == 0 {
		return event.ID, &adapter.ErrOutOfCapacity{Capacity: es.Capacity}
	}
	event.ID = entity.ID(len(es.events) + 1)
	es.events = append(es.events, event)
	return event.ID, nil
}

func (es *FakeEventStore) FetchByID(id entity.ID) (entity.Event, error) {
	es.init()

	for _, e := range es.events {
		if e.ID == id {
			return e, nil
		}
	}
	return entity.Event{}, &adapter.ErrEventNotFound{ID: id}
}

func (es *FakeEventStore) Clear() {
	es.init()
	es.events = nil
}

func (es *FakeEventStore) init() {
	if es == nil {
		fmt.Println("1")
		*es = FakeEventStore{}
	}
	if es.events == nil {
		fmt.Println("2")
		es.events = make([]entity.Event, 0)
	}
}

func (es *FakeEventStore) idleCapacity() int {
	fmt.Println(es.Capacity, len(es.events))
	return es.Capacity - len(es.events)
}
