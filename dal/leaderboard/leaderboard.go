package leaderboard

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID   `sql:",type:uuid,pk" json:"id"`
	Name      string      `sql:"name" json:"name"`
	Mode      int         `sql:"mode" json:"mode"`
	Capacity  int64       `sql:"capacity" json:"capacity"`
	CreatedAt pg.NullTime `sql:"created_at" json:"created_at"`
	UpdatedAt pg.NullTime `sql:"updated_at" json:"updated_at"`
}
