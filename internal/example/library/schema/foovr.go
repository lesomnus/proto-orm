// Code generated by "protoc-gen-orm-ent". DO NOT EDIT

package schema

import (
	ent "entgo.io/ent"
	field "entgo.io/ent/schema/field"
)

type FooVr struct {
	ent.Schema
}

func (FooVr) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable(),
		field.JSON("vr_double", []float64{}).
			Optional(),
		field.JSON("vr_int64", []int64{}).
			Optional(),
		field.JSON("vr_uint64", []uint64{}).
			Optional(),
		field.JSON("vr_bool", []bool{}).
			Optional(),
		field.JSON("vr_string", []string{}).
			Optional(),
		field.JSON("vr_bytes", [][]byte{}).
			Optional(),
	}
}
