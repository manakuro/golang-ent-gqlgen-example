package schema

import (
	"time"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent/dialect"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
		field.Int("age").Positive(),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("articles", Article.Type).
			StorageKey(
				edge.Column("user_id"),
				edge.Symbol("user_id"),
			),
	}
}
