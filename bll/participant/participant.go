package participant

import (
	"encoding/json"
	"time"

	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/google/uuid"
)

type Participant struct {
	ID            uuid.UUID         `json:"id"`
	ExternalID    string            `json:"external_id"`
	Name          string            `json:"name"`
	LeaderboardID uuid.UUID         `json:"leaderboard_id"`
	Score         int               `json:"score"`
	Metadata      map[string]string `json:"metadata"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type RankedParticipant struct {
	*Participant
	Rank int `json:"rank"`
}

func NewParticipant(model *participant.Participant) *Participant {
	metadata := map[string]string{}
	json.Unmarshal([]byte(model.Metadata), &metadata)
	return &Participant{
		ID:            model.ID,
		ExternalID:    model.ExternalID,
		Name:          model.Name,
		LeaderboardID: model.LeaderboardID,
		Score:         model.Score,
		Metadata:      metadata,
		CreatedAt:     model.CreatedAt.Time,
		UpdatedAt:     model.UpdatedAt.Time,
	}
}

func NewRankedParticipant(model *participant.RankedParticipant) *RankedParticipant {
	metadata := map[string]string{}
	json.Unmarshal([]byte(model.Metadata), &metadata)

	return &RankedParticipant{
		Participant: &Participant{
			ID:            model.ID,
			ExternalID:    model.ExternalID,
			Name:          model.Name,
			LeaderboardID: model.LeaderboardID,
			Score:         model.Score,
			Metadata:      metadata,
			CreatedAt:     model.CreatedAt.Time,
			UpdatedAt:     model.UpdatedAt.Time,
		},
		Rank: model.Rank,
	}
}
