package entity

import (
	"time"
)

type Experiment struct {
	Model
	Owner ID // todo: define User type
}

// ID is
type ID int64

// Model is
type Model struct {
	ID        ID
	CreatedAt time.Time
}

// EventType is
type EventType struct {
	Model
	Properties []PropertyKey
}

// Event is
type Event struct {
	Model
	Type  EventType
	ExpID ID
}

// PropertyKey is
type PropertyKey struct {
	Name string
	Type Type
}

// Type is
type Type string

// A list of acceptable data types from users.
const (
	Int    = Type("Int")
	String = Type("String")
	Bool   = Type("Bool")
)

// type User struct {
// 	Model
// }
