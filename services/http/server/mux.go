package server

import (
	"fmt"
	"net/http"
	"strings"

	lib "github.com/byyjoww/leaderboard/services/http"
	"github.com/gorilla/mux"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

type Mux struct {
	logger lib.Logger
	router *mux.Router
}

func NewMux(logger lib.Logger) *Mux {
	return &Mux{
		logger: logger,
		router: mux.NewRouter(),
	}
}

func (m *Mux) PathPrefixSubrouter(prefix string) *Mux {
	return &Mux{
		logger: m.logger,
		router: m.router.PathPrefix(prefix).Subrouter(),
	}
}

func (m *Mux) Subrouter() *Mux {
	return &Mux{
		logger: m.logger,
		router: m.router.NewRoute().Subrouter(),
	}
}

func (m *Mux) WithMiddlewares(middlewares ...Middleware) *Mux {
	for _, mdw := range middlewares {
		var (
			mFunc = func(next http.Handler) http.Handler {
				return mdw.ServeNext(m.logger, next)
			}
		)

		logger := m.logger.WithFields(map[string]interface{}{
			"middleware": fmt.Sprintf("%T", mdw),
		})
		logger.Infof("registering new middleware")

		m.router.Use(mFunc)
	}
	return m
}

func (m *Mux) AddHandlers(handlers ...Handler) {
	for _, h := range handlers {
		m.AddHTTPHandlers(
			NewHandlerConverter(h),
		)
	}
}

func (m *Mux) AddHTTPHandlers(handlers ...HTTPHandler) {
	for _, h := range handlers {
		var (
			method = h.GetMethod()
			path   = h.GetPath()
			serve  = func(w http.ResponseWriter, r *http.Request) {
				h.ServeHTTP(m.logger, w, r)
			}
		)

		logger := m.logger.WithFields(map[string]interface{}{
			"handler": fmt.Sprintf("%s", h),
			"method":  method,
			"path":    path,
		})
		logger.Infof("registering new handler")

		m.router.HandleFunc(path, serve).Methods(method)
		if !strings.HasPrefix(path, "/") {
			logger.Warnf("handler path doesn't start with a '/' character")
		}
	}
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func (m *Mux) WithMetrics() http.Handler {
	mdw := middleware.New(middleware.Config{
		Recorder:           metrics.NewRecorder(metrics.Config{}),
		DisableMeasureSize: true,
	})
	return std.Handler("", mdw, m)
}
