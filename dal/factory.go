package dal

import (
	"github.com/byyjoww/leaderboard/dal/leaderboard"
	"github.com/byyjoww/leaderboard/dal/player"
	"github.com/go-pg/pg"
)

type Factory interface {
	SetLogger(logger *PgLogger)
	NewLeaderboardDAL() leaderboard.LeaderboardDAL
	NewPlayerDAL() player.PlayerDAL
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

func (f *PgFactory) NewLeaderboardDAL() leaderboard.LeaderboardDAL {
	return leaderboard.NewDAL(f.db)
}

func (f *PgFactory) NewPlayerDAL() player.PlayerDAL {
	return player.NewDAL(f.db)
}
