package handler

import (
	app "github.com/byyjoww/leaderboard/services/http/server"
	"github.com/byyjoww/leaderboard/services/http/server/response"
)

func NewBadRequest(err error) app.Response {
	return response.NewJsonBadRequest(err)
}

func NewInternalServerError(err error) app.Response {
	return response.NewJsonInternalServerError(err)
}

func NewStatusOK(content interface{}) app.Response {
	return response.NewJsonStatusOK(content)
}
