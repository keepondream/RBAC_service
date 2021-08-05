// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/route"
)

// Route is the model entity for the Route schema.
type Route struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Tenant holds the value of the "tenant" field.
	// 域标识,可自定义用于区分哪个平台使用
	Tenant string `json:"tenant,omitempty"`
	// Name holds the value of the "name" field.
	// 路由名称
	Name string `json:"name,omitempty"`
	// URI holds the value of the "uri" field.
	// 拦截URL
	URI string `json:"uri,omitempty"`
	// Method holds the value of the "method" field.
	// URL请求方式
	Method route.Method `json:"method,omitempty"`
	// Data holds the value of the "data" field.
	// 自定义json数据
	Data *interface{} `json:"data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RouteQuery when eager-loading is set.
	Edges RouteEdges `json:"edges"`
}

// RouteEdges holds the relations/edges for other nodes in the graph.
type RouteEdges struct {
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e RouteEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[0] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Route) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case route.FieldData:
			values[i] = new([]byte)
		case route.FieldID:
			values[i] = new(sql.NullInt64)
		case route.FieldTenant, route.FieldName, route.FieldURI, route.FieldMethod:
			values[i] = new(sql.NullString)
		case route.FieldCreatedAt, route.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Route", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Route fields.
func (r *Route) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case route.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case route.FieldTenant:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant", values[i])
			} else if value.Valid {
				r.Tenant = value.String
			}
		case route.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case route.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				r.URI = value.String
			}
		case route.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				r.Method = route.Method(value.String)
			}
		case route.FieldData:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		case route.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case route.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryPermissions queries the "permissions" edge of the Route entity.
func (r *Route) QueryPermissions() *PermissionQuery {
	return (&RouteClient{config: r.config}).QueryPermissions(r)
}

// Update returns a builder for updating this Route.
// Note that you need to call Route.Unwrap() before calling this method if this Route
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Route) Update() *RouteUpdateOne {
	return (&RouteClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Route entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Route) Unwrap() *Route {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Route is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Route) String() string {
	var builder strings.Builder
	builder.WriteString("Route(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", tenant=")
	builder.WriteString(r.Tenant)
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", uri=")
	builder.WriteString(r.URI)
	builder.WriteString(", method=")
	builder.WriteString(fmt.Sprintf("%v", r.Method))
	builder.WriteString(", data=")
	builder.WriteString(fmt.Sprintf("%v", r.Data))
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Routes is a parsable slice of Route.
type Routes []*Route

func (r Routes) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
