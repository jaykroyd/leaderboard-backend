package server

import (
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
)

type Middleware interface {
	ServeNext(logger lib.Logger, next http.Handler) http.Handler
}
