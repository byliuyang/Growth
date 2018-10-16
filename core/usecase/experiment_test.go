package usecase

import (
	"math/rand"
	"testing"
	"time"

	"Growth/core/adapter/testadapter"
	"Growth/core/entity"

	"github.com/stretchr/testify/require"
)

func TestExperiment(t *testing.T) {
	e := NewExperiment(&testadapter.FakeExperimentStore{})

	t.Run("no data when the store is empty", func(t *testing.T) {
		e.FetchExperimentByID(entity.ID(1))
		require.Equal(t, entity.ID(1), e.ErrExperimentNotFound().ID)
	})

	t.Run("can find the same experiment which has been created", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		for i := 0; i < 10; i++ {
			userID := entity.ID(rand.Int63())
			exp := e.CreateExperiment(userID)
			require.NoError(t, e.ErrOther())
			require.Equal(t, userID, exp.Owner)

			exp2 := e.FetchExperimentByID(exp.ID)
			require.Nil(t, e.ErrExperimentNotFound())
			require.NoError(t, e.ErrOther())

			require.Equal(t, exp2.ID, exp.ID, entity.ID(i+1))
			require.Equal(t, userID, exp2.Owner)
		}
	})
}
