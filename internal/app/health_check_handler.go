package app

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HealthCheckHandler struct {
	logger logrus.FieldLogger
}

func NewHealthCheckHandler(logger logrus.FieldLogger) *HealthCheckHandler {
	return &HealthCheckHandler{
		logger: logger,
	}
}

func (h *HealthCheckHandler) GetMethod() string {
	return http.MethodGet
}

func (h *HealthCheckHandler) GetPath() string {
	return "/healthcheck"
}

func (h *HealthCheckHandler) Handle(r *http.Request) Response {
	logger := h.logger.WithFields(logrus.Fields{
		"source": fmt.Sprintf("%T", h),
		"method": h.GetMethod(),
		"route":  h.GetPath(),
	})
	logger.Info("new request received")

	return NewJsonResponse("The server is online\n", http.StatusOK)
}
