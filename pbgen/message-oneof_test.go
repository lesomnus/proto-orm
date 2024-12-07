package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestMessageOneof(t *testing.T) {
	t.Run("single field", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageOneof{
			Name: "key",
			Body: []pbgen.MessageOneofBody{
				pbgen.MessageOneofField{
					Type:   pbgen.TypeBytes,
					Name:   "id",
					Number: 1,
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`oneof key {
	bytes id = 1;
}`, v)
	})

	t.Run("multiple fields", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageOneof{
			Name: "key",
			Body: []pbgen.MessageOneofBody{
				pbgen.MessageOneofField{
					Type:   pbgen.TypeBytes,
					Name:   "id",
					Number: 1,
				},
				pbgen.Comment{Value: "Baz"},
				pbgen.MessageOneofField{
					Type:   "foo.bar.baz",
					Name:   "alias",
					Number: 2,
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`oneof key {
	bytes id = 1;
	// Baz
	foo.bar.baz alias = 2;
}`, v)
	})
}
