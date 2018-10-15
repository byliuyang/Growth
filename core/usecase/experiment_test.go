package usecase

import (
	"math/rand"
	"testing"
	"time"

	"Growth/core/entity"

	"github.com/stretchr/testify/require"
)

func TestCreateExperiment(t *testing.T) {
	t.Run("the owner of an experiment is correctly set", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		for i := 0; i < 10; i++ {
			id := entity.ID(rand.Int63())
			exp, err := CreateExperiment(id)
			require.NoError(t, err)
			require.Equal(t, id, exp.Owner)
		}
	})
}
