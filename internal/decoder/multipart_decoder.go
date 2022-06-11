package decoder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	ContentTypeMultipartFormData string = "multipart/form-data"
)

type MultipartFormDataDecoder struct {
	contentType string
}

func NewMultipartFormDataDecoder() *MultipartFormDataDecoder {
	return &MultipartFormDataDecoder{
		contentType: ContentTypeMultipartFormData,
	}
}

func (d *MultipartFormDataDecoder) HasContentType(contentType string) bool {
	return strings.HasPrefix(contentType, d.contentType)
}

func (d *MultipartFormDataDecoder) DecodeRequest(r *http.Request, contentType string, response interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	return d.decode(data, contentType, response)
}

func (d *MultipartFormDataDecoder) DecodeResponse(r *http.Response, contentType string, response interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	return d.decode(data, contentType, response)
}

func (d *MultipartFormDataDecoder) decode(data []byte, contentType string, response interface{}) error {
	_, params, _ := mime.ParseMediaType(contentType)
	buf := bytes.NewBuffer(data)
	mr := multipart.NewReader(buf, params["boundary"])

	form, err := mr.ReadForm(0)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}

	newForm := map[string]string{}
	for k, v := range form.Value {
		if len(v) > 0 {
			newForm[k] = v[0]
		}
	}

	data, err = json.Marshal(newForm)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}

	return nil
}
