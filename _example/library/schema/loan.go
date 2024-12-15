// Code generated by "protoc-gen-orm-ent". DO NOT EDIT

package schema

import (
	ent "entgo.io/ent"
	edge "entgo.io/ent/schema/edge"
	field "entgo.io/ent/schema/field"
	uuid "github.com/google/uuid"
	time "time"
)

type Loan struct {
	ent.Schema
}

func (Loan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.UUID("subject_id", uuid.UUID{}).
			Immutable(),
		field.UUID("borrower_id", uuid.UUID{}).
			Immutable(),
		field.Time("date_due"),
		field.Time("date_return").
			Optional().
			Nillable(),
		field.Time("date_created").
			Default(time.Now).
			Immutable(),
	}
}

func (Loan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subject", Book.Type).
			Field("subject_id").
			Unique().
			Required().
			Immutable(),
		edge.To("borrower", Member.Type).
			Field("borrower_id").
			Unique().
			Required().
			Immutable(),
	}
}
