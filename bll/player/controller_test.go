package player_test

import (
	"testing"
	"time"

	"github.com/byyjoww/leaderboard/bll/player"
	playerDal "github.com/byyjoww/leaderboard/dal/player"
	"github.com/byyjoww/leaderboard/mocks"
	"github.com/go-pg/pg/types"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	var (
		pDal *mocks.MockPlayerDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockPlayerDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			expected = &playerDal.Player{
				ID:            uuid.New(),
				LeaderboardID: uuid.New(),
				Score:         0,
				CreatedAt:     types.NullTime{Time: time.Now().UTC()},
				UpdatedAt:     types.NullTime{Time: time.Now().UTC()},
			}
		)

		pDal.EXPECT().GetByPK(expected.ID).Return(expected, nil)

		service := player.NewController(pDal)
		lb, err := service.Get(expected.ID)
		assert.Nil(t, err)
		assert.Equal(t, expected, lb)
	})
}

func TestUpdateScore(t *testing.T) {
	t.Parallel()

	var (
		pDal *mocks.MockPlayerDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockPlayerDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			amounts  = []int{10, 561, -1968, -1}
			expected = []int{1010, 1561, 0, 999}
		)

		for i, a := range amounts {
			p := &playerDal.Player{
				ID:            uuid.New(),
				LeaderboardID: uuid.New(),
				Score:         1000,
				CreatedAt:     types.NullTime{Time: time.Now().UTC()},
				UpdatedAt:     types.NullTime{Time: time.Now().UTC()},
			}

			pDal.EXPECT().GetByPK(p.ID).Return(p, nil)
			pDal.EXPECT().UpdateScore(p).Return(nil)

			service := player.NewController(pDal)
			after, err := service.UpdateScore(p.ID, a)
			assert.Nil(t, err)
			assert.Equal(t, expected[i], after)
		}
	})
}

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		pDal *mocks.MockPlayerDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockPlayerDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			expected = &playerDal.Player{
				LeaderboardID: uuid.New(),
				Score:         0,
			}
		)

		pDal.EXPECT().Create(gomock.Any()).Return(nil)

		service := player.NewController(pDal)
		player, err := service.Create(expected.LeaderboardID)
		assert.Nil(t, err)
		assert.Equal(t, expected, player)
	})
}

func TestRemove(t *testing.T) {
	t.Parallel()

	var (
		pDal *mocks.MockPlayerDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockPlayerDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			expected = &playerDal.Player{
				ID: uuid.New(),
			}
		)

		pDal.EXPECT().GetByPK(expected.ID).Return(expected, nil)
		pDal.EXPECT().Delete(expected).Return(nil)

		service := player.NewController(pDal)
		err := service.Remove(expected.ID)
		assert.Nil(t, err)
	})
}
