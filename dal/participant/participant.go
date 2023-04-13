package participant

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Participant struct {
	ID            uuid.UUID   `sql:",type:uuid,pk" json:"id"`
	Name          string      `sql:"name" json:"name"`
	LeaderboardID uuid.UUID   `sql:"leaderboard_id" json:"-"`
	Score         int         `sql:"score" json:"score"`
	Metadata      string      `sql:"metadata" json:"metadata"` // This is a json string of map[string]interface{}
	CreatedAt     pg.NullTime `sql:"created_at" json:"created_at"`
	UpdatedAt     pg.NullTime `sql:"updated_at" json:"updated_at"`
}

type RankedParticipant struct {
	tableName struct{} `sql:"participants"`
	*Participant
	Rank int `sql:"rank" json:"rank"`
}
