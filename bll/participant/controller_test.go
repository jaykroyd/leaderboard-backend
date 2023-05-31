package participant_test

import (
	"context"
	"testing"
	"time"

	"github.com/byyjoww/leaderboard/bll/participant"
	participantDal "github.com/byyjoww/leaderboard/dal/participant"
	"github.com/byyjoww/leaderboard/mocks"
	"github.com/go-pg/pg/types"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	var (
		pDal  *mocks.MockParticipantDAL
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockParticipantDAL(ctrl)
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			ctx      = context.Background()
			expected = &participantDal.Participant{
				ID:            uuid.New(),
				ExternalID:    uuid.NewString(),
				LeaderboardID: uuid.New(),
				Score:         0,
				CreatedAt:     types.NullTime{Time: time.Now().UTC()},
				UpdatedAt:     types.NullTime{Time: time.Now().UTC()},
			}
		)

		pDal.EXPECT().GetByPK(ctx, expected.ID).Return(expected, nil)

		service := participant.NewController(pDal, lbDal)
		lb, err := service.Get(ctx, expected.LeaderboardID, expected.ExternalID)
		assert.Nil(t, err)
		assert.Equal(t, expected, lb)
	})
}

func TestList(t *testing.T) {
	t.Parallel()

	var (
		pDal  *mocks.MockParticipantDAL
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockParticipantDAL(ctrl)
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			ctx           = context.Background()
			leaderboardId = uuid.New()
			expected      = []*participantDal.Participant{
				{
					ID:            uuid.New(),
					LeaderboardID: uuid.New(),
					Score:         0,
					CreatedAt:     types.NullTime{Time: time.Now().UTC()},
					UpdatedAt:     types.NullTime{Time: time.Now().UTC()},
				},
			}
		)

		pDal.EXPECT().List(ctx, leaderboardId, 10, 0).Return(expected, nil)

		service := participant.NewController(pDal, lbDal)
		lb, err := service.List(ctx, leaderboardId, 10, 0)
		assert.Nil(t, err)
		assert.Equal(t, expected, lb)
	})
}

func TestUpdateScore(t *testing.T) {
	t.Parallel()

	var (
		pDal  *mocks.MockParticipantDAL
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockParticipantDAL(ctrl)
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			ctx      = context.Background()
			amounts  = []int{10, 561, -1968, -1}
			expected = []int{1010, 1561, 0, 999}
		)

		for i, a := range amounts {
			p := &participantDal.Participant{
				ID:            uuid.New(),
				ExternalID:    uuid.NewString(),
				LeaderboardID: uuid.New(),
				Score:         1000,
				CreatedAt:     types.NullTime{Time: time.Now().UTC()},
				UpdatedAt:     types.NullTime{Time: time.Now().UTC()},
			}

			pDal.EXPECT().GetByPK(ctx, p.ID).Return(p, nil)
			pDal.EXPECT().UpdateScore(ctx, p).Return(nil)

			service := participant.NewController(pDal, lbDal)
			after, err := service.UpdateScore(ctx, p.LeaderboardID, p.ExternalID, a)
			assert.Nil(t, err)
			assert.Equal(t, expected[i], after)
		}
	})
}

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		pDal  *mocks.MockParticipantDAL
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockParticipantDAL(ctrl)
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			ctx      = context.Background()
			expected = &participantDal.Participant{
				LeaderboardID: uuid.New(),
				Score:         0,
			}
		)

		pDal.EXPECT().Create(ctx, gomock.Any()).Return(nil)

		service := participant.NewController(pDal, lbDal)
		participant, err := service.Create(ctx, expected.LeaderboardID, "123", "test", map[string]string{})
		assert.Nil(t, err)
		assert.Equal(t, expected, participant)
	})
}

func TestRemove(t *testing.T) {
	t.Parallel()

	var (
		pDal  *mocks.MockParticipantDAL
		lbDal *mocks.MockLeaderboardDAL
	)

	setup := func(ctrl *gomock.Controller) {
		pDal = mocks.NewMockParticipantDAL(ctrl)
		lbDal = mocks.NewMockLeaderboardDAL(ctrl)
	}

	t.Run("test_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setup(ctrl)

		var (
			ctx      = context.Background()
			expected = &participantDal.Participant{
				ID:         uuid.New(),
				ExternalID: uuid.NewString(),
			}
		)

		pDal.EXPECT().GetByPK(ctx, expected.ID).Return(expected, nil)
		pDal.EXPECT().Delete(ctx, expected).Return(nil)

		service := participant.NewController(pDal, lbDal)
		err := service.Remove(ctx, expected.LeaderboardID, expected.ExternalID)
		assert.Nil(t, err)
	})
}
