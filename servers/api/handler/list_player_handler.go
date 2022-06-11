package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ListPlayerHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller player.PlayerController
}

type ListPlayerRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
	Limit         int       `json:"limit"`
}

func NewListPlayerHandler(logger logrus.FieldLogger, decoder app.Decoder, controller player.PlayerController) *ListPlayerHandler {
	h := &ListPlayerHandler{
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

func (h *ListPlayerHandler) GetMethod() string {
	return http.MethodGet
}

func (h *ListPlayerHandler) GetPath() string {
	return "/player/list"
}

func (h *ListPlayerHandler) Handle(r *http.Request) app.Response {
	req := ListPlayerRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("getting player")
	players, err := h.controller.List(req.LeaderboardID, req.Limit, 0)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"players": players}).Info("successfully retrieved player")
	return app.NewStatusOK(players)
}
