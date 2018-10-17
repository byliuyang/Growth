package adapter

import (
	"fmt"

	"Growth/core/entity"
)

type ExperimentStore interface {
	Save(entity.Experiment) entity.Experiment
	FetchByID(id entity.ID) entity.Experiment
	FetchByOwner(ownerID entity.ID) []entity.Experiment
	ErrNotFound() *ErrExperimentNotFound
	ErrOther() error
}

type ErrExperimentNotFound struct {
	ID    entity.ID
	Stack string
}

func (e *ErrExperimentNotFound) Error() string {
	return fmt.Sprintf("experiment:%d not found", e.ID)
}
