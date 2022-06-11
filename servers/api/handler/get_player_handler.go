package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetPlayerHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller player.PlayerController
}

type GetPlayerRequest struct {
	PlayerID uuid.UUID `json:"player_id"`
}

func NewGetPlayerHandler(logger logrus.FieldLogger, decoder app.Decoder, controller player.PlayerController) *GetPlayerHandler {
	h := &GetPlayerHandler{
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

func (h *GetPlayerHandler) GetMethod() string {
	return http.MethodGet
}

func (h *GetPlayerHandler) GetPath() string {
	return "/player"
}

func (h *GetPlayerHandler) Handle(r *http.Request) app.Response {
	req := GetPlayerRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("getting player")
	player, err := h.controller.Get(req.PlayerID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"player": player}).Info("successfully retrieved player")
	return app.NewStatusOK(player)
}
