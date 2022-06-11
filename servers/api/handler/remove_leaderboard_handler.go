package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RemoveLeaderboardHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller leaderboard.LeaderboardController
}

type RemoveLeaderboardRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
}

func NewRemoveLeaderboardHandler(logger logrus.FieldLogger, decoder app.Decoder, controller leaderboard.LeaderboardController) *RemoveLeaderboardHandler {
	h := &RemoveLeaderboardHandler{
		logger:     logger,
		decoder:    decoder,
		controller: controller,
	}

	h.logger = h.logger.WithFields(logrus.Fields{
		"source": fmt.Sprintf("%T", h),
		"method": h.GetMethod(),
		"route":  h.GetPath(),
	})

	return h
}

func (h *RemoveLeaderboardHandler) GetMethod() string {
	return http.MethodDelete
}

func (h *RemoveLeaderboardHandler) GetPath() string {
	return "/leaderboard"
}

func (h *RemoveLeaderboardHandler) Handle(r *http.Request) app.Response {
	req := RemoveLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("deleting leaderboard")
	err := h.controller.Remove(req.LeaderboardID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"id": req.LeaderboardID}).Info("leaderboard deleted successfully")
	return app.NewStatusOK("OK")
}
