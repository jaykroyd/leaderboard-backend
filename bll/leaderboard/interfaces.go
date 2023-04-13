package leaderboard

import (
	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/google/uuid"
)

type LeaderboardController interface {
	Creator
	Provider
	Remove(leaderboardId uuid.UUID) error
	Reset(leaderboardId uuid.UUID) error
}

type Creator interface {
	Create(name string, capacity int64, mode int) (*leaderboard.Leaderboard, error)
}

type Provider interface {
	List(limit int, offset int) ([]*leaderboard.Leaderboard, error)
	ListByMode(mode int, limit int, offset int) ([]*leaderboard.Leaderboard, error)
	Get(leaderboardId uuid.UUID) (*leaderboard.Leaderboard, error)
}
