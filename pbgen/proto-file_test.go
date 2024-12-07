package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestProtoFile(t *testing.T) {
	t.Run("extension only", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`edition = "2023";
`, v)
	})

	t.Run("package", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			Package: "foo.bar.baz",
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`edition = "2023";

package foo.bar.baz;
`, v)
	})

	t.Run("import", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			Imports: []pbgen.Import{
				{Name: "/foo/a.proto"},
				{Name: "/foo/b.proto", Visibility: pbgen.VisibilityWeak},
				{Name: "/foo/c.proto", Visibility: pbgen.VisibilityPublic},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`edition = "2023";

import "/foo/a.proto";
import weak "/foo/b.proto";
import public "/foo/c.proto";
`, v)
	})

	t.Run("option", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			Options: []pbgen.Option{pbgen.OptionGoPackage("foo.bar.baz")},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`edition = "2023";

option go_package = "foo.bar.baz";
`, v)
	})

	t.Run("messages", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			TopLevelDefinitions: []pbgen.TopLevelDef{
				pbgen.Message{
					Name: "User",
					Body: []pbgen.MessageBody{
						pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
					},
				},
				pbgen.Message{
					Name: "Account",
					Body: []pbgen.MessageBody{
						pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
					},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`edition = "2023";

message User {
	bytes id = 1;
}

message Account {
	bytes id = 1;
}
`, v)
	})

	t.Run("messages with comment", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			TopLevelDefinitions: []pbgen.TopLevelDef{
				pbgen.Comment{Value: "foo"},
				pbgen.Message{
					Name: "User",
					Body: []pbgen.MessageBody{
						pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
					},
				},
				pbgen.Comment{Value: "bar"},
				pbgen.Message{
					Name: "Account",
					Body: []pbgen.MessageBody{
						pbgen.MessageField{Type: pbgen.TypeBytes, Name: "id", Number: 1},
					},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`edition = "2023";

// foo
message User {
	bytes id = 1;
}

// bar
message Account {
	bytes id = 1;
}
`, v)
	})
}
