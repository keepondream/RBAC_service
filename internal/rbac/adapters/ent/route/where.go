// Code generated by entc, DO NOT EDIT.

package route

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Tenant applies equality check predicate on the "tenant" field. It's identical to TenantEQ.
func Tenant(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenant), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURI), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TenantEQ applies the EQ predicate on the "tenant" field.
func TenantEQ(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenant), v))
	})
}

// TenantNEQ applies the NEQ predicate on the "tenant" field.
func TenantNEQ(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenant), v))
	})
}

// TenantIn applies the In predicate on the "tenant" field.
func TenantIn(vs ...string) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTenant), v...))
	})
}

// TenantNotIn applies the NotIn predicate on the "tenant" field.
func TenantNotIn(vs ...string) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTenant), v...))
	})
}

// TenantGT applies the GT predicate on the "tenant" field.
func TenantGT(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenant), v))
	})
}

// TenantGTE applies the GTE predicate on the "tenant" field.
func TenantGTE(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenant), v))
	})
}

// TenantLT applies the LT predicate on the "tenant" field.
func TenantLT(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenant), v))
	})
}

// TenantLTE applies the LTE predicate on the "tenant" field.
func TenantLTE(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenant), v))
	})
}

// TenantContains applies the Contains predicate on the "tenant" field.
func TenantContains(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTenant), v))
	})
}

// TenantHasPrefix applies the HasPrefix predicate on the "tenant" field.
func TenantHasPrefix(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTenant), v))
	})
}

// TenantHasSuffix applies the HasSuffix predicate on the "tenant" field.
func TenantHasSuffix(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTenant), v))
	})
}

// TenantEqualFold applies the EqualFold predicate on the "tenant" field.
func TenantEqualFold(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTenant), v))
	})
}

// TenantContainsFold applies the ContainsFold predicate on the "tenant" field.
func TenantContainsFold(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTenant), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURI), v))
	})
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURI), v))
	})
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURI), v...))
	})
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURI), v...))
	})
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURI), v))
	})
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURI), v))
	})
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURI), v))
	})
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURI), v))
	})
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURI), v))
	})
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURI), v))
	})
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURI), v))
	})
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURI), v))
	})
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURI), v))
	})
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v Method) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v Method) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMethod), v))
	})
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...Method) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMethod), v...))
	})
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...Method) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMethod), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Route {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Route(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PermissionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PermissionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		p(s.Not())
	})
}
