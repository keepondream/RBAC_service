package schema

import (
	"net/http"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Route holds the schema definition for the Route entity.
type Route struct {
	ent.Schema
}

// Fields of the Route.
func (Route) Fields() []ent.Field {
	var data interface{}
	return []ent.Field{
		field.String("tenant").NotEmpty().Comment("域标识,可自定义用于区分哪个平台使用"),
		field.String("name").NotEmpty().Comment("路由名称"),
		field.String("uri").NotEmpty().Comment("拦截URL"),
		field.Enum("method").Values(
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		).Comment("URL请求方式"),
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

// Edges of the Route.
func (Route) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("permissions", Permission.Type).Ref("routes"),
	}
}

// Indexes of the Route.
func (Route) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant", "uri", "method").Unique(),
	}
}
