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

type GetParticipantHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

func NewGetParticipantHandler(decoder server.Decoder, controller participant.ParticipantController) *GetParticipantHandler {
	return &GetParticipantHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *GetParticipantHandler) GetMethod() string {
	return http.MethodGet
}

func (h *GetParticipantHandler) GetPath() string {
	return "/leaderboards/{leaderboard_id}/participants/{external_id}"
}

func (h *GetParticipantHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("getting participant")

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

	participant, err := h.controller.Get(r.Context(), leaderboardUUID, externalID)
	if err != nil {
		logger.WithError(err).Error("failed to retrieve participant")
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{
		"participant": participant,
	}).Info("successfully retrieved participant")

	return NewStatusOK(participant)
}
