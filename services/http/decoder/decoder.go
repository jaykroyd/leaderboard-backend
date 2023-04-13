package decoder

import (
	"net/http"
	"net/textproto"

	"github.com/pkg/errors"
)

var (
	ErrNoContentTypeAvailable = errors.New("no decoders available for this content type")
	ErrDecodingContent        = errors.New("decoding content")
)

type DecoderImpl struct {
	decoders []ContentTypeDecoder
}

func New() *DecoderImpl {
	return &DecoderImpl{
		decoders: []ContentTypeDecoder{
			NewFormURLEncodedDecoder(),
			NewJsonDecoder(),
			NewMultipartFormDataDecoder(),
		},
	}
}

func NewCustom(decoders ...ContentTypeDecoder) *DecoderImpl {
	return &DecoderImpl{
		decoders: decoders,
	}
}

// Decode clones and decodes an http request into a byte array, while re-populating the original request so it can be subsequently read by a handler.
func (d *DecoderImpl) DecodeRequest(r *http.Request, i interface{}) error {
	contentTypes := r.Header.Values(textproto.CanonicalMIMEHeaderKey("Content-Type"))
	if len(contentTypes) < 1 {
		// No content type provided in the request
		if r.Method == http.MethodGet {
			contentTypes = append(contentTypes, ContentTypeFormUrlEncoded)
		} else {
			return nil
		}
	}

	contentType := contentTypes[0]
	for _, decoder := range d.decoders {
		if decoder.HasContentType(contentType) {
			return decoder.DecodeRequest(r, contentType, i)
		}
	}

	return errors.Wrap(ErrNoContentTypeAvailable, contentType)
}

// Decode clones and decodes an http response into a byte array.
func (d *DecoderImpl) DecodeResponse(r *http.Response, i interface{}) error {
	contentTypes := r.Header.Values(textproto.CanonicalMIMEHeaderKey("Content-Type"))
	if len(contentTypes) < 1 {
		return nil
	}

	contentType := contentTypes[0]
	for _, decoder := range d.decoders {
		if decoder.HasContentType(contentType) {
			return decoder.DecodeResponse(r, contentType, i)
		}
	}

	return errors.Wrap(ErrNoContentTypeAvailable, contentType)
}
