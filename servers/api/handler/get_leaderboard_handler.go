package handler

import (
	"net/http"

	"github.com/byyjoww/leaderboard/bll/leaderboard"
	leaderboardDal "github.com/byyjoww/leaderboard/dal/leaderboard"
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
	Leaderboard *leaderboardDal.Leaderboard `json:"leaderboard"`
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
	return "/leaderboards/{id}"
}

func (h *GetLeaderboardHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	vars := mux.Vars(r)
	leaderboardIdString := vars["id"]

	logger.WithFields(logging.Fields{
		"request": leaderboardIdString,
	}).Info("getting leaderboard")

	leaderboardId, err := uuid.Parse(leaderboardIdString)
	if err != nil {
		if err != nil {
			return NewBadRequest(errors.Wrap(err, "failed to parse leaderboard id"))
		}
	}

	leaderboard, err := h.controller.Get(leaderboardId)
	if err != nil {
		return NewInternalServerError(err)
	}

	logger.WithFields(logging.Fields{"leaderboard": leaderboard}).Infof("successfully retrieved leaderboard")
	return NewStatusOK(GetLeaderboardResponse{
		Leaderboard: leaderboard,
	})
}
