package leaderboard

import (
	"context"

	"github.com/google/uuid"
)

type LeaderboardController interface {
	Creator
	Provider
	Remove(ctx context.Context, leaderboardId uuid.UUID) error
	Reset(ctx context.Context, leaderboardId uuid.UUID) error
}

type Creator interface {
	Create(ctx context.Context, name string, capacity int, mode int) (*Leaderboard, error)
}

type Provider interface {
	List(ctx context.Context, limit int, offset int) ([]*Leaderboard, error)
	ListByMode(ctx context.Context, mode int, limit int, offset int) ([]*Leaderboard, error)
	Get(ctx context.Context, leaderboardId uuid.UUID) (*Leaderboard, error)
}
