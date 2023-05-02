package leaderboard

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID   `sql:",type:uuid,pk"`
	Name      string      `sql:"name"`
	Mode      int         `sql:"mode"`
	Capacity  int         `sql:"capacity"`
	CreatedAt pg.NullTime `sql:"created_at"`
	UpdatedAt pg.NullTime `sql:"updated_at"`
}
