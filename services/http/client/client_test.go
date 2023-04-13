package client_test

import (
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/byyjoww/leaderboard/services/http/client"
	"github.com/byyjoww/leaderboard/services/http/client/auth"
	"github.com/byyjoww/leaderboard/services/http/decoder"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Parallel()

	var (
		httpClient = http.DefaultClient
		decoder    = decoder.New()
	)

	startTestServer()

	client := client.New(httpClient, decoder).
		WithAuth(auth.NewBasic("user", "pass")).
		WithBaseUrl("http://0.0.0.0:6000")

	t.Run("test_client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		code, err := client.Get("/healthcheck").DoAndUnmarshal(nil)
		assert.Less(t, code, 300)
		assert.Nil(t, err)
	})

	t.Run("test_application_json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		resp := &map[string]interface{}{}
		code, err := client.Get("/applicationJson").DoAndUnmarshal(resp)
		r := *resp

		assert.Less(t, code, 300)
		assert.Nil(t, err)
		assert.Equal(t, "value", r["key"])
	})

	t.Run("test_multipart_form", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		resp := &map[string]string{}
		code, err := client.Get("/multipartFormData").DoAndUnmarshal(resp)
		r := *resp

		assert.Less(t, code, 300)
		assert.Nil(t, err)
		assert.Equal(t, "value", r["key"])
	})
}

func startTestServer() {
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	r.HandleFunc("/applicationJson", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key":"value"}`))
	}).Methods("GET")

	r.HandleFunc("/multipartFormData", func(w http.ResponseWriter, r *http.Request) {
		mw := multipart.NewWriter(w)
		defer mw.Close()

		w.Header().Set("Content-Type", mw.FormDataContentType())
		values := map[string]string{"key": "value"}
		for k, v := range values {
			mw.WriteField(k, v)
		}
	}).Methods("GET")

	go http.ListenAndServe("0.0.0.0:6000", r)
}
