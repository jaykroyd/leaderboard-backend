package player

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Player struct {
	ID            uuid.UUID   `sql:",type:uuid,pk"`
	Name          string      `sql:"name"`
	LeaderboardID uuid.UUID   `sql:"leaderboard_id"`
	Score         int         `sql:"score"`
	CreatedAt     pg.NullTime `sql:"created_at"`
	UpdatedAt     pg.NullTime `sql:"updated_at"`
}

type RankedPlayer struct {
	tableName struct{} `sql:"players"`
	*Player
	Rank int `sql:"rank"`
}
