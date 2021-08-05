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

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	var data interface{}
	return []ent.Field{
		field.String("tenant").NotEmpty().Comment("域标识,可自定义用于区分哪个平台使用"),
		field.String("name").NotEmpty().Comment("权限名称"),
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

// Indexes of the Fund.
func (Permission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant", "name").Unique(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("routes", Route.Type),
	}
}
