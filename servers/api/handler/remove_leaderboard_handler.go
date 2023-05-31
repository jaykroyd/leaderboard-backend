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

type RemoveLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.LeaderboardController
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
	return "/leaderboard/{leaderboard_id}"
}

func (h *RemoveLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("deleting leaderboard")

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

	err = h.controller.Remove(r.Context(), leaderboardUUID)
	if err != nil {
		logger.WithError(err).Error("failed to delete leaderboard")
		return NewInternalServerError(err)
	}

	logger.Info("leaderboard deleted successfully")
	return NewStatusOK("OK")
}
