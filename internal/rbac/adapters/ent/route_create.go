// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/route"
)

// RouteCreate is the builder for creating a Route entity.
type RouteCreate struct {
	config
	mutation *RouteMutation
	hooks    []Hook
}

// SetTenant sets the "tenant" field.
func (rc *RouteCreate) SetTenant(s string) *RouteCreate {
	rc.mutation.SetTenant(s)
	return rc
}

// SetName sets the "name" field.
func (rc *RouteCreate) SetName(s string) *RouteCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetURI sets the "uri" field.
func (rc *RouteCreate) SetURI(s string) *RouteCreate {
	rc.mutation.SetURI(s)
	return rc
}

// SetMethod sets the "method" field.
func (rc *RouteCreate) SetMethod(r route.Method) *RouteCreate {
	rc.mutation.SetMethod(r)
	return rc
}

// SetData sets the "data" field.
func (rc *RouteCreate) SetData(i *interface{}) *RouteCreate {
	rc.mutation.SetData(i)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RouteCreate) SetCreatedAt(t time.Time) *RouteCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RouteCreate) SetNillableCreatedAt(t *time.Time) *RouteCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RouteCreate) SetUpdatedAt(t time.Time) *RouteCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RouteCreate) SetNillableUpdatedAt(t *time.Time) *RouteCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (rc *RouteCreate) AddPermissionIDs(ids ...int) *RouteCreate {
	rc.mutation.AddPermissionIDs(ids...)
	return rc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (rc *RouteCreate) AddPermissions(p ...*Permission) *RouteCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionIDs(ids...)
}

// Mutation returns the RouteMutation object of the builder.
func (rc *RouteCreate) Mutation() *RouteMutation {
	return rc.mutation
}

// Save creates the Route in the database.
func (rc *RouteCreate) Save(ctx context.Context) (*Route, error) {
	var (
		err  error
		node *Route
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RouteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RouteCreate) SaveX(ctx context.Context) *Route {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (rc *RouteCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := route.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := route.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RouteCreate) check() error {
	if _, ok := rc.mutation.Tenant(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New("ent: missing required field \"tenant\"")}
	}
	if v, ok := rc.mutation.Tenant(); ok {
		if err := route.TenantValidator(v); err != nil {
			return &ValidationError{Name: "tenant", err: fmt.Errorf("ent: validator failed for field \"tenant\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := route.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := rc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New("ent: missing required field \"uri\"")}
	}
	if v, ok := rc.mutation.URI(); ok {
		if err := route.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf("ent: validator failed for field \"uri\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New("ent: missing required field \"method\"")}
	}
	if v, ok := rc.mutation.Method(); ok {
		if err := route.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf("ent: validator failed for field \"method\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New("ent: missing required field \"data\"")}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (rc *RouteCreate) sqlSave(ctx context.Context) (*Route, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RouteCreate) createSpec() (*Route, *sqlgraph.CreateSpec) {
	var (
		_node = &Route{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: route.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: route.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Tenant(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: route.FieldTenant,
		})
		_node.Tenant = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: route.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.URI(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: route.FieldURI,
		})
		_node.URI = value
	}
	if value, ok := rc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: route.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := rc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: route.FieldData,
		})
		_node.Data = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: route.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: route.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := rc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.PermissionsTable,
			Columns: route.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RouteCreateBulk is the builder for creating many Route entities in bulk.
type RouteCreateBulk struct {
	config
	builders []*RouteCreate
}

// Save creates the Route entities in the database.
func (rcb *RouteCreateBulk) Save(ctx context.Context) ([]*Route, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Route, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RouteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RouteCreateBulk) SaveX(ctx context.Context) []*Route {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}