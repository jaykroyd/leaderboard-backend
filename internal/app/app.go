package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger logrus.FieldLogger
	config Config
	router Router
}

type Config struct {
	Address string
	Auth    Auth
}

type Auth struct {
	Enabled bool
	User    string
	Pass    string
}

func NewApp(logger logrus.FieldLogger, config Config) *App {
	app := &App{
		logger: logger,
		config: config,
		router: NewMuxRouter(logger, mux.NewRouter()),
	}

	app.configureRoutes()
	return app
}

func (a *App) Router() Router {
	return a.router
}

func (a *App) configureRoutes() {
	a.logger.Infof("creating router at %s", a.config.Address)

	a.Router().AddHandlers(
		NewHealthCheckHandler(a.logger),
	)

	a.Router().AddMiddlewares(
		NewPanicLoggerMiddleware(a.logger),
		NewBasicAuthMiddleware(
			a.logger,
			a.config.Auth.User,
			a.config.Auth.Pass,
			a.config.Auth.Enabled,
		))

	a.logger.Info("base server configuration complete")
}

func (a *App) ListenAndServe() {
	a.logger.Info(fmt.Sprintf("Starting listening on %s", a.config.Address))
	err := http.ListenAndServe(a.config.Address, a.router.Handler())
	if err != nil {
		a.logger.WithError(err).Error("error on starting private server")
		os.Exit(1)
	}
}
