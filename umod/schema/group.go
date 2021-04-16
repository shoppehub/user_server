package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Comment("用户的名称，比如真实姓名"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
