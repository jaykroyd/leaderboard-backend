package server

import (
	"fmt"
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
)

var (
	_ HTTPHandler = (*HandlerConverter)(nil)
)

type HandlerConverter struct {
	handler Handler
}

func NewHandlerConverter(handler Handler) *HandlerConverter {
	return &HandlerConverter{
		handler: handler,
	}
}

func (h *HandlerConverter) GetMethod() string {
	return h.handler.GetMethod()
}

func (h *HandlerConverter) GetPath() string {
	return h.handler.GetPath()
}

func (h *HandlerConverter) ServeHTTP(logger lib.Logger, w http.ResponseWriter, r *http.Request) {
	logger = logger.WithFields(map[string]interface{}{
		"handler": fmt.Sprintf("%T", h.handler),
		"method":  h.GetMethod(),
		"route":   h.GetPath(),
	})
	logger.Info("new request received")

	resp := h.handler.Handle(logger, r)
	Respond(logger, w, resp)
}

func (h *HandlerConverter) String() string {
	return fmt.Sprintf("%T", h.handler)
}
