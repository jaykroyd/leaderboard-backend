package decoder

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

const (
	ContentTypeFormUrlEncoded string = "application/x-www-form-urlencoded"
)

type FormURLEncodedDecoder struct {
	contentType string
	decoder     *schema.Decoder
}

func NewFormURLEncodedDecoder() *FormURLEncodedDecoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return &FormURLEncodedDecoder{
		contentType: ContentTypeFormUrlEncoded,
		decoder:     decoder,
	}
}

func (d *FormURLEncodedDecoder) HasContentType(contentType string) bool {
	return strings.HasPrefix(contentType, d.contentType)
}

func (d *FormURLEncodedDecoder) DecodeRequest(r *http.Request, contentType string, response interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}

	formUnescaped, err := unescape(r.Form)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}

	err = d.decoder.Decode(response, formUnescaped)
	if err != nil {
		return errors.Wrap(ErrDecodingContent, err.Error())
	}
	return nil
}

func (d *FormURLEncodedDecoder) DecodeResponse(r *http.Response, contentType string, response interface{}) error {
	return nil
}

func unescape(values url.Values) (url.Values, error) {
	queryUnescaped, err := url.QueryUnescape(values.Encode())
	if err != nil {
		return url.Values{}, err
	}

	// url.QueryUnescape doesn't seem to replace \u00a0
	unescapedNoBreakSpace := strings.ReplaceAll(queryUnescaped, "\u00a0", " ")

	parsedValues, err := url.ParseQuery(unescapedNoBreakSpace)
	if err != nil {
		return url.Values{}, err
	}

	return parsedValues, nil
}
