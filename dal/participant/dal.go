package participant

import (
	"fmt"

	"github.com/byyjoww/leaderboard/constants"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ParticipantDAL interface {
	Exists(name string) (bool, error)
	GetByPK(id uuid.UUID) (*Participant, error)
	GetRankedByPK(id uuid.UUID) (*RankedParticipant, error)
	GetCount(leaderboardId uuid.UUID) (int, error)
	List(leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error)
	Create(participant *Participant) error
	UpdateScore(participant *Participant) error
	Delete(participant *Participant) error
}

type DAL struct {
	db *pg.DB
}

func NewDAL(db *pg.DB) *DAL {
	return &DAL{
		db: db,
	}
}

func (d *DAL) Exists(name string) (bool, error) {
	participant := &Participant{Name: name}
	return d.db.Model(participant).
		Where("name = ?", name).
		Exists()
}

func (d *DAL) GetByPK(id uuid.UUID) (*Participant, error) {
	participant := &Participant{ID: id}
	err := d.db.Model(participant).
		WherePK().
		Select()
	if err != nil {
		return nil, err
	}
	return participant, nil
}

func (d *DAL) GetRankedByPK(id uuid.UUID) (*RankedParticipant, error) {
	participant := &RankedParticipant{Participant: &Participant{ID: id}}
	err := d.db.Model(participant).
		WherePK().
		Column("*").
		ColumnExpr("Row_number() OVER (PARTITION BY leaderboard_id ORDER BY score DESC) as rank").
		Select()
	if err != nil {
		return nil, err
	}
	return participant, nil
}

func (d *DAL) GetCount(leaderboardId uuid.UUID) (int, error) {
	return d.db.Model((*Participant)(nil)).
		Where("leaderboard_id = ?", leaderboardId).
		Count()
}

func (d *DAL) List(leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error) {
	var participants []*RankedParticipant
	err := d.db.Model(&participants).
		Where("leaderboard_id = ?", leaderboardId).
		Order("score DESC").
		Column("*").
		ColumnExpr("Row_number() OVER (PARTITION BY leaderboard_id ORDER BY score DESC) as rank").
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}
	return participants, nil
}

func (d *DAL) Create(participant *Participant) error {
	_, err := d.db.Model(participant).
		Set("created_at = now()").
		Set("updated_at = now()").
		Insert()
	if err != nil {
		return err
	}
	return nil
}

func (d *DAL) UpdateScore(participant *Participant) error {
	_, err := d.db.Model(participant).
		WherePK().
		Set("score = ?", participant.Score).
		Set("updated_at = now()").
		Update()
	if err != nil {
		if err == pg.ErrNoRows {
			return errors.Wrap(
				constants.ErrParticipantNotFound,
				fmt.Sprintf("%s (id %s)", err.Error(), participant.ID),
			)
		}
		return err
	}
	return nil
}

func (d *DAL) Delete(participant *Participant) error {
	_, err := d.db.Model(participant).
		WherePK().
		Delete()
	if err != nil {
		if err == pg.ErrNoRows {
			return errors.Wrap(constants.ErrParticipantNotFound, err.Error())
		}
		return err
	}
	return nil
}
