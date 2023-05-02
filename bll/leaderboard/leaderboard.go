package leaderboard

import (
	"time"

	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Mode      int       `json:"mode"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewLeaderboard(model *leaderboard.Leaderboard) *Leaderboard {
	return &Leaderboard{
		ID:        model.ID,
		Name:      model.Name,
		Mode:      model.Mode,
		Capacity:  model.Capacity,
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: model.UpdatedAt.Time,
	}
}
