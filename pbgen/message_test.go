package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	p := pbgen.NewPrinter("")

	t.Run("single field", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Message{
			FullName: "User",
			Body: []pbgen.MessageBody{
				pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
			},
		}
		o := bytes.Buffer{}
		err := p.Print(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`message User {
	bytes id = 1;
}`, v)
	})

	t.Run("multiple fields", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Message{
			FullName: "User",
			Body: []pbgen.MessageBody{
				pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
				pbgen.MessageField{Type: pbgen.TypeString, Name: "name", Number: 2},
			},
		}
		o := bytes.Buffer{}
		err := p.Print(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`message User {
	bytes id = 1;
	string name = 2;
}`, v)
	})

	t.Run("multiple fields with comment", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Message{
			FullName: "User",
			Body: []pbgen.MessageBody{
				pbgen.Comment{Value: "foo\nbar\nbaz"},
				pbgen.Comment{Value: "Royale"},
				pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
				pbgen.Comment{Value: "with"},
				pbgen.MessageField{Type: pbgen.TypeString, Name: "name", Number: 2},
				pbgen.Comment{Value: "cheese"},
			},
		}
		o := bytes.Buffer{}
		err := p.Print(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`message User {
	// foo
	// bar
	// baz
	// Royale
	bytes id = 1;
	// with
	string name = 2;
	// cheese
}`, v)
	})

	t.Run("oneof", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Message{
			FullName: "User",
			Body: []pbgen.MessageBody{
				pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
				pbgen.MessageField{Type: pbgen.TypeString, Name: "name", Number: 2},
				pbgen.MessageOneof{
					Name: "auth",
					Body: []pbgen.MessageOneofBody{
						pbgen.MessageOneofField{Type: pbgen.TypeString, Name: "password", Number: 3},
						pbgen.MessageOneofField{Type: pbgen.TypeString, Name: "email", Number: 4},
					},
				},
			},
		}
		o := bytes.Buffer{}
		err := p.Print(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`message User {
	bytes id = 1;
	string name = 2;
	oneof auth {
		string password = 3;
		string email = 4;
	}
}`, v)
	})
}
