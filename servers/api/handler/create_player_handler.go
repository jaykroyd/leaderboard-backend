package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreatePlayerHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller player.PlayerController
}

type CreatePlayerRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
}

func NewCreatePlayerHandler(logger logrus.FieldLogger, decoder app.Decoder, controller player.PlayerController) *CreatePlayerHandler {
	h := &CreatePlayerHandler{
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

func (h *CreatePlayerHandler) GetMethod() string {
	return http.MethodPost
}

func (h *CreatePlayerHandler) GetPath() string {
	return ""
}

func (h *CreatePlayerHandler) Handle(r *http.Request) app.Response {
	req := CreatePlayerRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("creating new player")
	player, err := h.controller.Create(req.LeaderboardID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.Info("new player created successfully")
	return app.NewStatusOK(player)
}
