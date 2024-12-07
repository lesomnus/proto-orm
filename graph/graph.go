package graph

import (
	"errors"
	"fmt"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Graph map[protoreflect.FullName]*Entity

func NewGraph(fs []*protogen.File) (Graph, error) {
	g := Graph{}
	errs := []error{}
	for _, f := range fs {
		o_f := proto.GetExtension(f.Desc.Options(), orm.E_All).(*orm.FileOptions)
		if o_f == nil {
			o_f = &orm.FileOptions{}
		}

		is_enabled_all := false
		if o_f.Messages != nil {
			is_enabled_all = !o_f.Messages.Disabled
		}
		for _, m := range f.Messages {
			o_m := proto.GetExtension(m.Desc.Options(), orm.E_Message).(*orm.MessageOptions)
			if o_m == nil {
				o_m = &orm.MessageOptions{}
			}

			is_enabled := is_enabled_all
			if o_m.Enabled != nil {
				is_enabled = *o_m.Enabled
			}
			if !is_enabled {
				continue
			}

			e, err := NewEntity(f, m, o_m)
			if err != nil {
				err := fmt.Errorf("%s: %w", m.Desc.Name(), err)
				errs = append(errs, err)
				continue
			}

			g[e.FullName()] = e
		}
	}

	// Find source for messages of RPC.
	rpcs := map[protoreflect.FullName]*RpcMessage{}
	for _, e := range g {
		for _, r := range e.Rpcs {
			rpcs[r.Req.FullName] = r.Req
			rpcs[r.Res.FullName] = r.Res
		}
	}
	if len(rpcs) > 0 {
		for _, f := range fs {
			for _, s := range f.Services {
				for _, m := range s.Methods {
					if r, ok := rpcs[m.Input.Desc.FullName()]; ok {
						r.Source = m.Input
					}
					if r, ok := rpcs[m.Output.Desc.FullName()]; ok {
						r.Source = m.Output
					}
				}
			}
		}
	}

	// Make edges.
	for _, e := range g {
		for _, f := range e.Fields {
			if !f.IsEdge() {
				continue
			}

			o := proto.GetExtension(f.Desc.Options(), orm.E_Field).(*orm.FieldOptions)
			edge, err := g.newEdge(f, o)
			if err != nil {
				err := fmt.Errorf("%s: %w", f.Desc.FullName(), err)
				errs = append(errs, err)
				continue
			}

			f.Edge = edge
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return g, nil
}
