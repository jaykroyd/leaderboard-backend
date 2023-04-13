package dal

import (
	"github.com/byyjoww/leaderboard/constants"
	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/byyjoww/leaderboard/dal/participant"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
)

type Factory interface {
	SetLogger(logger *PgLogger)
	NewLeaderboardDAL() leaderboard.LeaderboardDAL
	NewParticipantDAL() participant.ParticipantDAL
}

type PgFactory struct {
	db *pg.DB
}

func NewPgFactory(config Config) *PgFactory {
	return &PgFactory{
		db: pg.Connect(&pg.Options{
			User:        config.User,
			Password:    config.Pass,
			Addr:        config.Address(),
			Database:    config.Database,
			PoolSize:    config.PoolSize,
			MaxRetries:  config.MaxRetries,
			PoolTimeout: config.ConnectionTimeout,
		}),
	}
}

func (f *PgFactory) SetLogger(logger *PgLogger) {
	f.db.AddQueryHook(logger)
}

func (f *PgFactory) Ping() error {
	_, err := f.db.Exec("SELECT 1")
	if err != nil {
		return errors.Wrap(constants.ErrConnectingToDatabase, err.Error())
	}
	return nil
}

func (f *PgFactory) NewLeaderboardDAL() leaderboard.LeaderboardDAL {
	return leaderboard.NewDAL(f.db)
}

func (f *PgFactory) NewParticipantDAL() participant.ParticipantDAL {
	return participant.NewDAL(f.db)
}
