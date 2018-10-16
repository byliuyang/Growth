package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type EventType int

type Event struct {
	Type EventType
	Data interface{}
}

type EventBus struct {
	qs map[EventType][]chan Out
}

func (bus *EventBus) init() {
	if bus.qs == nil {
		bus.qs = make(map[EventType][]chan Out)
	}
}

func (bus *EventBus) Emit(e Event) error {
	bus.init()
	for _, out := range bus.qs[e.Type] {
		go func(out chan Out) {
			out <- Out{
				Data: e.Data,
				Error: nil,
			}
		}(out)
	}
	return nil
}

type Out struct {
	Data interface{}
	Error error
}

func (bus *EventBus) On(et EventType) chan Out {
	bus.init()
	// register a queue for this listener
	outs := bus.qs[et]
	out := make(chan Out)
	bus.qs[et] = append(outs, out)
	return out
}

const (
	EventX EventType = iota
	EventY
	EventZ
)

func TestX(t *testing.T) {
	defer func() {
		recover()
	}()

	bus := &EventBus{}

	xChan := bus.On(EventX)
	xChan2 := bus.On(EventX)

	bus.Emit(Event{
		Type: EventType(EventX),
		Data: "1",
	})
	bus.Emit(Event{
		Type: EventType(EventX),
		Data: "1",
	})

	out := <-xChan
	require.Equal(t, "1", out.Data)
	require.Equal(t, nil, out.Error)

	out = <-xChan2
	require.Equal(t, "1", out.Data)
	require.Equal(t, nil, out.Error)

}
