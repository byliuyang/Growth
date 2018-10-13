package usecase

import (
	"Growth/core/adapter"
	"Growth/core/entity"
)

// PushEvent receives an event pushed by a user and store it.
func PushEvent(event entity.Event, store adapter.EventStore) (e entity.Event, err error) {
	id, err := store.Save(event)
	if err != nil {
		return
	}
	e = event
	e.ID = id
	return
}

func FetchEventByID(id entity.ID, store adapter.EventStore) (entity.Event, error) {
	return store.FetchByID(id)
}
