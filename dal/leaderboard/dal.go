package leaderboard

import (
	"github.com/byyjoww/leaderboard/constants"
	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type LeaderboardDAL interface {
	GetByPK(id uuid.UUID) (*Leaderboard, error)
	List(limit int, offset int) ([]*Leaderboard, error)
	ListByMode(mode int, limit int, offset int) ([]*Leaderboard, error)
	Create(leaderboard *Leaderboard) error
	Delete(leaderboard *Leaderboard) error
	Reset(leaderboardId uuid.UUID) error
}

type DAL struct {
	db *pg.DB
}

func NewDAL(db *pg.DB) *DAL {
	return &DAL{
		db: db,
	}
}

func (d *DAL) GetByPK(id uuid.UUID) (*Leaderboard, error) {
	leaderboard := &Leaderboard{ID: id}
	err := d.db.Model(leaderboard).
		WherePK().
		Select()
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}

func (d *DAL) List(limit int, offset int) ([]*Leaderboard, error) {
	var leaderboards []*Leaderboard
	err := d.db.Model(&leaderboards).
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}
	return leaderboards, nil
}

func (d *DAL) ListByMode(mode int, limit int, offset int) ([]*Leaderboard, error) {
	var leaderboards []*Leaderboard
	err := d.db.Model(&leaderboards).
		Where("mode = ?", mode).
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}
	return leaderboards, nil
}

func (d *DAL) Create(leaderboard *Leaderboard) error {
	_, err := d.db.Model(leaderboard).
		Set("created_at = now()").
		Set("updated_at = now()").
		Insert()
	if err != nil {
		return err
	}
	return nil
}

func (d *DAL) Delete(leaderboard *Leaderboard) error {
	return d.db.RunInTransaction(func(tx *pg.Tx) error {
		// List all participants in the leaderboard
		participants := []*participant.Participant{}
		err := tx.Model(&participants).
			Where("leaderboard_id = ?", leaderboard.ID).
			Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		if len(participants) > 0 {
			// Delete all participants in the leaderboard
			err = tx.Delete(&participants)
			if err != nil {
				if err != pg.ErrNoRows {
					return err
				}
			}
		}

		// Delete the leaderboard
		_, err = tx.Model(leaderboard).
			WherePK().
			Delete()
		if err != nil {
			if err == pg.ErrNoRows {
				return errors.Wrap(constants.ErrLeaderboardNotFound, err.Error())
			}
			return err
		}

		return nil
	})
}

func (d *DAL) Reset(leaderboardId uuid.UUID) error {
	// List all participants in the leaderboard
	participants := []*participant.Participant{}
	err := d.db.Model(&participants).
		Where("leaderboard_id = ?", leaderboardId).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	for _, p := range participants {
		p.Score = 0
	}

	if len(participants) > 0 {
		// Update the score to 0 for all participants in leaderboard
		err = d.db.Update(&participants)
		if err != nil {
			if err != pg.ErrNoRows {
				return err
			}
		}
	}
	return nil
}
