package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetScoreHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller player.PlayerController
}

type GetScoreRequest struct {
	PlayerID uuid.UUID `json:"player_id"`
}

func NewGetScoreHandler(logger logrus.FieldLogger, decoder app.Decoder, controller player.PlayerController) *GetScoreHandler {
	h := &GetScoreHandler{
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

func (h *GetScoreHandler) GetMethod() string {
	return http.MethodGet
}

func (h *GetScoreHandler) GetPath() string {
	return "/score"
}

func (h *GetScoreHandler) Handle(r *http.Request) app.Response {
	req := GetScoreRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("getting player score")
	score, err := h.controller.Get(req.PlayerID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"score": score}).Info("successfully retrieved player score")
	return app.NewStatusOK(score)
}
