package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/environment"
	"pkg.dsb.dev/transport"
)

func TestRangeFromHeader(t *testing.T) {
	tt := []struct {
		Name          string
		Range         string
		ExpectedStart int64
		ExpectedEnd   int64
		ExpectsError  bool
	}{
		{
			Name:          "It should parse the start and end from the header",
			Range:         "bytes=0-9",
			ExpectedStart: 0,
			ExpectedEnd:   9,
		},
		{
			Name:          "It should return -1 if no end is specified",
			Range:         "bytes=0-",
			ExpectedStart: 0,
			ExpectedEnd:   -1,
		},
		{
			Name:         "It should return an error for an invalid header",
			Range:        "test",
			ExpectsError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.Header.Set("Range", tc.Range)

			start, end, err := transport.RangeFromHeader(r)
			if tc.ExpectsError {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.ExpectedStart, start)
			assert.Equal(t, tc.ExpectedEnd, end)
		})
	}
}

func TestRangeFromValues(t *testing.T) {
	tt := []struct {
		Name     string
		Expected string
		Start    int64
		End      int64
		Size     int64
	}{
		{
			Name:     "It should build a Content-Range header",
			Expected: "bytes 0-9/10",
			Start:    0,
			End:      9,
			Size:     10,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			actual := transport.RangeFromValues(tc.Start, tc.End, tc.Size)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}

func TestMethodNotAllowed(t *testing.T) {
	h := transport.MethodNotAllowed()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestNotFound(t *testing.T) {
	h := transport.NotFound()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestHTTP_Decode(t *testing.T) {
	tr := transport.HTTP{}

	// Valid body
	a := struct {
		Value string `json:"value" validate:"required"`
	}{
		Value: "test",
	}

	m, err := json.Marshal(a)
	assert.NoError(t, err)
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(m))
	assert.NoError(t, tr.Decode(r, &a))

	// Invalid body
	b := struct {
		Value string `json:"value" validate:"required"`
	}{}

	m, err = json.Marshal(b)
	assert.NoError(t, err)
	r = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(m))
	assert.Error(t, tr.Decode(r, &b))
}

func TestHTTP_Error(t *testing.T) {
	tr := transport.HTTP{}
	ctx := environment.NewContext()
	w := httptest.NewRecorder()

	tr.Error(ctx, w, http.StatusBadRequest, "hello")

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHTTP_ErrorWithStack(t *testing.T) {
	tr := transport.HTTP{}
	ctx := environment.NewContext()
	w := httptest.NewRecorder()

	tr.ErrorWithStack(ctx, w, http.StatusBadRequest, "hello")

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHTTP_Respond(t *testing.T) {
	tr := transport.HTTP{}
	ctx := environment.NewContext()
	w := httptest.NewRecorder()

	tr.Respond(ctx, w, nil, http.StatusCreated)

	assert.Equal(t, http.StatusCreated, w.Code)
}
