package client

import "net/http"

type Auth interface {
	SetAuth(req *http.Request)
	String() string
}
