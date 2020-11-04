package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

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

func TestDialGRPC(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		assert.NoError(t, server.ServeGRPC(ctx))
	}()

	conn, err := server.DialGRPC(ctx, "localhost:5000", grpc.WithInsecure())
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	assert.NoError(t, conn.Close())
	cancel()
}
