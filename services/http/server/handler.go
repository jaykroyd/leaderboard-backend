package server

import (
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
)

type Handler interface {
	GetMethod() string
	GetPath() string
	Handle(logger lib.Logger, r *http.Request) Response
}

type HTTPHandler interface {
	GetMethod() string
	GetPath() string
	ServeHTTP(logger lib.Logger, w http.ResponseWriter, r *http.Request)
}
