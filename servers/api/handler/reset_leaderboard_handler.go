package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
)

type ResetLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.LeaderboardController
}

type ResetLeaderboardRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
	Limit         int       `json:"limit"`
	Offset        int       `json:"offset"`
}

func NewResetLeaderboardHandler(decoder server.Decoder, controller leaderboard.LeaderboardController) *ResetLeaderboardHandler {
	return &ResetLeaderboardHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *ResetLeaderboardHandler) GetMethod() string {
	return http.MethodPatch
}

func (h *ResetLeaderboardHandler) GetPath() string {
	return "/leaderboard/reset"
}

func (h *ResetLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	req := ResetLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger.WithFields(logging.Fields{"request": req}).Info("resetting leaderboard")
	err := h.controller.Reset(req.LeaderboardID)
	if err != nil {
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{"leaderboard": req.LeaderboardID}).Info("successfully reset leaderboard")
	return NewStatusOK("OK")
}
