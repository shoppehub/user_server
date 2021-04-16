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
		field.String("nickName"),
		field.String("username"),
		field.String("password"),
		field.String("email").Comment("用户的邮箱，支持邮箱登录"),
		field.String("mobile").Comment("手机号码，支持手机号码登录"),
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
