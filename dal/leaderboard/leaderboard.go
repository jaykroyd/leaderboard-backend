package leaderboard

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID   `sql:",type:uuid,pk"`
	CreatedAt pg.NullTime `sql:"created_at"`
	UpdatedAt pg.NullTime `sql:"updated_at"`
}
