package cmd

import (
	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/bll/participant"
	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/dal"
	"github.com/byyjoww/leaderboard/logging"
	apiApp "github.com/byyjoww/leaderboard/servers/api/http"
	telemetryApp "github.com/byyjoww/leaderboard/servers/telemetry/http"
	"github.com/byyjoww/leaderboard/services/http/server"
)

func startAPI(logger logging.Logger, cfg config.Config) {
	dalFactory := dal.NewPgFactory(dal.Config{
		User:              cfg.Postgres.User,
		Pass:              cfg.Postgres.Pass,
		Host:              cfg.Postgres.Host,
		Port:              cfg.Postgres.Port,
		Database:          cfg.Postgres.Database,
		PoolSize:          cfg.Postgres.PoolSize,
		MaxRetries:        cfg.Postgres.MaxRetries,
		ConnectionTimeout: cfg.Postgres.ConnectionTimeout,
	})

	if err := dalFactory.Ping(); err != nil {
		logger.WithFields(logging.Fields{
			"host":     cfg.Postgres.Host,
			"port":     cfg.Postgres.Port,
			"database": cfg.Postgres.Database,
		}).WithError(err).Fatal("failed to confirm active database connection")
	}

	lbDal := dalFactory.NewLeaderboardDAL()
	participantDal := dalFactory.NewParticipantDAL()

	api := apiApp.New(logger, cfg.Http.API, apiApp.Controllers{
		Leaderboard: leaderboard.NewController(lbDal),
		Participant: participant.NewController(participantDal, lbDal),
	})

	telemetry := telemetryApp.New(logger, cfg.Http.Telemetry)

	logger.Info("App initialized succesfully")
	if err := server.ListenAndServe(api, telemetry); err != nil {
		logger.WithError(err).Fatal("failed to start server")
	}
}
