package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Router interface {
	Subrouter(prefix string) Router
	AddHandlers(handlers ...Handler)
	AddMiddlewares(middlewares ...Middleware)
	Handler() http.Handler
}

type MuxRouter struct {
	logger logrus.FieldLogger
	router *mux.Router
}

func NewMuxRouter(logger logrus.FieldLogger, router *mux.Router) *MuxRouter {
	return &MuxRouter{
		logger: logger,
		router: router,
	}
}

func (a *MuxRouter) Subrouter(prefix string) Router {
	return NewMuxRouter(a.logger, a.router.PathPrefix(prefix).Subrouter())
}

func (a *MuxRouter) Handler() http.Handler {
	return a.router
}

func (a *MuxRouter) AddMiddlewares(middlewares ...Middleware) {
	for _, m := range middlewares {
		a.logger.Infof("registering new middleware %T", m)
		a.router.Use(m.ServeHTTP)
	}
}

func (a *MuxRouter) AddHandlers(handlers ...Handler) {
	for _, h := range handlers {
		var (
			method  = h.GetMethod()
			path    = h.GetPath()
			wrapper = a.handleFunc(h)
		)

		logger := a.logger.WithFields(logrus.Fields{
			"method": method,
			"path":   path,
		})
		logger.Infof("registering new handler %T", h)

		a.router.HandleFunc(path, wrapper).Methods(method)
		if !strings.HasPrefix(path, "/") {
			logger.Warnf("handler path doesn't start with a '/' character")
		}
	}
}

func (a *MuxRouter) handleFunc(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := a.logger.WithFields(logrus.Fields{
			"source": fmt.Sprintf("%T", h),
			"method": h.GetMethod(),
			"route":  h.GetPath(),
		})
		logger.Info("new request received")

		response, contentType, status := a.handle(h, r)
		if status == 0 {
			a.logger.Errorf("the handler at %s: %s is attempting to return status code 0", h.GetMethod(), h.GetPath())
			status = http.StatusInternalServerError
		}

		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(status)
		if len(response) > 0 {
			if _, err := w.Write(response); err != nil {
				a.logger.WithError(err).Error("could not write response.")
			}
		}
	}
}

func (a *MuxRouter) handle(h Handler, r *http.Request) ([]byte, string, int) {
	resp := h.Handle(r)
	bytes, err := resp.GetContent()
	if err != nil {
		a.logger.WithError(err).Error()
		return []byte{}, resp.GetContentType(), resp.GetStatusCode()
	}

	return bytes, resp.GetContentType(), resp.GetStatusCode()
}
