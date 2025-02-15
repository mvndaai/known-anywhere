package server_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-anywhere/internal/router/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrimSpacesMiddleware(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?qa=qb", http.NoBody)
	req.Header.Add("ha", "  hb  ")
	params := req.URL.Query()
	params.Set("qc", " qd ")
	req.URL.RawQuery = params.Encode()

	testHandler := func(_ http.ResponseWriter, r *http.Request) {
		expectedHeaders := http.Header{"Ha": []string{"hb"}}
		assert.Equal(t, expectedHeaders, r.Header)

		expectedParams := url.Values{"qa": []string{"qb"}, "qc": []string{"qd"}}
		assert.Equal(t, expectedParams, r.URL.Query())
	}

	m := server.TrimSpaces(testHandler)
	m.ServeHTTP(httptest.NewRecorder(), req)
}

func TestLogHeadersAndParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?queryKey=query_value&queryKey=query_value", http.NoBody)
	req.Header.Add("Header-Key", "header_value")
	req.Header.Add("Authorization", "bearer abc")

	// Test clearDuplicates func
	req.Header.Add("DupToClear", "1")
	req.Header.Add("DupToClear", "1")
	req.Header.Add("NonOrderedDup", "1")
	req.Header.Add("NonOrderedDup", "2")
	req.Header.Add("NonOrderedDup", "1")

	testHandler := func(_ http.ResponseWriter, r *http.Request) {
		f := ctxerr.Fields(r.Context())

		headers := f["headers"]
		require.NotNil(t, headers)
		h, err := headers.(http.Header)
		require.NotNil(t, err)
		assert.Equal(t, "header_value", h.Get("Header-Key"))
		assert.Equal(t, "", h.Get("Authorization"))
		assert.Len(t, h, 3)
		assert.EqualValues(t, []string{"1"}, h.Values("DupToClear"))
		assert.EqualValues(t, []string{"1", "2"}, h.Values("NonOrderedDup"))

		parameters := f["parameters"]
		require.NotNil(t, parameters)
		p, err := parameters.(url.Values)
		require.NotNil(t, err)
		assert.Equal(t, "query_value", p.Get("queryKey"))
		assert.EqualValues(t, []string{"query_value"}, p["queryKey"])
		assert.Len(t, p, 1)
	}

	m := server.LogHeadersAndParams(testHandler)
	m.ServeHTTP(httptest.NewRecorder(), req)
}
