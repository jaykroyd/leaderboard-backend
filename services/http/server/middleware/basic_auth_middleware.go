package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
	app "github.com/byyjoww/leaderboard/services/http/server"
	"github.com/byyjoww/leaderboard/services/http/server/response"
)

var (
	_ app.Middleware = (*BasicAuthMiddleware)(nil)
	_ Authorizer     = (*BasicAuthMiddleware)(nil)
)

type BasicAuthMiddleware struct {
	user    string
	pass    string
	enabled bool
}

func NewBasicAuthMiddleware(user string, pass string, enabled bool) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{
		user:    user,
		pass:    pass,
		enabled: enabled,
	}
}

func (b *BasicAuthMiddleware) ServeNext(logger lib.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok := b.Authorize(w, r); ok {
			next.ServeHTTP(w, r)
			return
		}
		app.Respond(logger, w, response.NewJsonUnauthorized())
	})
}

func (b *BasicAuthMiddleware) Authorize(vw http.ResponseWriter, r *http.Request) bool {
	if !b.enabled {
		// logger.Debug("authentication is disabled")
		return true
	}

	user, pass, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(user))
		passwordHash := sha256.Sum256([]byte(pass))
		expectedUsernameHash := sha256.Sum256([]byte(b.user))
		expectedPasswordHash := sha256.Sum256([]byte(b.pass))
		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
		return usernameMatch && passwordMatch
	}

	return false
}
