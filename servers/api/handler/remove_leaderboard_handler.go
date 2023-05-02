package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
)

type RemoveLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.LeaderboardController
}

type RemoveLeaderboardRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
}

func NewRemoveLeaderboardHandler(decoder server.Decoder, controller leaderboard.LeaderboardController) *RemoveLeaderboardHandler {
	return &RemoveLeaderboardHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *RemoveLeaderboardHandler) GetMethod() string {
	return http.MethodDelete
}

func (h *RemoveLeaderboardHandler) GetPath() string {
	return "/leaderboard"
}

func (h *RemoveLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	req := RemoveLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger.WithFields(logging.Fields{"request": req}).Info("deleting leaderboard")
	err := h.controller.Remove(req.LeaderboardID)
	if err != nil {
		logger.WithError(err).Error("failed to delete leaderboard")
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{"id": req.LeaderboardID}).Info("leaderboard deleted successfully")
	return NewStatusOK("OK")
}
