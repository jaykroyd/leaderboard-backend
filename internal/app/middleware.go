package app

import "net/http"

type Middleware interface {
	ServeHTTP(next http.Handler) http.Handler
}
