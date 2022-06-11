package app

import "net/http"

type Decoder interface {
	DecodeRequest(req *http.Request, i interface{}) error
}
