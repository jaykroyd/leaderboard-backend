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

type UpdateScoreHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

type AddScoreRequest struct {
	Amount int `json:"amount"`
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
	return "/leaderboards/{leaderboard_id}/participants/{participant_id}/score"
}

func (h *UpdateScoreHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	var (
		vars          = mux.Vars(r)
		leaderboardID = vars["leaderboard_id"]
		participantID = vars["participant_id"]
	)

	req := AddScoreRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardID,
		"participant_id": participantID,
	}).Info("updating participant score")

	score, err := h.controller.UpdateScore(uuid.MustParse(participantID), req.Amount)
	if err != nil {
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{"delta": req.Amount, "after": score}).Info("score updated successfully")
	return NewStatusOK(map[string]int{
		"new_score": score,
	})
}
