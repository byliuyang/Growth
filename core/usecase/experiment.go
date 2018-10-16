package usecase

import (
	"Growth/core/adapter"
	"Growth/core/entity"
)

func NewExperiment(store adapter.ExperimentStore) *Experiment {
	return &Experiment{store: store}
}

type Experiment struct {
	store                 adapter.ExperimentStore
	errExperimentNotFound *adapter.ErrExperimentNotFound
	errOther              error
}

func (e *Experiment) CreateExperiment(userID entity.ID) (entity.Experiment) {
	exp := e.store.Save(entity.Experiment{Owner: userID})
	e.errOther = e.store.ErrOther()
	return exp
}

func (e *Experiment) FetchByOwner(ownerID entity.ID) []entity.Experiment {
	return e.store.FetchByOwner(ownerID)
}

func (e *Experiment) FetchByID(id entity.ID) entity.Experiment {
	exp := e.store.FetchByID(id)
	e.errExperimentNotFound = e.store.ErrNotFound()
	e.errOther = e.store.ErrOther()
	return exp
}

func (e *Experiment) ErrExperimentNotFound() *adapter.ErrExperimentNotFound {
	defer func() {
		e.errExperimentNotFound = nil
	}()
	return e.errExperimentNotFound
}

func (e *Experiment) ErrOther() error {
	defer func() {
		e.errOther = nil
	}()
	return e.errOther
}
