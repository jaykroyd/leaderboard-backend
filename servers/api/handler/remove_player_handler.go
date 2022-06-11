package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RemovePlayerHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller player.PlayerController
}

type RemovePlayerRequest struct {
	PlayerID uuid.UUID `json:"player_id"`
}

func NewRemovePlayerHandler(logger logrus.FieldLogger, decoder app.Decoder, controller player.PlayerController) *RemovePlayerHandler {
	h := &RemovePlayerHandler{
		logger:     logger,
		decoder:    decoder,
		controller: controller,
	}

	h.logger = h.logger.WithFields(logrus.Fields{
		"source": fmt.Sprintf("%T", h),
		"method": h.GetMethod(),
		"route":  h.GetPath(),
	})

	return h
}

func (h *RemovePlayerHandler) GetMethod() string {
	return http.MethodDelete
}

func (h *RemovePlayerHandler) GetPath() string {
	return ""
}

func (h *RemovePlayerHandler) Handle(r *http.Request) app.Response {
	req := RemovePlayerRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("deleting player")
	err := h.controller.Remove(req.PlayerID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"id": req.PlayerID}).Info("player deleted successfully")
	return app.NewStatusOK("OK")
}
