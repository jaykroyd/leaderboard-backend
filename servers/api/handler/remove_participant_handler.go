package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/participant"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	return "/leaderboards/{leaderboard_id}/participants/{participant_id}"
}

func (h *RemoveParticipantHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	var (
		vars          = mux.Vars(r)
		leaderboardID = vars["leaderboard_id"]
		participantID = vars["participant_id"]
	)

	logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardID,
		"participant_id": participantID,
	}).Info("deleting participant")

	err := h.controller.Remove(uuid.MustParse(participantID))
	if err != nil {
		return NewInternalServerError(err)
	}

	logger.Info("participant deleted successfully")
	return NewStatusOK("OK")
}
