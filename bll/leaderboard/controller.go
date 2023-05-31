package leaderboard

import (
	"context"

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

func (c *Controller) Get(ctx context.Context, leaderboardId uuid.UUID) (*Leaderboard, error) {
	model, err := c.dal.GetByPK(ctx, leaderboardId)
	if err != nil {
		return nil, err
	}

	return NewLeaderboard(model), nil
}

func (c *Controller) List(ctx context.Context, limit int, offset int) ([]*Leaderboard, error) {
	models, err := c.dal.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	leaderboards := make([]*Leaderboard, 0, len(models))
	for _, l := range models {
		leaderboards = append(leaderboards, NewLeaderboard(l))
	}

	return leaderboards, nil
}

func (c *Controller) ListByMode(ctx context.Context, mode int, limit int, offset int) ([]*Leaderboard, error) {
	models, err := c.dal.ListByMode(ctx, mode, limit, offset)
	if err != nil {
		return nil, err
	}

	leaderboards := make([]*Leaderboard, 0, len(models))
	for _, l := range models {
		leaderboards = append(leaderboards, NewLeaderboard(l))
	}

	return leaderboards, nil
}

func (c *Controller) Create(ctx context.Context, name string, capacity int, mode int) (*Leaderboard, error) {
	model := &leaderboard.Leaderboard{
		Name:     name,
		Mode:     mode,
		Capacity: capacity,
	}

	if err := c.dal.Create(ctx, model); err != nil {
		return nil, err
	}

	return NewLeaderboard(model), nil
}

func (c *Controller) Remove(ctx context.Context, leaderboardId uuid.UUID) error {
	model, err := c.dal.GetByPK(ctx, leaderboardId)
	if err != nil {
		return err
	}

	return c.dal.Delete(ctx, model)
}

func (c *Controller) Reset(ctx context.Context, leaderboardId uuid.UUID) error {
	return c.dal.Reset(ctx, leaderboardId)
}
