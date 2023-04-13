package handler

import (
	"fmt"
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsHandler struct {
	metricsHandler http.Handler
	endpoint       string
}

func NewMetricsHandler(endpoint string) *MetricsHandler {
	return &MetricsHandler{
		metricsHandler: promhttp.Handler(),
		endpoint:       endpoint,
	}
}

func (h *MetricsHandler) GetMethod() string {
	return http.MethodGet
}

func (h *MetricsHandler) GetPath() string {
	return h.endpoint
}

func (h *MetricsHandler) ServeHTTP(logger lib.Logger, w http.ResponseWriter, r *http.Request) {
	h.metricsHandler.ServeHTTP(w, r)
}

func (h *MetricsHandler) String() string {
	return fmt.Sprintf("%T", h)
}
