package middleware_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/middleware"
	"pkg.dsb.dev/transport"
)

func TestPanic(t *testing.T) {
	fn := middleware.Panic()
	hnd := fn.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	}))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	assert.NotPanics(t, func() {
		hnd.ServeHTTP(w, r)
	})

	var body transport.Error

	assert.NoError(t, json.NewDecoder(w.Body).Decode(&body))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotEmpty(t, body.Message)
	assert.NotEmpty(t, body.Stack)
}
