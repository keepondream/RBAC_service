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

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetTenant sets the "tenant" field.
func (pc *PermissionCreate) SetTenant(s string) *PermissionCreate {
	pc.mutation.SetTenant(s)
	return pc
}

// SetName sets the "name" field.
func (pc *PermissionCreate) SetName(s string) *PermissionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetData sets the "data" field.
func (pc *PermissionCreate) SetData(i *interface{}) *PermissionCreate {
	pc.mutation.SetData(i)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PermissionCreate) SetCreatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PermissionCreate) SetUpdatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (pc *PermissionCreate) AddRouteIDs(ids ...int) *PermissionCreate {
	pc.mutation.AddRouteIDs(ids...)
	return pc
}

// AddRoutes adds the "routes" edges to the Route entity.
func (pc *PermissionCreate) AddRoutes(r ...*Route) *PermissionCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRouteIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := permission.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := permission.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.Tenant(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New("ent: missing required field \"tenant\"")}
	}
	if v, ok := pc.mutation.Tenant(); ok {
		if err := permission.TenantValidator(v); err != nil {
			return &ValidationError{Name: "tenant", err: fmt.Errorf("ent: validator failed for field \"tenant\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New("ent: missing required field \"data\"")}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: permission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Tenant(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldTenant,
		})
		_node.Tenant = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: permission.FieldData,
		})
		_node.Data = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RoutesTable,
			Columns: permission.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: route.FieldID,
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

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
