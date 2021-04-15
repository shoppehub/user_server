package schema

import (
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
		field.String("name"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		GmtMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
