package blob_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/storage/blob"
)

func TestBucket_Iterate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	bkt, err := blob.OpenBucket(ctx, "mem://")
	assert.NoError(t, err)
	assert.NotNil(t, bkt)

	tree := []string{
		"a",
		"/test/b",
		"/test/test/c",
		"/test/test/nothing/test/d",
	}

	// Populate some random files in the bucket
	for _, key := range tree {
		wr, err := bkt.NewWriter(ctx, key)
		assert.NoError(t, err)
		assert.NotNil(t, wr)

		buf := bytes.NewBuffer([]byte("hello world"))
		_, err = io.Copy(wr, buf)
		assert.NoError(t, err)
		assert.NoError(t, wr.Close())
	}

	// Iterate over the files
	count := 0
	assert.NoError(t, bkt.Iterate(ctx, func(ctx context.Context, item blob.Blob) error {
		assert.NotZero(t, item)
		assert.Contains(t, tree, item.Key)
		assert.NotZero(t, item.ModTime)
		assert.EqualValues(t, len("hello world"), item.Size)
		count++
		return nil
	}))

	assert.Equal(t, len(tree), count)
	assert.NoError(t, bkt.Close())
}
