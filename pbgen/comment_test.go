package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestComment(t *testing.T) {
	t.Run("single line", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Comment{Value: "foo"}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`// foo`, v)
	})

	t.Run("multiple lines", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Comment{Value: "foo\nbar"}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`// foo
// bar`, v)
	})
}
