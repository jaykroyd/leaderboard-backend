package app

import (
	"net/http"
)

type Handler interface {
	GetMethod() string
	GetPath() string
	Handle(r *http.Request) Response
}
