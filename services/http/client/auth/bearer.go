package auth

import (
	"fmt"
	"net/http"
)

type BearerAuth struct {
	token string
}

func NewBearer(token string) *BearerAuth {
	return &BearerAuth{
		token: token,
	}
}

func (a *BearerAuth) SetAuth(req *http.Request) {
	req.Header.Set("Authorization", a.getBearer())
}

func (a *BearerAuth) String() string {
	return a.getBearer()
}

func (a *BearerAuth) getBearer() string {
	return fmt.Sprintf("Bearer: %s", a.token)
}
