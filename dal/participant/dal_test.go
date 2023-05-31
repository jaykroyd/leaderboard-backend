package participant_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/byyjoww/leaderboard/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Parallel()

	t.Run("test_success", func(t *testing.T) {

		var (
			ctx           = context.Background()
			leaderboardId = uuid.New()
			participants  = []*participant.Participant{
				{
					LeaderboardID: leaderboardId,
					ExternalID:    "1",
					Metadata:      "{}",
					Name:          "test1",
					Score:         10,
				},
				{
					LeaderboardID: leaderboardId,
					ExternalID:    "2",
					Metadata:      "{}",
					Name:          "test2",
					Score:         55,
				},
				{
					LeaderboardID: leaderboardId,
					ExternalID:    "3",
					Metadata:      "{}",
					Name:          "test3",
					Score:         7853,
				},
				{
					LeaderboardID: leaderboardId,
					ExternalID:    "4",
					Metadata:      "{}",
					Name:          "test4",
					Score:         702,
				},
				{
					LeaderboardID: leaderboardId,
					ExternalID:    "5",
					Metadata:      "{}",
					Name:          "test5",
					Score:         7,
				},
				{
					LeaderboardID: leaderboardId,
					ExternalID:    "6",
					Metadata:      "{}",
					Name:          "test6",
					Score:         7,
				},
			}
			expected = []*participant.Participant{
				participants[2],
				participants[3],
				participants[1],
				participants[0],
				participants[4],
				participants[5],
			}
		)

		factory := test.GetTestDalFactory()
		// factory.SetLogger(dal.NewPgLogger(logrus.New()))
		service := factory.NewParticipantDAL()

		for _, p := range participants {
			if err := service.Create(ctx, p); err != nil {
				t.Fatal(err)
			}
		}

		results, err := service.List(ctx, leaderboardId, 100, 0)
		assert.Nil(t, err)
		assert.Len(t, results, len(expected))
		for i, e := range expected {
			assert.Equal(t, e.ID, results[i].ID)
		}

		rankedParticipant, err := service.GetRankedByExternalID(ctx, leaderboardId, "1")
		fmt.Println(rankedParticipant.Rank)
		assert.Nil(t, err)
		assert.Equal(t, 4, rankedParticipant.Rank)

		for _, p := range participants {
			service.Delete(ctx, p)
		}

		t.Fail()
	})
}
