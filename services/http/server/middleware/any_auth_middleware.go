package middleware

import (
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
	app "github.com/byyjoww/leaderboard/services/http/server"
	"github.com/byyjoww/leaderboard/services/http/server/response"
)

var (
	_ app.Middleware = (*AnyAuthMiddleware)(nil)
)

type AnyAuthMiddleware struct {
	authorizers []Authorizer
}

type Authorizer interface {
	Authorize(w http.ResponseWriter, r *http.Request) bool
}

func NewAnyAuthMiddleware(authorizers ...Authorizer) *AnyAuthMiddleware {
	return &AnyAuthMiddleware{
		authorizers: authorizers,
	}
}

func (a *AnyAuthMiddleware) ServeNext(logger lib.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			for _, auth := range a.authorizers {
				if ok := auth.Authorize(w, r); ok {
					next.ServeHTTP(w, r)
					return
				}
			}
			app.Respond(logger, w, response.NewJsonUnauthorized())
		})
}
