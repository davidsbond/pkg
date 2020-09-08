package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/middleware"
	"pkg.dsb.dev/requestid"
)

func TestRequestID(t *testing.T) {
	fn := middleware.RequestID()
	hnd := fn.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		ctxID := requestid.FromContext(r.Context())

		assert.Equal(t, id, ctxID)
		assert.NotEmpty(t, id)
		assert.NotEmpty(t, ctxID)
	}))

	// Request id is generated by the server
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	hnd.ServeHTTP(w, r)
	actual := w.Header().Get("X-Request-ID")
	assert.NotEmpty(t, actual)

	// Request id is generated by the client
	expected := "my-id"
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/", nil)
	r.Header.Set("X-Request-ID", expected)

	hnd.ServeHTTP(w, r)
	actual = w.Header().Get("X-Request-ID")
	assert.Equal(t, expected, actual)
}
