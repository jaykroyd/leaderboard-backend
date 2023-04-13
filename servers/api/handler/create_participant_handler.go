package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/participant"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
)

type CreateParticipantHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

type CreateParticipantRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
	ID            string    `json:"id"`
	Name          string    `json:"name"`
}

func NewCreateParticipantHandler(decoder server.Decoder, controller participant.ParticipantController) *CreateParticipantHandler {
	return &CreateParticipantHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *CreateParticipantHandler) GetMethod() string {
	return http.MethodPost
}

func (h *CreateParticipantHandler) GetPath() string {
	return "/participants"
}

func (h *CreateParticipantHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	req := CreateParticipantRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger.WithFields(logging.Fields{"request": req}).Info("creating new participant")
	participant, err := h.controller.Create(req.LeaderboardID, req.Name)
	if err != nil {
		return NewInternalServerError(err)
	}

	logger.Info("new participant created successfully")
	return NewStatusOK(participant)
}