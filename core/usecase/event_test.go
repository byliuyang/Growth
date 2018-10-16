package usecase

import (
	"testing"

	"Growth/core/adapter"
	"Growth/core/adapter/testadapter"
	"Growth/core/entity"

	"github.com/stretchr/testify/require"
)

func TestPushEvent(t *testing.T) {
	eventStore := &testadapter.FakeEventStore{Capacity: 10}
	_, err := PushEvent(entity.Event{}, eventStore)
	require.NoError(t, err)
}

func TestFetchEventById(t *testing.T) {
	eventStore := &testadapter.FakeEventStore{Capacity: 10}

	t.Run("event not found when store is empty", func(t *testing.T) {
		e, err := FetchEventByID(1, eventStore)
		require.IsType(t, &adapter.ErrEventNotFound{}, err) // todo: an error leak here, should define use case level error for it
		require.Equal(t, "event:1 not found", err.Error())
		require.Equal(t, entity.Event{}, e)
	})

	t.Run("return event when id is found in store", func(t *testing.T) {
		id, err := eventStore.Save(entity.Event{})
		require.NoError(t, err)

		event, err := FetchEventByID(id, eventStore)
		require.NoError(t, err)
		require.Equal(t,
			entity.Event{
				Model: entity.Model{ID: 1},
			},
			event)
	})
}

func TestPushFetchIntegration(t *testing.T) {

	eventStore := &testadapter.FakeEventStore{Capacity: 2}

	pushed1, err := PushEvent(entity.Event{}, eventStore)
	require.NoError(t, err)
	pushed2, err := PushEvent(entity.Event{}, eventStore)
	require.NoError(t, err)

	fetched1, err := FetchEventByID(pushed1.ID, eventStore)
	require.NoError(t, err)
	fetched2, err := FetchEventByID(pushed2.ID, eventStore)
	require.NoError(t, err)

	require.Equal(t, pushed1, fetched1)
	require.Equal(t, pushed2, fetched2)
}

func TestCapacity(t *testing.T) {
	eventStore := &testadapter.FakeEventStore{Capacity: 2}
	PushEvent(entity.Event{}, eventStore)
	PushEvent(entity.Event{}, eventStore)

	_, err := PushEvent(entity.Event{}, eventStore)
	require.Equal(t, &adapter.ErrOutOfCapacity{Capacity: 2}, err)
}
