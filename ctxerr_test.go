package ctxerr

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestCommonUsage(t *testing.T) {
	err := errors.New("something happened")
	ctx := New(err).WithField("testfield", "testvalue")
	assert.Error(t, *ctx)
	assert.Equal(t, err, errors.Unwrap(ctx))
	assert.Nil(t, ctx.GetField("test"))
	assert.Equal(t, ctx.GetField("testfield"), "testvalue")
	ctx2 := ctx.WithField("testfield", "some other value")
	assert.NotEqual(t, ctx.GetField("testfield"), ctx2.GetField("testfield"))
	ctx3 := ctx.WithField("some other field", "some another value")
	assert.Equal(t, ctx.GetField("testfield"), ctx3.GetField("testfield"))
}
