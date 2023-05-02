package participant

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Participant struct {
	ID            uuid.UUID   `sql:",type:uuid,pk""`
	ExternalID    string      `sql:"external_id""`
	Name          string      `sql:"name"`
	LeaderboardID uuid.UUID   `sql:"leaderboard_id"`
	Score         int         `sql:"score"`
	Metadata      string      `sql:"metadata"` // This is a json string of map[string]string
	CreatedAt     pg.NullTime `sql:"created_at"`
	UpdatedAt     pg.NullTime `sql:"updated_at"`
}

type RankedParticipant struct {
	tableName struct{} `sql:"participants"`
	*Participant
	Rank int `sql:"rank"`
}
