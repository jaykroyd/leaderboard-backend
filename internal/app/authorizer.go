package app

import "net/http"

type Authorizer interface {
	Authorize(r *http.Request) bool
}
