package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type RoleRequirementMiddleware struct {
	logger       logrus.FieldLogger
	identifier   RoleIdentifier
	requiredRole Role
}

func NewRoleRequirementMiddleware(logger logrus.FieldLogger, identifier RoleIdentifier, requiredRole Role) *RoleRequirementMiddleware {
	return &RoleRequirementMiddleware{
		logger:       logger.WithField("requirement", fmt.Sprintf("%T", requiredRole)),
		identifier:   identifier,
		requiredRole: requiredRole,
	}
}

func (b *RoleRequirementMiddleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := b.identifier.Identify(&HubUser{
			user:  r.Header.Get("x-forwarded-user"),
			email: r.Header.Get("x-forwarded-email"),
		})

		if b.requiredRole.IsMetBy(role) {
			next.ServeHTTP(w, r)
			return
		}

		msg := "Forbidden"
		b.logger.Debug(msg)
		b, _ := json.Marshal(msg)
		// w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(b)
	})
}
