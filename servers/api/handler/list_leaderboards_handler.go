package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
)

type ListLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.Provider
}

type ListLeaderboardsRequest struct {
	Mode   int `json:"mode"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListLeaderboardsResponse struct {
	Leaderboards []*leaderboard.Leaderboard `json:"leaderboards"`
}

func NewListLeaderboardsHandler(decoder server.Decoder, controller leaderboard.Provider) *ListLeaderboardHandler {
	return &ListLeaderboardHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *ListLeaderboardHandler) GetMethod() string {
	return http.MethodGet
}

func (h *ListLeaderboardHandler) GetPath() string {
	return "/leaderboards"
}

func (h *ListLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("listing leaderboards")

	req := ListLeaderboardsRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger = logger.WithFields(logging.Fields{
		"request": req,
	})

	hasMode := r.URL.Query().Has("mode")
	var leaderboards []*leaderboard.Leaderboard
	var err error

	if hasMode {
		leaderboards, err = h.controller.ListByMode(r.Context(), req.Mode, req.Limit, req.Offset)
	} else {
		leaderboards, err = h.controller.List(r.Context(), req.Limit, req.Offset)
	}

	if err != nil {
		logger.WithError(err).Error("failed to retrieve leaderboards")
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{
		"leaderboards": leaderboards,
		"amount":       len(leaderboards)},
	).Info("successfully retrieved leaderboards")

	return NewStatusOK(ListLeaderboardsResponse{
		Leaderboards: leaderboards,
	})
}
