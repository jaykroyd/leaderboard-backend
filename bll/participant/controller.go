package participant

import (
	"errors"

	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/google/uuid"
)

type ParticipantController interface {
	Get(participantID uuid.UUID) (*participant.RankedParticipant, error)
	List(leaderboardId uuid.UUID, limit int, offset int) ([]*participant.RankedParticipant, error)
	UpdateScore(participantID uuid.UUID, amount int) (int, error)
	Create(leaderboardId uuid.UUID, name string) (*participant.Participant, error)
	Remove(participantID uuid.UUID) error
}

type Controller struct {
	dal participant.ParticipantDAL
}

func NewController(dal participant.ParticipantDAL) *Controller {
	return &Controller{
		dal: dal,
	}
}

func (c *Controller) Get(participantID uuid.UUID) (*participant.RankedParticipant, error) {
	return c.dal.GetRankedByPK(participantID)
}

func (c *Controller) List(leaderboardId uuid.UUID, limit int, offset int) ([]*participant.RankedParticipant, error) {
	return c.dal.List(leaderboardId, limit, offset)
}

func (c *Controller) UpdateScore(participantID uuid.UUID, amount int) (int, error) {
	p, err := c.dal.GetByPK(participantID)
	if err != nil {
		return 0, nil
	}

	prev := p.Score
	p.Score += amount
	if p.Score < 0 {
		p.Score = 0
	}

	err = c.dal.UpdateScore(p)
	if err != nil {
		return prev, err
	}

	return p.Score, nil
}

func (c *Controller) Create(leaderboardId uuid.UUID, name string) (*participant.Participant, error) {
	if err := c.validateName(name); err != nil {
		return nil, err
	}

	p := &participant.Participant{
		LeaderboardID: leaderboardId,
		Name:          name,
		Score:         0,
	}

	return p, c.dal.Create(p)
}

func (c *Controller) Remove(participantID uuid.UUID) error {
	p, err := c.dal.GetByPK(participantID)
	if err != nil {
		return err
	}

	return c.dal.Delete(p)
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
