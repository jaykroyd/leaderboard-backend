package leaderboard

import (
	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/google/uuid"
)

const (
	LeaderboardModeHighscore   int = 0
	LeaderboardModeIncremental int = 1
)

type Controller struct {
	dal leaderboard.LeaderboardDAL
}

func NewController(dal leaderboard.LeaderboardDAL) *Controller {
	return &Controller{
		dal: dal,
	}
}

func (c *Controller) Get(leaderboardId uuid.UUID) (*leaderboard.Leaderboard, error) {
	return c.dal.GetByPK(leaderboardId)
}

func (c *Controller) List(limit int, offset int) ([]*leaderboard.Leaderboard, error) {
	return c.dal.List(limit, offset)
}

func (c *Controller) ListByMode(mode int, limit int, offset int) ([]*leaderboard.Leaderboard, error) {
	return c.dal.ListByMode(mode, limit, offset)
}

func (c *Controller) Create(name string, capacity int64, mode int) (*leaderboard.Leaderboard, error) {
	lb := &leaderboard.Leaderboard{
		Name:     name,
		Mode:     mode,
		Capacity: capacity,
	}
	return lb, c.dal.Create(lb)
}

func (c *Controller) Remove(leaderboardId uuid.UUID) error {
	lb, err := c.Get(leaderboardId)
	if err != nil {
		return err
	}
	return c.dal.Delete(lb)
}

func (c *Controller) Reset(leaderboardId uuid.UUID) error {
	return c.dal.Reset(leaderboardId)
}
