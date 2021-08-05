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

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	var data interface{}
	return []ent.Field{
		field.String("tenant").NotEmpty().Comment("域标识,可自定义用于区分哪个平台使用"),
		field.String("name").NotEmpty().Comment("名称"),
		field.String("type").NotEmpty().Comment("节点类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等"),
		field.Int("parent_id").Optional().Comment("父节点ID"),
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

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Node.Type).
			From("parent").
			Unique().Field("parent_id"),
		edge.From("groups", Group.Type).
			Ref("nodes"),
	}
}

// Indexes of the Node.
func (Node) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant", "name", "type").Unique(),
	}
}
