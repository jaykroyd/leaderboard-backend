package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetLeaderboardHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller leaderboard.LeaderboardController
}

type GetLeaderboardRequest struct {
	LeaderboardID uuid.UUID `json:"leaderboard_id"`
}

func NewGetLeaderboardHandler(logger logrus.FieldLogger, decoder app.Decoder, controller leaderboard.LeaderboardController) *GetLeaderboardHandler {
	h := &GetLeaderboardHandler{
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

func (h *GetLeaderboardHandler) GetMethod() string {
	return http.MethodGet
}

func (h *GetLeaderboardHandler) GetPath() string {
	return "/leaderboard"
}

func (h *GetLeaderboardHandler) Handle(r *http.Request) app.Response {
	req := GetLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	h.logger.WithFields(logrus.Fields{"request": req}).Info("getting leaderboard")
	lb, err := h.controller.Get(req.LeaderboardID)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"leaderboard": lb}).Infof("successfully retrieved leaderboard")
	return app.NewStatusOK(lb)
}
