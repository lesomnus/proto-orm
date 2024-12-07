package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestRpc(t *testing.T) {
	t.Run("unary", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: "User"},
			Response: pbgen.RpcType{Type: "User"},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (User);`, v)
	})

	t.Run("stream request", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: "User", Stream: true},
			Response: pbgen.RpcType{Type: "User"},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (stream User) returns (User);`, v)
	})

	t.Run("stream response", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: "User"},
			Response: pbgen.RpcType{Type: "User", Stream: true},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (stream User);`, v)
	})

	t.Run("options", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: "User"},
			Response: pbgen.RpcType{Type: "User"},
			Options: []pbgen.Option{
				{
					Name: "foo.bar",
					Value: pbgen.Data{
						Fields: []pbgen.DataField{
							{
								Name:  "a",
								Value: pbgen.UnsafeLiteral{Value: "\"b\""},
							},
							{
								Name:  "c",
								Value: pbgen.UnsafeLiteral{Value: "\"d\""},
							},
						},
					},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (User) {
	option (foo.bar) = {
		a: "b"
		c: "d"
	};
}`, v)
	})

	t.Run("options with nested message", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: "User"},
			Response: pbgen.RpcType{Type: "User"},
			Options: []pbgen.Option{
				{
					Name: "foo.bar",
					Value: pbgen.Data{
						Fields: []pbgen.DataField{
							{
								Name:  "a",
								Value: pbgen.DataString{Value: "b"},
							},
							{
								Name:  "c",
								Value: pbgen.DataString{Value: "d"},
							},
							{
								Name: "nested",
								Value: pbgen.Data{
									Fields: []pbgen.DataField{
										{
											Name:  "nested_a",
											Value: pbgen.DataString{Value: "b"},
										},
										{
											Name:  "nested_c",
											Value: pbgen.DataString{Value: "d"},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (User) {
	option (foo.bar) = {
		a: "b"
		c: "d"
		nested: {
			nested_a: "b"
			nested_c: "d"
		}
	};
}`, v)
	})

	t.Run("options with list", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: "User"},
			Response: pbgen.RpcType{Type: "User"},
			Options: []pbgen.Option{
				{
					Name: "foo.bar",
					Value: pbgen.Data{
						Fields: []pbgen.DataField{
							{
								Name:  "a",
								Value: pbgen.DataString{Value: "b"},
							},
							{
								Name:  "c",
								Value: pbgen.DataString{Value: "d"},
							},
							{
								Name: "list",
								Value: pbgen.DataList{
									Values: []pbgen.Data{
										{
											Fields: []pbgen.DataField{
												{
													Name:  "a",
													Value: pbgen.DataString{Value: "b"},
												},
												{
													Name:  "c",
													Value: pbgen.DataString{Value: "d"},
												},
											},
										},
										{
											Fields: []pbgen.DataField{
												{
													Name:  "e",
													Value: pbgen.DataString{Value: "f"},
												},
												{
													Name:  "g",
													Value: pbgen.DataString{Value: "h"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (User) {
	option (foo.bar) = {
		a: "b"
		c: "d"
		list: [
			{
				a: "b"
				c: "d"
			},
			{
				e: "f"
				g: "h"
			}
		]
	};
}`, v)
	})
}
