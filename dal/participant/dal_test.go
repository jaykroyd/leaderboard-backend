package participant_test

import (
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
			leaderboardId = uuid.New()
			participants  = []*participant.Participant{
				{
					Name:  "test1",
					Score: 10,
				},
				{
					Name:  "test2",
					Score: 55,
				},
				{
					Name:  "test3",
					Score: 7853,
				},
				{
					Name:  "test4",
					Score: 702,
				},
				{
					Name:  "test5",
					Score: 7,
				},
				{
					Name:  "test6",
					Score: 7,
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
			p.LeaderboardID = leaderboardId
			if err := service.Create(p); err != nil {
				t.Fatal(err)
			}
		}

		results, err := service.List(leaderboardId, 10, 0)
		if err != nil {
			t.Fatal(err)
		}

		for i, e := range expected {
			assert.Equal(t, e.ID, results[i].ID)
		}
	})
}
