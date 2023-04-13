package middleware

import (
	"fmt"
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
	app "github.com/byyjoww/leaderboard/services/http/server"
	"github.com/byyjoww/leaderboard/services/http/server/response"
)

type RoleIdentifier interface {
	Identify(user User) Role
}

type Role interface {
	IsMetBy(role Role) bool
}

type User interface {
	ID() string
}

type UserDecoder interface {
	Decode(r *http.Request) (User, error)
}

var (
	_ app.Middleware = (*RoleRequirementMiddleware)(nil)
)

type RoleRequirementMiddleware struct {
	logger       lib.Logger
	decoder      UserDecoder
	identifier   RoleIdentifier
	requiredRole Role
}

func NewRoleRequirementMiddleware(logger lib.Logger, decoder UserDecoder, identifier RoleIdentifier, requiredRole Role) *RoleRequirementMiddleware {
	return &RoleRequirementMiddleware{
		logger:       logger.WithField("requirement", fmt.Sprintf("%T", requiredRole)),
		decoder:      decoder,
		identifier:   identifier,
		requiredRole: requiredRole,
	}
}

func (b *RoleRequirementMiddleware) ServeNext(logger lib.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := b.decoder.Decode(r)
		if err != nil {
			b.logger.WithError(err).Error("failed to decode user role")
			app.Respond(logger, w, response.NewJsonBadRequest(err))
			return
		}

		role := b.identifier.Identify(user)
		if b.requiredRole.IsMetBy(role) {
			next.ServeHTTP(w, r)
			return
		}
		app.Respond(logger, w, response.NewJsonUnauthorized())
	})
}
