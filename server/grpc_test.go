package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/server"
)

func TestServeGRPC(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	finished := make(chan struct{})
	go func() {
		assert.NoError(t, server.ServeGRPC(ctx))
		finished <- struct{}{}
	}()

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	cancel()

	select {
	case <-finished:
		return
	case <-ticker.C:
		assert.Fail(t, "server did not shut down after 10 seconds")
	}
}
