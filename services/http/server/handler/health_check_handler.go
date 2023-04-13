package handler

import (
	"fmt"
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
	app "github.com/byyjoww/leaderboard/services/http/server"
	"github.com/byyjoww/leaderboard/services/http/server/response"
)

var (
	_ app.Handler = (*HealthCheckHandler)(nil)
)

type HealthCheckHandler struct {
	endpoint string
}

func NewHealthCheckHandler(endpoint string) *HealthCheckHandler {
	return &HealthCheckHandler{
		endpoint: endpoint,
	}
}

func (h *HealthCheckHandler) GetMethod() string {
	return http.MethodGet
}

func (h *HealthCheckHandler) GetPath() string {
	return h.endpoint
}

func (h *HealthCheckHandler) Handle(logger lib.Logger, r *http.Request) app.Response {
	return response.NewJsonStatusOK("The server is online\n")
}

func (h *HealthCheckHandler) String() string {
	return fmt.Sprintf("%T", h)
}
