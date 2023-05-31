package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type GetLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.Provider
}

type GetLeaderboardResponse struct {
	Leaderboard *leaderboard.Leaderboard `json:"leaderboard"`
}

func NewGetLeaderboardHandler(decoder server.Decoder, controller leaderboard.Provider) *GetLeaderboardHandler {
	return &GetLeaderboardHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *GetLeaderboardHandler) GetMethod() string {
	return http.MethodGet
}

func (h *GetLeaderboardHandler) GetPath() string {
	return "/leaderboards/{leaderboard_id}"
}

func (h *GetLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("getting leaderboard")

	var (
		vars          = mux.Vars(r)
		leaderboardId = vars["leaderboard_id"]
	)

	logger = logger.WithFields(logging.Fields{
		"leaderboard_id": leaderboardId,
	})

	leaderboardUUID, err := uuid.Parse(leaderboardId)
	if err != nil {
		logger.WithError(err).Error("failed to parse leaderboard id")
		return NewBadRequest(errors.Wrap(err, "failed to parse leaderboard id"))
	}

	leaderboard, err := h.controller.Get(r.Context(), leaderboardUUID)
	if err != nil {
		logger.WithError(err).Error("failed to retrieve leaderboard")
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{
		"leaderboard": leaderboard,
	}).Infof("successfully retrieved leaderboard")

	return NewStatusOK(GetLeaderboardResponse{
		Leaderboard: leaderboard,
	})
}
