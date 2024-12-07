package graph

import (
	"errors"
	"fmt"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/encoding/protowire"
)

type Edge struct {
	Source *Field
	Target *Entity

	From *Field
	Bind *Field
}

func (g Graph) newEdge(f *Field, o *orm.FieldOptions) (*Edge, error) {
	if o == nil {
		o = &orm.FieldOptions{}
	}
	e := &Edge{Source: f}

	target, ok := g[f.Message.Desc.FullName()]
	if !ok {
		return nil, fmt.Errorf("target message not found")
	}
	e.Target = target

	errs := []error{}
	if o.From > 0 {
		if f_from := target.FieldByNumber(protowire.Number(o.From)); f_from == nil {
			err := fmt.Errorf("referencing field not found")
			errs = append(errs, err)
		} else {
			e.From = f_from
		}
	}
	if o.Bind > 0 {
		if f_bound := f.Entity.FieldByNumber(protowire.Number(o.Bind)); f_bound == nil {
			err := fmt.Errorf("bound field not found")
			errs = append(errs, err)
		} else {
			e.Bind = f_bound
			f_bound.Bound = f

			// Note that uniqueness of the index does not make its members unique.
			// Assume unique index of columns X and Y, there may be (a, b) and (a, c), so X is not unique.
			if o.Immutable {
				f_bound.Immutable = true
			}
			if f.Entity.Key.Type() == orm.Type_TYPE_UUID {
				f_bound.GivenType = orm.Type_TYPE_UUID
			}
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return e, nil
}

func (e *Edge) HasBind() bool {
	return e.Bind != nil
}

func (e *Edge) HasMany() bool {
	return e.Source.IsList()
}

func (e *Edge) IsSelfLoop() bool {
	if e.From == nil {
		return false
	}
	return e.From.Entity == e.Source.Entity
}
