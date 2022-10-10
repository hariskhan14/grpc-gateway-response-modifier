package service_test

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	assert := require.New(t)
	client := &http.Client{}
	url := testCtx.grpcServer.URL + "/user/123"

	t.Run("no version", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Version", "2.0")

		resp, err := client.Do(req)
		assert.NoError(err)
		assert.Equal(http.StatusNotAcceptable, resp.StatusCode)
		assert.Equal("application/json", resp.Header.Get("content-type"))
	})

	t.Run("unsupported version", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Version", "")

		resp, err := client.Do(req)
		assert.NoError(err)
		assert.Equal(http.StatusNotAcceptable, resp.StatusCode)
		assert.Equal("application/json", resp.Header.Get("content-type"))
	})

	t.Run("supported version", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Version", "1.0")

		resp, err := client.Do(req)
		assert.NoError(err)
		assert.Equal(http.StatusOK, resp.StatusCode)
		assert.Equal("application/json", resp.Header.Get("content-type"))
	})
}
