package testgrpc_test

import (
	"context"
	"testing"

	"github.com/lesomnus/proto-orm/internal/example/library"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	t.Run("default", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, library.LockerById(c.v_locker.GetId()))
		x.NoError(err)
		x.Equal(c.v_locker.GetId(), v.GetId())
		x.Equal(c.v_locker.GetNumber(), v.GetNumber())
		// Direct edges are loaded.
		x.True(v.HasOwner())

		w := v.GetOwner()
		x.Equal(c.v_member.GetId(), w.GetId())
		// Transient edges are not loaded.
		x.False(w.HasParent())
	}))

	// TODO: specific fields, nested select
}
