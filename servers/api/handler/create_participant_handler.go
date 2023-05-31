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
	LeaderboardID uuid.UUID         `json:"leaderboard_id"`
	ExternalID    string            `json:"external_id"`
	Name          string            `json:"name"`
	Metadata      map[string]string `json:"metadata"`
}

type CreateParticipantResponse struct {
	Participant *participant.Participant `json:"participant"`
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
	logger.Info("creating new participant")

	req := CreateParticipantRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger = logger.WithFields(logging.Fields{
		"request": req,
	})

	participant, err := h.controller.Create(r.Context(), req.LeaderboardID, req.ExternalID, req.Name, req.Metadata)
	if err != nil {
		logger.WithError(err).Error("failed to create participant")
		return NewInternalServerError(err)
	}

	logger.Info("new participant created successfully")
	return NewStatusOK(CreateParticipantResponse{
		Participant: participant,
	})
}
