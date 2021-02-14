package ftp_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/environment"
	"pkg.dsb.dev/testutil"
)

func TestConn_Walk(t *testing.T) {
	t.Parallel()

	ctx := environment.NewContext()
	conn := testutil.WithFTPServer(t)

	assert.NoError(t, conn.Walk(ctx, "/", func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)
		return nil
	}))
}
