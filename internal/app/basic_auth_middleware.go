package app

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type BasicAuthMiddleware struct {
	logger  logrus.FieldLogger
	user    string
	pass    string
	enabled bool
}

func NewBasicAuthMiddleware(logger logrus.FieldLogger, user string, pass string, enabled bool) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{
		logger:  logger.WithField("authorizer", "basicAuth"),
		user:    user,
		pass:    pass,
		enabled: enabled,
	}
}

func (b *BasicAuthMiddleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok := b.authorize(r); ok {
			next.ServeHTTP(w, r)
			return
		}

		msg := "Unauthorized"
		b.logger.Debug(msg)
		b, _ := json.Marshal(msg)
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(b)
	})
}

func (b *BasicAuthMiddleware) authorize(r *http.Request) bool {
	if !b.enabled {
		b.logger.Info("authentication is disabled")
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
