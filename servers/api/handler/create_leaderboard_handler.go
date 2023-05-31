package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	"github.com/byyjoww/leaderboard/logging"
	app "github.com/byyjoww/leaderboard/services/http"
	"github.com/byyjoww/leaderboard/services/http/server"
)

type CreateLeaderboardHandler struct {
	decoder    server.Decoder
	controller leaderboard.Creator
}

type CreateLeaderboardRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Mode     int    `json:"mode"`
}

type CreateLeaderboardResponse struct {
	Leaderboard *leaderboard.Leaderboard `json:"leaderboard"`
}

func NewCreateLeaderboardHandler(decoder server.Decoder, controller leaderboard.Creator) *CreateLeaderboardHandler {
	return &CreateLeaderboardHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *CreateLeaderboardHandler) GetMethod() string {
	return http.MethodPost
}

func (h *CreateLeaderboardHandler) GetPath() string {
	return "/leaderboards"
}

func (h *CreateLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	logger.Info("creating leaderboard")

	req := CreateLeaderboardRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return NewBadRequest(err)
	}

	logger = logger.WithFields(logging.Fields{
		"request": req,
	})

	leaderboard, err := h.controller.Create(r.Context(), req.Name, req.Capacity, req.Mode)
	if err != nil {
		logger.WithError(err).Error("failed to create leaderboard")
		return NewInternalServerError(err)
	}

	logger.Info("new leaderboard created successfully")
	return NewStatusOK(CreateLeaderboardResponse{
		Leaderboard: leaderboard,
	})
}
