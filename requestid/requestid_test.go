package requestid_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/requestid"
)

func TestExtract(t *testing.T) {
	t.Parallel()
	const key = "X-Request-ID"

	tt := []struct {
		Name          string
		RequestHeader http.Header
	}{
		{
			Name: "It should use an existing request id",
			RequestHeader: map[string][]string{
				key: {"test-id"},
			},
		},
		{
			Name:          "It should generate an identifier",
			RequestHeader: make(http.Header),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			respHeader := make(http.Header)
			_ = requestid.Extract(context.Background(), tc.RequestHeader, respHeader)

			assert.NotEmpty(t, tc.RequestHeader.Get(key))
			assert.NotEmpty(t, respHeader.Get(key))
			assert.EqualValues(t, tc.RequestHeader.Get(key), respHeader.Get(key))
		})
	}
}
