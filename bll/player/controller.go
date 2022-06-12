package player

import (
	"errors"

	"github.com/byyjoww/leaderboard/dal/player"
	"github.com/google/uuid"
)

type PlayerController interface {
	Get(playerId uuid.UUID) (*player.RankedPlayer, error)
	List(leaderboardId uuid.UUID, limit int, offset int) ([]*player.RankedPlayer, error)
	UpdateScore(playerId uuid.UUID, amount int) (int, error)
	Create(leaderboardId uuid.UUID, name string) (*player.Player, error)
	Remove(playerId uuid.UUID) error
}

type Controller struct {
	dal player.PlayerDAL
}

func NewController(dal player.PlayerDAL) *Controller {
	return &Controller{
		dal: dal,
	}
}

func (c *Controller) Get(playerId uuid.UUID) (*player.RankedPlayer, error) {
	return c.dal.GetRankedByPK(playerId)
}

func (c *Controller) List(leaderboardId uuid.UUID, limit int, offset int) ([]*player.RankedPlayer, error) {
	return c.dal.List(leaderboardId, limit, offset)
}

func (c *Controller) UpdateScore(playerId uuid.UUID, amount int) (int, error) {
	p, err := c.dal.GetByPK(playerId)
	if err != nil {
		return 0, nil
	}

	prev := p.Score
	p.Score += amount
	if p.Score < 0 {
		p.Score = 0
	}

	err = c.dal.UpdateScore(p)
	if err != nil {
		return prev, err
	}

	return p.Score, nil
}

func (c *Controller) Create(leaderboardId uuid.UUID, name string) (*player.Player, error) {
	if len(name) > 50 {
		return nil, errors.New("player name is too long")
	}

	p := &player.Player{
		LeaderboardID: leaderboardId,
		Name:          name,
		Score:         0,
	}

	return p, c.dal.Create(p)
}

func (c *Controller) Remove(playerId uuid.UUID) error {
	p, err := c.dal.GetByPK(playerId)
	if err != nil {
		return err
	}

	return c.dal.Delete(p)
}
