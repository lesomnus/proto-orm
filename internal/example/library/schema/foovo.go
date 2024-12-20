// Code generated by "protoc-gen-orm-ent". DO NOT EDIT

package schema

import (
	ent "entgo.io/ent"
	field "entgo.io/ent/schema/field"
)

type FooVo struct {
	ent.Schema
}

func (FooVo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable(),
		field.Float("vo_double").
			Optional(),
		field.Int64("vo_int64").
			Optional(),
		field.Uint64("vo_uint64").
			Optional(),
		field.Bool("vo_bool").
			Optional(),
		field.String("vo_string").
			Optional(),
		field.Bytes("vo_bytes").
			Optional(),
	}
}