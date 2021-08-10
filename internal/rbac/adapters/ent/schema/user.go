package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	var data interface{}
	return []ent.Field{
		field.String("tenant").NotEmpty().Comment("域标识,可自定义用于区分哪个平台使用"),
		field.String("uuid").NotEmpty().Comment("用户唯一标识,由认证服务或者第三方服务提供的唯一标识"),
		field.Int("parent_id").Optional().Comment("父节点ID"),
		field.Bool("is_super").Default(false).Comment("是否为超级管理员,该标识意味着当前用户不需要鉴权,有系统最大权限"),
		field.JSON("data", &data).Comment("自定义json数据"),
		field.Time("created_at").SchemaType(map[string]string{
			dialect.Postgres: "timestamptz(0)",
		}).Default(time.Now().UTC).Comment("创建时间").Annotations(entsql.Annotation{
			Default: "(now() at time zone 'utc')",
		}),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.Postgres: "timestamptz(0)",
		}).Default(time.Now().UTC).Comment("更新时间").Annotations(entsql.Annotation{
			Default: "(now() at time zone 'utc')",
		}).UpdateDefault(time.Now().UTC),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant", "uuid").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", User.Type).
			From("parent").
			Unique().Field("parent_id"),
		edge.From("groups", Group.Type).
			Ref("users"),
		edge.From("nodes", Node.Type).
			Ref("users"),
		edge.From("permissions", Permission.Type).
			Ref("users"),
	}
}
