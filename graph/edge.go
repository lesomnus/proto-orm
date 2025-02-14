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
	*FieldMeta
	// source *protogen.Field
	entity *Entity // Message that defines this edge.

	// Reverse is assigned in User-to-Pet edge.
	// Inverse is assigned in Pet-from-User edge.
	// -> Reverse always exists if Inverse is exists since "from" edge requires "to" edge.
	//
	//              +User[pet].Target
	//              |   +User[pet].Reverse -> User to Pet
	//              |   |
	// User.pet --> Pet.owner --> User.pet
	//                            |     |
	//                            |     +Pet[owner].Inverse. -> Pet from User
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
			FieldMeta: NewFieldMeta(f),
			entity:    e,

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
			err := fmt.Errorf("back-reference field must be an edge")
			errs = append(errs, err)
		} else if t.Target.FullName() != e.FullName() {
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

			t_actual := t.Type()
			t_expect := target.Key.Type()
			t_to_be := t_expect
			if t_actual != t_expect {
				err := fmt.Errorf("expected bound type to be %s but was %s", t_expect, t_actual)
				errs = append(errs, err)
			} else {
				t.typ = t_to_be
			}
		}
	}

	if opt := o.Ownership; opt != nil {
		n := opt.Name
		if n == "" {
			n = e.Name()
		}
		if _, ok := target.Fields[n]; ok {
			err := fmt.Errorf(`make inverse edge: name "%s" already exist`, n)
			errs = append(errs, err)
		} else {
			v.Reverse = &Edge{
				FieldMeta: &FieldMeta{
					kind:     protoreflect.MessageKind,
					fullName: target.FullName().Append(protoreflect.Name(opt.Name)),

					number: protowire.Number(opt.Number),
					goName: opt.Name,

					isList: opt.Shared,
					isMap:  false,
				},
				entity: target,

				Target:  e,
				Inverse: v,

				Nullable:  opt.Nullable,
				Immutable: opt.Immutable,
			}

			target.Fields[n] = v.Reverse
			target.Edges[n] = v.Reverse
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

func (e *Edge) ProtoType() string {
	return string(e.Target.Source.Desc.FullName())
}

func (e *Edge) ImportPath() string {
	return e.Target.Source.Desc.ParentFile().Path()
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
