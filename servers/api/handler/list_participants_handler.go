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

type ListParticipantsHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

type ListParticipantsRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListParticipantsResponse struct {
	Participants []*participant.RankedParticipant `json:"participants"`
}

func NewListParticipantsHandler(decoder server.Decoder, controller participant.ParticipantController) *ListParticipantsHandler {
	return &ListParticipantsHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *ListParticipantsHandler) GetMethod() string {
	return http.MethodGet
}

func (h *ListParticipantsHandler) GetPath() string {
	return "/leaderboards/{leaderboard_id}/participants"
}

func (h *ListParticipantsHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("listing participants")

	var (
		vars          = mux.Vars(r)
		leaderboardID = vars["leaderboard_id"]
	)

	req := ListParticipantsRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	leaderboardUUID, err := uuid.Parse(leaderboardID)
	if err != nil {
		logger.WithError(err).Error("parsing leaderboard uuid")
		return NewBadRequest(err)
	}

	logger = logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardUUID,
		"request":        req,
	})

	participants, err := h.controller.List(r.Context(), leaderboardUUID, req.Limit, req.Offset)
	if err != nil {
		logger.WithError(err).Error("failed to retrieve participants")
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{
		"participants": participants,
		"amount":       len(participants),
	}).Info("successfully retrieved participants")

	return NewStatusOK(ListParticipantsResponse{
		Participants: participants,
	})
}
