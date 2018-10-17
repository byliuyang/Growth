package testadapter

import (
	"runtime/debug"

	"Growth/core/adapter"
	"Growth/core/entity"
)

type FakeExperimentStore struct {
	experiments []entity.Experiment
	errNotFound *adapter.ErrExperimentNotFound
	errOther    error
}

func (store *FakeExperimentStore) Save(e entity.Experiment) entity.Experiment {
	store.init()
	e.ID = entity.ID(len(store.experiments) + 1)
	store.experiments = append(store.experiments, e)
	return e
}

func (store *FakeExperimentStore) FetchByOwner(id entity.ID) (exps []entity.Experiment) {
	for _, e := range store.experiments {
		if e.Owner == id {
			exps = append(exps, e)
		}
	}
	return
}

func (store *FakeExperimentStore) FetchByID(id entity.ID) entity.Experiment {
	for _, e := range store.experiments {
		if e.ID == id {
			return e
		}
	}
	store.errNotFound =&adapter.ErrExperimentNotFound{ID: id, Stack: string(debug.Stack())}
	return entity.Experiment{}
}

func (store *FakeExperimentStore) ErrNotFound() *adapter.ErrExperimentNotFound {
	defer func() {
		store.errNotFound = nil
	}()
	return store.errNotFound
}

func (store *FakeExperimentStore) ErrNotFoundTrace() *adapter.ErrExperimentNotFound {
	defer func() {
		store.errNotFound = nil
	}()
	return store.errNotFound
}

func (store *FakeExperimentStore) ErrOther() error {
	defer func() {
		store.errOther = nil
	}()
	return store.errOther
}

func (store *FakeExperimentStore) init() {
	if store == nil {
		*store = FakeExperimentStore{}
	}
	if store.experiments == nil {
		store.experiments = make([]entity.Experiment, 0)
	}
}
