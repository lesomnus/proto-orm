// Code generated by "protoc-gen-orm-ent", DO NOT EDIT.

package schema

import (
	ent "entgo.io/ent"
	edge "entgo.io/ent/schema/edge"
	field "entgo.io/ent/schema/field"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	time "time"
)

type Author struct {
	ent.Schema
}

func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("alias").
			Unique(),
		field.String("name"),
		field.JSON("labels", map[string]string{}).
			Default(map[string]string{}),
		field.JSON("profile", &library.Profile{}).
			Default(&library.Profile{}),
		field.Time("date_created").
			Default(time.Now).
			Immutable(),
	}
}

func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type),
	}
}
