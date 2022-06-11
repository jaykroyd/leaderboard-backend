package decoder

import "net/http"

type ContentTypeDecoder interface {
	HasContentType(string) bool
	DecodeRequest(r *http.Request, contentType string, response interface{}) error
	DecodeResponse(r *http.Response, contentType string, response interface{}) error
}
