// Code generated by "protoc-gen-orm-ent", DO NOT EDIT.

package schema

import (
	ent "entgo.io/ent"
	edge "entgo.io/ent/schema/edge"
	field "entgo.io/ent/schema/field"
	index "entgo.io/ent/schema/index"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	time "time"
)

type Member struct {
	ent.Schema
}

func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("name").
			Default(""),
		field.JSON("labels", map[string]string{}).
			Default(map[string]string{}),
		field.JSON("profile", &library.Profile{}),
		field.Int("level"),
		field.Time("date_created").
			Default(func() time.Time { return time.Now().UTC() }).
			Immutable(),
	}
}

func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("locker", Locker.Type).
			Unique(),
		edge.From("parent", Member.Type).Ref("children").
			Unique(),
		edge.To("children", Member.Type),
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("parent").
			Unique(),
	}
}
