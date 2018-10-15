package usecase

import "Growth/core/entity"

func CreateExperiment(userID entity.ID) (entity.Experiment, error) {
	return entity.Experiment{Model: entity.Model{ID: userID}}, nil
}
