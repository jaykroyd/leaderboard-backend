package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/sirupsen/logrus"
)

type CreateLeaderboardHandler struct {
	logger     logrus.FieldLogger
	controller leaderboard.LeaderboardController
}

func NewCreateLeaderboardHandler(logger logrus.FieldLogger, controller leaderboard.LeaderboardController) *CreateLeaderboardHandler {
	h := &CreateLeaderboardHandler{
		logger:     logger,
		controller: controller,
	}

	h.logger = h.logger.WithFields(logrus.Fields{
		"source": fmt.Sprintf("%T", h),
		"method": h.GetMethod(),
		"route":  h.GetPath(),
	})

	return h
}

func (h *CreateLeaderboardHandler) GetMethod() string {
	return http.MethodPost
}

func (h *CreateLeaderboardHandler) GetPath() string {
	return "/leaderboard"
}

func (h *CreateLeaderboardHandler) Handle(r *http.Request) app.Response {
	leaderboard, err := h.controller.Create()
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.Info("new leaderboard created successfully")
	return app.NewStatusOK(leaderboard)
}
