package leaderboard

import (
	"github.com/google/uuid"
)

type LeaderboardController interface {
	Creator
	Provider
	Remove(leaderboardId uuid.UUID) error
	Reset(leaderboardId uuid.UUID) error
}

type Creator interface {
	Create(name string, capacity int, mode int) (*Leaderboard, error)
}

type Provider interface {
	List(limit int, offset int) ([]*Leaderboard, error)
	ListByMode(mode int, limit int, offset int) ([]*Leaderboard, error)
	Get(leaderboardId uuid.UUID) (*Leaderboard, error)
}
