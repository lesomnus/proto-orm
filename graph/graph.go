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

	ws := []func(){}
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

			e := newEntity(f, m)
			g[e.FullName()] = e

			// parse entities after declaration.
			ws = append(ws, func() {
				if err := e.prase(g, o_m); err != nil {
					err = fmt.Errorf("%s: %w", e.FullName(), err)
					errs = append(errs, err)
				}
			})
		}
	}

	// Prase entities.
	for _, w := range ws {
		w()
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return g, nil
}
