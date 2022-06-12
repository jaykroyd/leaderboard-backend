package player_test

import (
	"testing"

	"github.com/byyjoww/leaderboard/dal/player"
	"github.com/byyjoww/leaderboard/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Parallel()

	t.Run("test_success", func(t *testing.T) {

		var (
			leaderboardId = uuid.New()
			players       = []*player.Player{
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
			expected = []*player.Player{
				players[2],
				players[3],
				players[1],
				players[0],
				players[4],
				players[5],
			}
		)

		factory := test.GetTestDalFactory()
		// factory.SetLogger(dal.NewPgLogger(logrus.New()))
		service := factory.NewPlayerDAL()

		for _, p := range players {
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
