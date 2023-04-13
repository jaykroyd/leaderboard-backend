package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type HttpRequest struct {
	client      HTTPClient
	decoder     Decoder
	ctx         context.Context
	method      string
	url         string
	contentType string
	auth        Auth
	body        io.Reader
	query       map[string]string
	headers     map[string]string
}

func (r *HttpRequest) WithHeaders(headers map[string]string) *HttpRequest {
	for k, v := range headers {
		if _, ok := r.headers[k]; !ok {
			r.headers[k] = v
		}
	}
	return r
}

func (r *HttpRequest) WithQueryParams(params map[string]string) *HttpRequest {
	for k, v := range params {
		if _, ok := r.query[k]; !ok {
			r.query[k] = v
		}
	}
	return r
}

func (r *HttpRequest) WithJsonBody(body interface{}) *HttpRequest {
	b, err := json.Marshal(body)
	if err == nil {
		r.body = bytes.NewBuffer(b)
		r.contentType = ContentTypeApplicationJson
	}
	return r
}

func (r *HttpRequest) WithFormBody(body map[string]string) *HttpRequest {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	defer w.Close()

	for k, v := range body {
		w.WriteField(k, v)
	}

	r.contentType = w.FormDataContentType()
	r.body = b
	return r
}

func (r *HttpRequest) WithAuth(auth Auth) *HttpRequest {
	r.auth = auth
	return r
}

func (r *HttpRequest) WithContext(ctx context.Context) *HttpRequest {
	r.ctx = ctx
	return r
}

func (r *HttpRequest) Replace(replacements map[string]string) *HttpRequest {
	for k, v := range replacements {
		r.url = strings.Replace(r.url, k, v, 1)
	}
	return r
}

func (r *HttpRequest) Do() (*http.Response, error) {
	uri, err := r.buildUrl()
	if err != nil || uri.String() == "" {
		return nil, errors.Wrap(ErrBuildingUrl, err.Error())
	}

	if r.body == nil {
		r.body = bytes.NewBuffer([]byte{})
	}

	if r.ctx == nil {
		r.ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(r.ctx, r.method, uri.String(), r.body)
	if err != nil {
		return nil, err
	}

	r.buildHeaders(req)
	return r.client.Do(req)
}

func (r *HttpRequest) DoAndUnmarshal(result interface{}) (int, error) {
	resp, err := r.Do()
	if err != nil {
		return 0, err
	}

	if isErrorCode(resp.StatusCode) {
		return r.handleError(resp)
	}

	return r.handleResponse(result, resp)
}

func (r *HttpRequest) handleResponse(result interface{}, resp *http.Response) (int, error) {
	if hasContent(resp.StatusCode, resp.ContentLength) {
		if err := r.decoder.DecodeResponse(resp, result); err != nil {
			return resp.StatusCode, errors.Wrap(ErrDecodingResponse, err.Error())
		}
	}
	return resp.StatusCode, nil
}

func (r *HttpRequest) handleError(resp *http.Response) (int, error) {
	if !hasContent(resp.StatusCode, resp.ContentLength) {
		return resp.StatusCode, errors.Wrap(ErrStatusNotOk, "no content")
	}

	errContent, err := buildErrorContent(resp)
	if err != nil {
		return resp.StatusCode, errors.Wrap(ErrDecodingResponse, err.Error())
	}

	return resp.StatusCode, errors.Wrap(ErrStatusNotOk, errContent)
}

func (r *HttpRequest) buildUrl() (*url.URL, error) {
	uri, err := url.Parse(r.url)
	if err != nil {
		return nil, errors.Wrap(ErrParsingUri, r.url)
	}

	query := uri.Query()
	for k, v := range r.query {
		query.Set(k, v)
	}

	uri.RawQuery = query.Encode()
	return uri, nil
}

func (r *HttpRequest) buildHeaders(req *http.Request) {
	if r.contentType != "" {
		req.Header.Set("Content-Type", r.contentType)
	} else {
		req.Header.Set("Content-Length", "0")
	}

	if r.auth != nil {
		r.auth.SetAuth(req)
	}

	for k, v := range r.headers {
		if _, ok := req.Header[k]; !ok {
			req.Header.Set(k, v)
		}
	}
}

func (r *HttpRequest) String() string {
	return fmt.Sprintf("method: %s | url: %s | content-type: %s | headers: %s | body: %v | query: %s",
		r.method, r.url, r.contentType, r.headers, r.body, r.query)
}

func buildErrorContent(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	errContent := []string{}
	if resp.Header != nil {
		jsonHeader, err := json.Marshal(resp.Header)
		if err != nil {
			return "", fmt.Errorf("failed to marshal response header: %w", err)
		}
		errContent = append(errContent, fmt.Sprintf("Response Headers: %s", string(jsonHeader)))
	}
	if len(data) > 0 {
		errContent = append(errContent, fmt.Sprintf("Response Body: %s", string(data)))
	}
	if resp.Request != nil && resp.Request.URL != nil {
		errContent = append(errContent, fmt.Sprintf("Request Url: %s", resp.Request.URL.String()))
	}

	return strings.Join(errContent, " | "), nil
}

func hasContent(statusCode int, contentLength int64) bool {
	return statusCode != 204 || contentLength > 0
}

func isErrorCode(statusCode int) bool {
	return statusCode >= 300
}
