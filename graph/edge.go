package graph

import (
	"errors"
	"fmt"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Edge struct {
	source *protogen.Field
	entity *Entity // Message that defines this edge.

	//
	//              +User[pet].Target
	//              |   +User[pet].Reverse
	//              |   |
	// User.pet --> Pet.owner --> User.pet
	//                            |     |
	//                            |     +Pet[owner].Inverse.
	//                            +Pet[owner].Target
	//
	Target  *Entity
	Reverse *Edge // Back-reference field for this edge.
	Inverse *Edge // Edge that references this edge.
	Bind    *Attr

	Ignored   bool
	Unique    bool
	Nullable  bool
	Immutable bool
}

func (e *Entity) parseEdge(f *protogen.Field, o *orm.EdgeOptions, target *Entity) (*Edge, error) {
	if o == nil {
		o = &orm.EdgeOptions{}
	}

	n := string(f.Desc.Name())
	v, ok := e.Edges[n]
	if !ok {
		v = &Edge{
			source: f,
			entity: e,

			Target: target,

			Unique:    !f.Desc.IsList(),
			Nullable:  o.Nullable,
			Immutable: o.Immutable,
		}
	}

	e.Fields[n] = v
	e.Edges[n] = v

	errs := []error{}
	if o.From > 0 {
		f := target.FieldByNumber(protowire.Number(o.From))
		if f == nil {
			err := fmt.Errorf("back-reference field not found: %d", o.From)
			errs = append(errs, err)
		} else if t, ok := f.(*Edge); !ok {
			err := fmt.Errorf("back-reference field must be edge")
			errs = append(errs, err)
		} else if t.source.Message.Desc.FullName() != e.FullName() {
			err := fmt.Errorf("back-reference edge does not reference this entity")
			errs = append(errs, err)
		} else {
			v.Inverse = t
			t.Reverse = v
		}
	}
	if o.Bind > 0 {
		f := e.FieldByNumber(protowire.Number(o.Bind))
		if f == nil {
			err := fmt.Errorf("bound field not found")
			errs = append(errs, err)
		} else if t, ok := f.(*Attr); !ok {
			err := fmt.Errorf("bound field must be a scalar field")
			errs = append(errs, err)
		} else {
			v.Bind = t
			t.Bound = v

			t_actual := t.typ
			t_expect := target.Key.Type()
			t_to_be := t_expect
			if t_expect == orm.Type_TYPE_UUID {
				t_expect = orm.Type_TYPE_BYTES
			}
			if t_actual != t_expect {
				err := fmt.Errorf("expected bound type to be %s but %s", t_expect, t_actual)
				errs = append(errs, err)
			} else {
				t.typ = t_to_be
			}
		}
	}
	if o.Immutable {
		v.setImmutable()
	}
	if o.Nullable {
		v.setNullable()
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return v, nil
}

func (e *Edge) setImmutable() {
	e.Immutable = true
	if e.Bind != nil {
		e.Bind.Immutable = true
	}
}

func (e *Edge) setNullable() {
	e.Nullable = true
	if e.Bind != nil {
		e.Bind.Nullable = true
	}
}

func (e *Edge) FullName() protoreflect.FullName {
	return e.source.Desc.FullName()
}

func (e *Edge) Name() string {
	return string(e.source.Desc.Name())
}

func (e *Edge) Number() protowire.Number {
	return e.source.Desc.Number()
}

func (e *Edge) ProtoType() string {
	return string(e.source.Message.Desc.FullName())
}

func (a *Edge) ImportPath() string {
	return a.source.Message.Desc.ParentFile().Path()
}

func (e *Edge) GoName() string {
	return e.source.GoName
}

func (e *Edge) Source() *protogen.Field {
	return e.source
}

func (e *Edge) Entity() *Entity {
	return e.entity
}

func (e *Edge) HasBind() bool {
	return e.Bind != nil
}

func (e *Edge) IsIgnored() bool {
	return e.Ignored
}

func (e *Edge) IsRequired() bool {
	return !e.Nullable && !e.IsList()
}

func (e *Edge) IsOptional() bool {
	return !e.IsRequired()
}

func (e *Edge) IsList() bool {
	return e.source.Desc.IsList()
}

func (e *Edge) IsUnique() bool {
	// TODO: I think edge is unique if O2O.
	// In graph, `Edge` is directional so current Unique actually means a singular.
	return e.Unique
}

func (e *Edge) IsNullable() bool {
	return e.Nullable
}

func (e *Edge) IsImmutable() bool {
	return e.Immutable
}

func (e *Edge) HasInverse() bool {
	return e.Inverse != nil
}

func (e *Edge) IsSelfLoop() bool {
	if e.Inverse == nil {
		return false
	}
	return e.Inverse.entity == e.entity
}

func (e *Edge) IsUnidirectional() bool {
	return e.Reverse == nil && e.Inverse == nil && !e.IsList()
}

func (e *Edge) IsExclusive() bool {
	if e.IsUnidirectional() {
		return false
	}

	return !e.IsList()
}

func (e *Edge) Opposite() *Edge {
	if e.Reverse != nil {
		return e.Reverse
	} else {
		return e.Inverse
	}
}
