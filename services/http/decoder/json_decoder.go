package decoder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	ContentTypeApplicationJson string = "application/json"
)

type JsonDecoder struct {
	contentType string
}

func NewJsonDecoder() *JsonDecoder {
	return &JsonDecoder{
		contentType: ContentTypeApplicationJson,
	}
}

func (d *JsonDecoder) HasContentType(contentType string) bool {
	return strings.HasPrefix(contentType, d.contentType)
}

func (d *JsonDecoder) DecodeRequest(r *http.Request, contentType string, response interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	err = json.Unmarshal(data, response)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	return nil
}

func (d *JsonDecoder) DecodeResponse(r *http.Response, contentType string, response interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	err = json.Unmarshal(data, response)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	return nil
}
