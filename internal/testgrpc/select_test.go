package testgrpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/lesomnus/proto-orm/internal/example/library"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSelect(t *testing.T) {
	v_t := true

	t.Run("default to all fields with edge IDs", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, c.v_locker.Pick())
		x.NoError(err)
		x.Equal(c.v_locker.GetId(), v.GetId())
		x.Equal(c.v_locker.GetName(), v.GetName())
		x.Equal(c.v_locker.GetNumber(), v.GetNumber())

		// Direct edges are loaded.
		x.True(v.HasOwner())

		w := v.GetOwner()
		x.Equal(c.v_member.GetId(), w.GetId())

		// Transient edges are not loaded.
		x.False(w.HasParent())
	}))
	t.Run("all fields with edge ID", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, c.v_locker.Pick().Select(library.LockerSelect_builder{
			All:   &v_t,
			Owner: library.MemberSelect_builder{}.Build(),
		}.Build()))
		x.NoError(err)
		x.Equal(c.v_locker.GetId(), v.GetId())
		x.Equal(c.v_locker.GetName(), v.GetName())
		x.Equal(c.v_locker.GetNumber(), v.GetNumber())

		// Direct edges are loaded.
		x.True(v.HasOwner())

		w := v.GetOwner()
		x.Equal(c.v_member.GetId(), w.GetId())

		// Transient edges are not loaded.
		x.False(w.HasParent())
	}))
	t.Run("all fields without edge IDs", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, c.v_locker.Pick().Select(library.LockerSelect_builder{
			All: &v_t,
		}.Build()))
		x.NoError(err)
		x.Equal(c.v_locker.GetId(), v.GetId())
		x.Equal(c.v_locker.GetName(), v.GetName())
		x.Equal(c.v_locker.GetNumber(), v.GetNumber())

		// Direct edges are not loaded.
		x.False(v.HasOwner())
	}))
	t.Run("specific fields", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, c.v_locker.Pick().Select(library.LockerSelect_builder{
			Name: &v_t,
		}.Build()))
		x.NoError(err)
		x.Equal(c.v_locker.GetName(), v.GetName())

		x.NotEqual(c.v_locker.GetOwner(), v.GetOwner())
		x.Empty(v.GetOwner())
		x.NotEqual(c.v_locker.GetNumber(), v.GetNumber())
		x.Empty(v.GetNumber())
	}))
	t.Run("specific field from edge", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, c.v_locker.Pick().Select(library.LockerSelect_builder{
			Owner: library.MemberSelect_builder{Name: &v_t}.Build(),
		}.Build()))
		x.NoError(err)
		x.True(v.HasOwner())

		x.NotEqual(c.v_locker.GetName(), v.GetName())
		x.Empty(v.GetName())
		x.NotEqual(c.v_locker.GetNumber(), v.GetNumber())
		x.Empty(v.GetNumber())

		w := v.GetOwner()
		x.Equal(c.v_member.GetId(), w.GetId())
		x.Equal(c.v_member.GetName(), w.GetName())

		x.NotEqual(c.v_member.GetLevel(), w.GetLevel())
		x.Zero(w.GetLevel())
		x.NotEqual(c.v_member.GetDateCreated(), w.GetDateCreated())
		x.Equal(timestamppb.New(time.Time{}), w.GetDateCreated())

		// Edges are not loaded.
		x.False(w.HasParent())
	}))
	t.Run("all fields from edge", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		v, err := c.Locker.Get(ctx, c.v_locker.Pick().Select(library.LockerSelect_builder{
			Owner: library.MemberSelect_builder{All: &v_t}.Build(),
		}.Build()))
		x.NoError(err)
		x.True(v.HasOwner())

		x.NotEqual(c.v_locker.GetName(), v.GetName())
		x.Empty(v.GetName())
		x.NotEqual(c.v_locker.GetNumber(), v.GetNumber())
		x.Empty(v.GetNumber())

		w := v.GetOwner()
		x.Equal(c.v_member.GetId(), w.GetId())
		x.Equal(c.v_member.GetName(), w.GetName())
		x.Equal(c.v_member.GetLabels(), w.GetLabels())
		x.Equal(c.v_member.GetProfile(), w.GetProfile())
		x.Equal(c.v_member.GetLevel(), w.GetLevel())
		x.Equal(c.v_member.GetDateCreated(), w.GetDateCreated())

		// Edges are not loaded.
		x.False(w.HasParent())
	}))
}
