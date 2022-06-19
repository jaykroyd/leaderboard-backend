package leaderboard

import (
	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/google/uuid"
)

type LeaderboardController interface {
	List(limit int, offset int) ([]*leaderboard.Leaderboard, error)
	Get(leaderboardId uuid.UUID) (*leaderboard.Leaderboard, error)
	Create() (*leaderboard.Leaderboard, error)
	Remove(leaderboardId uuid.UUID) error
	Reset(leaderboardId uuid.UUID) error
}

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

func (c *Controller) Create() (*leaderboard.Leaderboard, error) {
	lb := &leaderboard.Leaderboard{}
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
