package auth

import (
	"fmt"
	"net/http"
)

type Basic struct {
	user string
	pass string
}

func NewBasic(user string, pass string) *Basic {
	return &Basic{
		user: user,
		pass: pass,
	}
}

func (a *Basic) SetAuth(req *http.Request) {
	req.SetBasicAuth(a.user, a.pass)
}

func (a *Basic) String() string {
	return fmt.Sprintf("%s:%s", a.user, a.pass)
}
