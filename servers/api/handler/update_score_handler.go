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

type UpdateScoreHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

type UpdateScoreRequest struct {
	Score int `json:"score"`
}

type UpdateScoreResponse struct {
	Participant *participant.RankedParticipant `json:"participant"`
}

func NewUpdateScoreHandler(decoder server.Decoder, controller participant.ParticipantController) *UpdateScoreHandler {
	return &UpdateScoreHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *UpdateScoreHandler) GetMethod() string {
	return http.MethodPatch
}

func (h *UpdateScoreHandler) GetPath() string {
	return "/leaderboards/{leaderboard_id}/participants/{external_id}/score"
}

func (h *UpdateScoreHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("updating participant score")

	var (
		vars          = mux.Vars(r)
		leaderboardID = vars["leaderboard_id"]
		externalID    = vars["external_id"]
	)

	logger = logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardID,
		"external_id":    externalID,
	})

	req := UpdateScoreRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	leaderboardUUID, err := uuid.Parse(leaderboardID)
	if err != nil {
		logger.WithError(err).Error("failed to parse leaderboard id")
		return NewBadRequest(errors.Wrap(err, "failed to parse leaderboard id"))
	}

	participant, err := h.controller.UpdateScore(r.Context(), leaderboardUUID, externalID, req.Score)
	if err != nil {
		logger.WithError(err).Error("failed to update score")
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{
		"delta":       req.Score,
		"after":       participant.Score,
		"participant": participant,
	}).Info("score updated successfully")

	return NewStatusOK(UpdateScoreResponse{
		Participant: participant,
	})
}
