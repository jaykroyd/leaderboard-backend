package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ResetLeaderboardHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller leaderboard.LeaderboardController
}

type ResetLeaderboardRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
	Limit         int       `json:"limit"`
	Offset        int       `json:"offset"`
}

func NewResetLeaderboardHandler(logger logrus.FieldLogger, decoder app.Decoder, controller leaderboard.LeaderboardController) *ResetLeaderboardHandler {
	h := &ResetLeaderboardHandler{
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

func (h *ResetLeaderboardHandler) GetMethod() string {
	return http.MethodPatch
}

func (h *ResetLeaderboardHandler) GetPath() string {
	return "/leaderboard/reset"
}

func (h *ResetLeaderboardHandler) Handle(r *http.Request) app.Response {
	req := ResetLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("resetting leaderboard")
	err := h.controller.Reset(req.LeaderboardID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"leaderboard": req.LeaderboardID}).Info("successfully reset leaderboard")
	return app.NewStatusOK("OK")
}
