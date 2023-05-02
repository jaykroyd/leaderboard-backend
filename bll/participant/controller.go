package participant

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/google/uuid"
)

type ParticipantController interface {
	Get(participantID uuid.UUID) (*RankedParticipant, error)
	List(leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error)
	UpdateScore(participantID uuid.UUID, amount int) (int, error)
	Create(leaderboardId uuid.UUID, externalId string, name string, metadata map[string]string) (*Participant, error)
	Remove(participantID uuid.UUID) error
}

type Controller struct {
	dal   participant.ParticipantDAL
	lbDal leaderboard.LeaderboardDAL
}

func NewController(dal participant.ParticipantDAL, lbDal leaderboard.LeaderboardDAL) *Controller {
	return &Controller{
		dal:   dal,
		lbDal: lbDal,
	}
}

func (c *Controller) Get(participantID uuid.UUID) (*RankedParticipant, error) {
	model, err := c.dal.GetRankedByPK(participantID)
	if err != nil {
		return nil, err
	}

	return NewRankedParticipant(model), nil
}

func (c *Controller) List(leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error) {
	models, err := c.dal.List(leaderboardId, limit, offset)
	if err != nil {
		return []*RankedParticipant{}, err
	}

	if len(models) <= 0 {
		return []*RankedParticipant{}, nil
	}

	participants := make([]*RankedParticipant, 0, len(models))
	for _, p := range models {
		participants = append(participants, NewRankedParticipant(p))
	}

	return participants, nil
}

func (c *Controller) UpdateScore(participantID uuid.UUID, amount int) (int, error) {
	model, err := c.dal.GetByPK(participantID)
	if err != nil {
		return 0, nil
	}

	prev := model.Score
	model.Score += amount
	if model.Score < 0 {
		model.Score = 0
	}

	err = c.dal.UpdateScore(model)
	if err != nil {
		return prev, err
	}

	return model.Score, nil
}

func (c *Controller) Create(leaderboardId uuid.UUID, externalId string, name string, metadata map[string]string) (*Participant, error) {
	lb, err := c.lbDal.GetByPK(leaderboardId)
	if err != nil {
		return nil, err
	}

	amount, err := c.dal.GetCount(leaderboardId)
	if err != nil {
		return nil, err
	}

	if amount >= lb.Capacity {
		return nil, fmt.Errorf("leaderboard already at capacity")
	}

	if err := c.validateName(name); err != nil {
		return nil, err
	}

	mData, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	model := &participant.Participant{
		LeaderboardID: leaderboardId,
		ExternalID:    externalId,
		Name:          name,
		Metadata:      string(mData),
		Score:         0,
	}

	if err := c.dal.Create(model); err != nil {
		return nil, err
	}

	return NewParticipant(model), nil
}

func (c *Controller) Remove(participantID uuid.UUID) error {
	model, err := c.dal.GetByPK(participantID)
	if err != nil {
		return err
	}

	return c.dal.Delete(model)
}

func (c *Controller) validateName(name string) error {
	if len(name) > 50 {
		return errors.New("participant name is too long")
	}

	// exists, err := c.dal.Exists(name)
	// if err != nil {
	// 	return err
	// }

	// if exists {
	// 	return fmt.Errorf("participant with name %s already exists", name)
	// }

	return nil
}
