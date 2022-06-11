package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/player"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateScoreHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller player.PlayerController
}

type AddScoreRequest struct {
	PlayerID uuid.UUID `json:"player_id"`
	Amount   int       `json:"amount"`
}

func NewUpdateScoreHandler(logger logrus.FieldLogger, decoder app.Decoder, controller player.PlayerController) *UpdateScoreHandler {
	h := &UpdateScoreHandler{
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

func (h *UpdateScoreHandler) GetMethod() string {
	return http.MethodPatch
}

func (h *UpdateScoreHandler) GetPath() string {
	return "/score"
}

func (h *UpdateScoreHandler) Handle(r *http.Request) app.Response {
	req := AddScoreRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("updating player score")
	score, err := h.controller.UpdateScore(req.PlayerID, req.Amount)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"delta": req.Amount, "after": score}).Info("score updated successfully")
	return app.NewStatusOK(score)
}
