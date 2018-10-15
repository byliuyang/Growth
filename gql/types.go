package gql

import "Growth/core/entity"

type Event struct {
	event entity.Event
}

func (e *Event) Id() int32 {
	return int32(e.event.ID)
}
