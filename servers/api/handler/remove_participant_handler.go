package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/participant"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type RemoveParticipantHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

func NewRemoveParticipantHandler(decoder server.Decoder, controller participant.ParticipantController) *RemoveParticipantHandler {
	return &RemoveParticipantHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *RemoveParticipantHandler) GetMethod() string {
	return http.MethodDelete
}

func (h *RemoveParticipantHandler) GetPath() string {
	return "/leaderboards/{leaderboard_id}/participants/{external_id}"
}

func (h *RemoveParticipantHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("deleting participant")

	var (
		vars          = mux.Vars(r)
		leaderboardID = vars["leaderboard_id"]
		externalID    = vars["external_id"]
	)

	logger = logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardID,
		"external_id":    externalID,
	})

	leaderboardUUID, err := uuid.Parse(leaderboardID)
	if err != nil {
		logger.WithError(err).Error("failed to parse leaderboard id")
		return NewBadRequest(errors.Wrap(err, "failed to parse leaderboard id"))
	}

	err = h.controller.Remove(r.Context(), leaderboardUUID, externalID)
	if err != nil {
		logger.WithError(err).Error("failed to delete participant")
		return NewInternalServerError(err)
	}

	logger.Info("participant deleted successfully")
	return NewStatusOK("OK")
}
