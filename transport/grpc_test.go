package transport_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pkg.dsb.dev/transport"
	v1 "pkg.dsb.dev/transport/v1"
)

func TestGRPC_ErrorWithStack(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	gRPC := transport.GRPC{}
	err := gRPC.ErrorWithStack(ctx, codes.Internal, "this is a test")

	st, ok := status.FromError(err)

	assert.True(t, ok)
	assert.NotNil(t, st)
	assert.Equal(t, "this is a test", st.Message())
	assert.Equal(t, codes.Internal, st.Code())
	if assert.Len(t, st.Details(), 1) {
		dt := st.Details()[0]

		ed, ok := dt.(*v1.GRPCErrorDetails)

		assert.True(t, ok)
		assert.NotNil(t, ed)
		assert.NotEmpty(t, ed.StackTrace)
	}
}

func TestGRPC_Error(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	gRPC := transport.GRPC{}
	err := gRPC.Error(ctx, codes.Internal, "this is a test")

	st, ok := status.FromError(err)

	assert.True(t, ok)
	assert.NotNil(t, st)
	assert.Equal(t, "this is a test", st.Message())
	assert.Equal(t, codes.Internal, st.Code())
}
