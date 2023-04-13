package client

import "net/http"

const (
	ContentTypeApplicationJson string = "application/json"
)

type Decoder interface {
	DecodeResponse(r *http.Response, i interface{}) error
}
