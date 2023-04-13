package http

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/bll/participant"
	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/logging"
	"github.com/byyjoww/leaderboard/servers/api/handler"
	"github.com/byyjoww/leaderboard/services/http/decoder"
	"github.com/byyjoww/leaderboard/services/http/server"
	appMiddleware "github.com/byyjoww/leaderboard/services/http/server/middleware"
)

type Controllers struct {
	Leaderboard leaderboard.LeaderboardController
	Participant participant.ParticipantController
}

func New(logger logging.Logger, config config.HTTPServer, controllers Controllers) server.App {
	httpLogger := logging.NewHttpLogger(logger)
	decoder := decoder.New()

	mux := server.NewMux(httpLogger).WithMiddlewares(
		appMiddleware.NewPanicLoggerMiddleware(),
	).PathPrefixSubrouter("/api/v1")

	mux.WithMiddlewares(
		appMiddleware.NewBasicAuthMiddleware(
			config.Auth.User,
			config.Auth.Pass,
			config.Auth.Enabled,
		),
	)

	mux.AddHandlers(
		// Leaderboards
		handler.NewCreateLeaderboardHandler(decoder, controllers.Leaderboard),
		handler.NewGetLeaderboardHandler(decoder, controllers.Leaderboard),
		handler.NewListLeaderboardsHandler(decoder, controllers.Leaderboard),
		handler.NewRemoveLeaderboardHandler(decoder, controllers.Leaderboard),

		// Participants
		handler.NewUpdateScoreHandler(decoder, controllers.Participant),
		handler.NewGetParticipantHandler(decoder, controllers.Participant),
		handler.NewListParticipantsHandler(decoder, controllers.Participant),
		handler.NewCreateParticipantHandler(decoder, controllers.Participant),
		handler.NewRemoveParticipantHandler(decoder, controllers.Participant),
	)

	return &http.Server{
		Addr:    config.Address,
		Handler: mux.WithMetrics(),
	}
}
