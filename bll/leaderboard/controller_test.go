package leaderboard_test

import (
	"testing"
	"time"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	leaderboardDal "github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/byyjoww/leaderboard/mocks"
	"github.com/go-pg/pg/types"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		lbDal.EXPECT().Create(gomock.Any()).Return(nil)

		service := leaderboard.NewController(lbDal)
		_, err := service.Create("test", 10, leaderboard.LeaderboardModeHighscore)
		assert.Nil(t, err)
	})
}

func TestGet(t *testing.T) {
	t.Parallel()

	var (
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			expected = &leaderboardDal.Leaderboard{
				ID: uuid.New(),
				CreatedAt: types.NullTime{
					Time: time.Now().UTC(),
				},
				UpdatedAt: types.NullTime{
					Time: time.Now().UTC(),
				},
			}
		)

		lbDal.EXPECT().GetByPK(expected.ID).Return(expected, nil)

		service := leaderboard.NewController(lbDal)
		lb, err := service.Get(expected.ID)
		assert.Nil(t, err)
		assert.Equal(t, expected, lb)
	})
}

func TestList(t *testing.T) {
	t.Parallel()

	var (
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			expected = []*leaderboardDal.Leaderboard{
				{
					ID: uuid.New(),
					CreatedAt: types.NullTime{
						Time: time.Now().UTC(),
					},
					UpdatedAt: types.NullTime{
						Time: time.Now().UTC(),
					},
				},
			}
		)

		lbDal.EXPECT().List(10, 0).Return(expected, nil)

		service := leaderboard.NewController(lbDal)
		lbs, err := service.List(10, 0)
		assert.Nil(t, err)
		assert.Equal(t, expected, lbs)
	})
}

func TestRemove(t *testing.T) {
	t.Parallel()

	var (
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			expected = &leaderboardDal.Leaderboard{
				ID: uuid.New(),
			}
		)

		lbDal.EXPECT().GetByPK(expected.ID).Return(expected, nil)
		lbDal.EXPECT().Delete(expected).Return(nil)

		service := leaderboard.NewController(lbDal)
		err := service.Remove(expected.ID)
		assert.Nil(t, err)
	})
}
