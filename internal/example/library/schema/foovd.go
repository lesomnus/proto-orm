// Code generated by "protoc-gen-orm-ent". DO NOT EDIT

package schema

import (
	ent "entgo.io/ent"
	field "entgo.io/ent/schema/field"
)

type FooVd struct {
	ent.Schema
}

func (FooVd) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable(),
		field.Float("vd_double"),
		field.Int64("vd_int64"),
		field.Uint64("vd_uint64"),
		field.Bool("vd_bool"),
		field.String("vd_string"),
		field.Bytes("vd_bytes"),
	}
}
