package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/participant"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
)

type GetParticipantHandler struct {
	decoder    server.Decoder
	controller participant.ParticipantController
}

type GetParticipantRequest struct {
	ID uuid.UUID `json:"id"`
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
	return "/leaderboards/{leaderboard_id}/participants/{participant_id}"
}

func (h *GetParticipantHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	req := GetParticipantRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger.WithFields(logging.Fields{"request": req}).Info("getting participant")
	participant, err := h.controller.Get(req.ID)
	if err != nil {
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{"participant": participant}).Info("successfully retrieved participant")
	return NewStatusOK(participant)
}
