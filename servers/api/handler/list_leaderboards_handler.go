package handler

import (
	"fmt"
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/internal/app"
	"github.com/sirupsen/logrus"
)

type ListLeaderboardHandler struct {
	logger     logrus.FieldLogger
	decoder    app.Decoder
	controller leaderboard.LeaderboardController
}

type ListLeaderboardRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func NewListLeaderboardsHandler(logger logrus.FieldLogger, decoder app.Decoder, controller leaderboard.LeaderboardController) *ListLeaderboardHandler {
	h := &ListLeaderboardHandler{
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

func (h *ListLeaderboardHandler) GetMethod() string {
	return http.MethodGet
}

func (h *ListLeaderboardHandler) GetPath() string {
	return "/leaderboard/list"
}

func (h *ListLeaderboardHandler) Handle(r *http.Request) app.Response {
	req := ListLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		h.logger.WithError(err).Error("error decoding request")
		return app.NewBadRequest(err)
	}

	lbs, err := h.controller.List(req.Limit, req.Offset)
	if err != nil {
		return app.NewInternalServerError(err)
	}

	h.logger.WithFields(logrus.Fields{"leaderboards": lbs, "amount": len(lbs)}).Info("successfully retrieved leaderboards")
	return app.NewStatusOK(lbs)
}
