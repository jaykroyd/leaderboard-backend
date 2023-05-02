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

func (c *Controller) Get(leaderboardId uuid.UUID) (*Leaderboard, error) {
	model, err := c.dal.GetByPK(leaderboardId)
	if err != nil {
		return nil, err
	}

	return NewLeaderboard(model), nil
}

func (c *Controller) List(limit int, offset int) ([]*Leaderboard, error) {
	models, err := c.dal.List(limit, offset)
	if err != nil {
		return nil, err
	}

	leaderboards := make([]*Leaderboard, 0, len(models))
	for _, l := range models {
		leaderboards = append(leaderboards, NewLeaderboard(l))
	}

	return leaderboards, nil
}

func (c *Controller) ListByMode(mode int, limit int, offset int) ([]*Leaderboard, error) {
	models, err := c.dal.ListByMode(mode, limit, offset)
	if err != nil {
		return nil, err
	}

	leaderboards := make([]*Leaderboard, 0, len(models))
	for _, l := range models {
		leaderboards = append(leaderboards, NewLeaderboard(l))
	}

	return leaderboards, nil
}

func (c *Controller) Create(name string, capacity int, mode int) (*Leaderboard, error) {
	model := &leaderboard.Leaderboard{
		Name:     name,
		Mode:     mode,
		Capacity: capacity,
	}

	if err := c.dal.Create(model); err != nil {
		return nil, err
	}

	return NewLeaderboard(model), nil
}

func (c *Controller) Remove(leaderboardId uuid.UUID) error {
	model, err := c.dal.GetByPK(leaderboardId)
	if err != nil {
		return err
	}

	return c.dal.Delete(model)
}

func (c *Controller) Reset(leaderboardId uuid.UUID) error {
	return c.dal.Reset(leaderboardId)
}
