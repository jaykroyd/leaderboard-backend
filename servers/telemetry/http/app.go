package http

import (
	"net/http"

	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/logging"
	"github.com/byyjoww/leaderboard/services/http/server"

	"github.com/byyjoww/leaderboard/services/http/server/handler"
	"github.com/byyjoww/leaderboard/services/http/server/middleware"
)

func New(logger logging.Logger, config config.HTTPServer) server.App {
	httpLogger := logging.NewHttpLogger(logger)
	mux := server.NewMux(httpLogger).WithMiddlewares(
		middleware.NewPanicLoggerMiddleware(),
		middleware.NewBasicAuthMiddleware(
			config.Auth.User,
			config.Auth.Pass,
			config.Auth.Enabled,
		),
	)

	mux.AddHandlers(
		handler.NewHealthCheckHandler("/health"),
	)

	mux.AddHTTPHandlers(
		handler.NewMetricsHandler("/metrics"),
	)

	return &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
}
