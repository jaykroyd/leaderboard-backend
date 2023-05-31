package participant

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ParticipantController interface {
	Get(ctx context.Context, leaderboardId uuid.UUID, externalId string) (*RankedParticipant, error)
	List(ctx context.Context, leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error)
	UpdateScore(ctx context.Context, leaderboardId uuid.UUID, externalId string, score int) (*RankedParticipant, error)
	Create(ctx context.Context, leaderboardId uuid.UUID, externalId string, name string, metadata map[string]string) (*Participant, error)
	Remove(ctx context.Context, leaderboardId uuid.UUID, externalId string) error
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

func (c *Controller) Get(ctx context.Context, leaderboardId uuid.UUID, externalId string) (*RankedParticipant, error) {
	model, err := c.dal.GetRankedByExternalID(ctx, leaderboardId, externalId)
	if err != nil {
		return nil, err
	}

	return NewRankedParticipant(model), nil
}

func (c *Controller) List(ctx context.Context, leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error) {
	models, err := c.dal.List(ctx, leaderboardId, limit, offset)
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

func (c *Controller) UpdateScore(ctx context.Context, leaderboardId uuid.UUID, externalId string, score int) (*RankedParticipant, error) {
	model, err := c.dal.GetByExternalID(ctx, leaderboardId, externalId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get participant from dal")
	}

	model.Score += score
	if model.Score < 0 {
		model.Score = 0
	}

	if err = c.dal.UpdateScore(ctx, model); err != nil {
		return nil, errors.Wrap(err, "failed to update score in dal")
	}

	rModel, err := c.dal.GetRankedByPK(ctx, model.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get final ranked participant from dal")
	}

	return NewRankedParticipant(rModel), nil
}

func (c *Controller) Create(ctx context.Context, leaderboardId uuid.UUID, externalId string, name string, metadata map[string]string) (*Participant, error) {
	lb, err := c.lbDal.GetByPK(ctx, leaderboardId)
	if err != nil {
		return nil, err
	}

	amount, err := c.dal.GetCount(ctx, leaderboardId)
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

	if err := c.dal.Create(ctx, model); err != nil {
		return nil, err
	}

	return NewParticipant(model), nil
}

func (c *Controller) Remove(ctx context.Context, leaderboardId uuid.UUID, externalId string) error {
	model, err := c.dal.GetByExternalID(ctx, leaderboardId, externalId)
	if err != nil {
		return err
	}

	return c.dal.Delete(ctx, model)
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
