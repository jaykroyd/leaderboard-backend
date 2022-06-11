package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/sirupsen/logrus"
)

type ListLeaderboardHandler struct {
	logger     logrus.FieldLogger
	controller leaderboard.LeaderboardController
}

func NewListLeaderboardsHandler(logger logrus.FieldLogger, controller leaderboard.LeaderboardController) *ListLeaderboardHandler {
	h := &ListLeaderboardHandler{
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

func (h *ListLeaderboardHandler) GetMethod() string {
	return http.MethodGet
}

func (h *ListLeaderboardHandler) GetPath() string {
	return "/leaderboards"
}

func (h *ListLeaderboardHandler) Handle(r *http.Request) app.Response {
	lbs, err := h.controller.List()
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"leaderboards": lbs, "amount": len(lbs)}).Info("successfully retrieved leaderboards")
	return app.NewStatusOK(lbs)
}
