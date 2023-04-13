package auth

import (
	"net/http"
)

type Auth struct {
	authorization string
}

func New(authorization string) *Auth {
	return &Auth{
		authorization: authorization,
	}
}

func (a *Auth) SetAuth(req *http.Request) {
	req.Header.Set("Authorization", a.authorization)
}

func (a *Auth) String() string {
	return a.authorization
}
