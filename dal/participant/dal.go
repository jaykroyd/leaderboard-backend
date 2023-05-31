package participant

import (
	"context"
	"fmt"

	"github.com/byyjoww/leaderboard/constants"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ParticipantDAL interface {
	Exists(ctx context.Context, name string) (bool, error)
	GetByPK(ctx context.Context, id uuid.UUID) (*Participant, error)
	GetRankedByPK(ctx context.Context, id uuid.UUID) (*RankedParticipant, error)
	GetByExternalID(ctx context.Context, leaderboardId uuid.UUID, externalID string) (*Participant, error)
	GetRankedByExternalID(ctx context.Context, leaderboardId uuid.UUID, externalID string) (*RankedParticipant, error)
	GetCount(ctx context.Context, leaderboardId uuid.UUID) (int, error)
	List(ctx context.Context, leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error)
	Create(ctx context.Context, participant *Participant) error
	UpdateScore(ctx context.Context, participant *Participant) error
	Delete(ctx context.Context, participant *Participant) error
}

type DAL struct {
	db *pg.DB
}

func NewDAL(db *pg.DB) *DAL {
	return &DAL{
		db: db,
	}
}

func (d *DAL) Exists(ctx context.Context, name string) (bool, error) {
	participant := &Participant{Name: name}
	return d.db.Model(participant).
		Context(ctx).
		Where("name = ?", name).
		Exists()
}

func (d *DAL) GetByPK(ctx context.Context, id uuid.UUID) (*Participant, error) {
	participant := &Participant{ID: id}
	err := d.db.Model(participant).
		Context(ctx).
		WherePK().
		Select()
	return participant, err
}

func (d *DAL) GetRankedByPK(ctx context.Context, id uuid.UUID) (*RankedParticipant, error) {
	participant := &RankedParticipant{Participant: &Participant{ID: id}}
	err := d.db.Model(participant).
		Context(ctx).
		Column("*").
		ColumnExpr("ROW_NUMBER() OVER (PARTITION BY section_id ORDER BY score DESC) as rank").
		Where("external_id = ?", id).
		Select()
	return participant, err
}

func (d *DAL) GetByExternalID(ctx context.Context, leaderboardId uuid.UUID, externalID string) (*Participant, error) {
	participant := &Participant{
		LeaderboardID: leaderboardId,
		ExternalID:    externalID,
	}
	err := d.db.Model(participant).
		Context(ctx).
		Where("external_id = ?", externalID).
		Where("leaderboard_id = ?", leaderboardId).
		Select()
	return participant, err
}

func (d *DAL) GetRankedByExternalID(ctx context.Context, leaderboardId uuid.UUID, externalID string) (*RankedParticipant, error) {
	participant := &RankedParticipant{
		Participant: &Participant{
			LeaderboardID: leaderboardId,
			ExternalID:    externalID,
		},
	}

	err := d.db.Model(participant).
		Context(ctx).
		Where("external_id = ?", externalID).
		Where("leaderboard_id = ?", leaderboardId).
		Column("*").
		ColumnExpr("(SELECT COUNT(1) + 1 FROM participants pi WHERE pi.leaderboard_id = ? and pi.score > ? and pi.created_at < ?) as rank",
			participant.LeaderboardID, participant.Score, participant.CreatedAt).
		Select()
	return participant, err
}

func (d *DAL) GetCount(ctx context.Context, leaderboardId uuid.UUID) (int, error) {
	return d.db.Model((*Participant)(nil)).
		Context(ctx).
		Where("leaderboard_id = ?", leaderboardId).
		Count()
}

func (d *DAL) List(ctx context.Context, leaderboardId uuid.UUID, limit int, offset int) ([]*RankedParticipant, error) {
	var participants []*RankedParticipant
	err := d.db.Model(&participants).
		Context(ctx).
		Where("leaderboard_id = ?", leaderboardId).
		Order("score DESC").
		Column("*").
		ColumnExpr("ROW_NUMBER() OVER (PARTITION BY leaderboard_id ORDER BY score DESC) as rank").
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}
	return participants, nil
}

func (d *DAL) Create(ctx context.Context, participant *Participant) error {
	_, err := d.db.Model(participant).
		Context(ctx).
		Set("created_at = now()").
		Set("updated_at = now()").
		Insert()
	return err
}

func (d *DAL) UpdateScore(ctx context.Context, participant *Participant) error {
	_, err := d.db.Model(participant).
		Context(ctx).
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

func (d *DAL) Delete(ctx context.Context, participant *Participant) error {
	_, err := d.db.Model(participant).
		Context(ctx).
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
