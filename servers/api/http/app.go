package http

import (
	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/dal"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/byyjoww/leaderboard/internal/decoder"
	"github.com/byyjoww/leaderboard/servers/api/handler"
	"github.com/sirupsen/logrus"
)

func New(logger logrus.FieldLogger, config config.Config) *app.App {
	a := app.NewApp(logger, app.Config{
		Address: config.Http.Address,
		Auth: app.Auth{
			Enabled: config.Http.Auth.Enabled,
			User:    config.Http.Auth.User,
			Pass:    config.Http.Auth.Pass,
		},
	})

	dalFactory := dal.NewPgFactory(dal.Config{
		User:              config.Postgres.User,
		Pass:              config.Postgres.Pass,
		Host:              config.Postgres.Host,
		Port:              config.Postgres.Port,
		Database:          config.Postgres.Database,
		PoolSize:          config.Postgres.PoolSize,
		MaxRetries:        config.Postgres.MaxRetries,
		ConnectionTimeout: config.Postgres.ConnectionTimeout,
	})

	lbDal := dalFactory.NewLeaderboardDAL()
	playerDal := dalFactory.NewPlayerDAL()

	lbController := leaderboard.NewController(lbDal)
	playerController := player.NewController(playerDal)

	decoder := decoder.New()
	api := a.Router().Subrouter("/api")
	api.AddHandlers(
		handler.NewCreateLeaderboardHandler(logger, lbController),
		handler.NewGetLeaderboardHandler(logger, decoder, lbController),
		handler.NewListLeaderboardsHandler(logger, lbController),
		handler.NewRemoveLeaderboardHandler(logger, decoder, lbController),
	)

	players := api.Subrouter("/player")
	players.AddHandlers(
		handler.NewUpdateScoreHandler(logger, decoder, playerController),
		handler.NewGetScoreHandler(logger, decoder, playerController),
		handler.NewCreatePlayerHandler(logger, decoder, playerController),
		handler.NewRemovePlayerHandler(logger, decoder, playerController),
	)

	return a
}
