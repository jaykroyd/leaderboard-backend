package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type ResetLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.LeaderboardController
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
	return "/leaderboard/{leaderboard_id}/reset"
}

func (h *ResetLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("resetting leaderboard")

	var (
		vars          = mux.Vars(r)
		leaderboardID = vars["leaderboard_id"]
	)

	logger = logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardID,
	})

	leaderboardUUID, err := uuid.Parse(leaderboardID)
	if err != nil {
		logger.WithError(err).Error("failed to parse leaderboard id")
		return NewBadRequest(errors.Wrap(err, "failed to parse leaderboard id"))
	}

	err = h.controller.Reset(r.Context(), leaderboardUUID)
	if err != nil {
		logger.WithError(err).Error("failed to reset leaderboard")
		return NewInternalServerError(err)
	}

	logger.Info("successfully reset leaderboard")
	return NewStatusOK("OK")
}
