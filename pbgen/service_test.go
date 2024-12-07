package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("single rpc", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Service{
			Name: "UserService",
			Body: []pbgen.ServiceBody{
				pbgen.Rpc{
					Name:     "Create",
					Request:  pbgen.RpcType{Type: "User"},
					Response: pbgen.RpcType{Type: "User"},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`service UserService {
	rpc Create (User) returns (User);
}`, v)
	})

	t.Run("multiple rpcs", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Service{
			Name: "UserService",
			Body: []pbgen.ServiceBody{
				pbgen.Rpc{
					Name:     "Create",
					Request:  pbgen.RpcType{Type: "User"},
					Response: pbgen.RpcType{Type: "User"},
				},
				pbgen.Rpc{
					Name:     "Update",
					Request:  pbgen.RpcType{Type: "User"},
					Response: pbgen.RpcType{Type: "User"},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`service UserService {
	rpc Create (User) returns (User);
	rpc Update (User) returns (User);
}`, v)
	})

	t.Run("multiple rpcs with comment", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Service{
			Name: "UserService",
			Body: []pbgen.ServiceBody{
				pbgen.Comment{Value: "foo\nbar\nbaz"},
				pbgen.Comment{Value: "Royale"},
				pbgen.Rpc{
					Name:     "Create",
					Request:  pbgen.RpcType{Type: "User"},
					Response: pbgen.RpcType{Type: "User"},
				},
				pbgen.Comment{Value: "with"},
				pbgen.Rpc{
					Name:     "Update",
					Request:  pbgen.RpcType{Type: "User"},
					Response: pbgen.RpcType{Type: "User"},
				},
				pbgen.Comment{Value: "cheese"},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`service UserService {
	// foo
	// bar
	// baz
	// Royale
	rpc Create (User) returns (User);
	// with
	rpc Update (User) returns (User);
	// cheese
}`, v)
	})
}
