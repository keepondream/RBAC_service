// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/predicate"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/route"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where adds a new predicate for the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetTenant sets the "tenant" field.
func (pu *PermissionUpdate) SetTenant(s string) *PermissionUpdate {
	pu.mutation.SetTenant(s)
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetData sets the "data" field.
func (pu *PermissionUpdate) SetData(i *interface{}) *PermissionUpdate {
	pu.mutation.SetData(i)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PermissionUpdate) SetCreatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCreatedAt(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (pu *PermissionUpdate) AddRouteIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddRouteIDs(ids...)
	return pu
}

// AddRoutes adds the "routes" edges to the Route entity.
func (pu *PermissionUpdate) AddRoutes(r ...*Route) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRouteIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRoutes clears all "routes" edges to the Route entity.
func (pu *PermissionUpdate) ClearRoutes() *PermissionUpdate {
	pu.mutation.ClearRoutes()
	return pu
}

// RemoveRouteIDs removes the "routes" edge to Route entities by IDs.
func (pu *PermissionUpdate) RemoveRouteIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveRouteIDs(ids...)
	return pu
}

// RemoveRoutes removes "routes" edges to Route entities.
func (pu *PermissionUpdate) RemoveRoutes(r ...*Route) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRouteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.Tenant(); ok {
		if err := permission.TenantValidator(v); err != nil {
			return &ValidationError{Name: "tenant", err: fmt.Errorf("ent: validator failed for field \"tenant\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Tenant(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldTenant,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldName,
		})
	}
	if value, ok := pu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: permission.FieldData,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
	}
	if pu.mutation.RoutesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRoutesIDs(); len(nodes) > 0 && !pu.mutation.RoutesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RoutesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetTenant sets the "tenant" field.
func (puo *PermissionUpdateOne) SetTenant(s string) *PermissionUpdateOne {
	puo.mutation.SetTenant(s)
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetData sets the "data" field.
func (puo *PermissionUpdateOne) SetData(i *interface{}) *PermissionUpdateOne {
	puo.mutation.SetData(i)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PermissionUpdateOne) SetCreatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCreatedAt(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (puo *PermissionUpdateOne) AddRouteIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddRouteIDs(ids...)
	return puo
}

// AddRoutes adds the "routes" edges to the Route entity.
func (puo *PermissionUpdateOne) AddRoutes(r ...*Route) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRouteIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRoutes clears all "routes" edges to the Route entity.
func (puo *PermissionUpdateOne) ClearRoutes() *PermissionUpdateOne {
	puo.mutation.ClearRoutes()
	return puo
}

// RemoveRouteIDs removes the "routes" edge to Route entities by IDs.
func (puo *PermissionUpdateOne) RemoveRouteIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveRouteIDs(ids...)
	return puo
}

// RemoveRoutes removes "routes" edges to Route entities.
func (puo *PermissionUpdateOne) RemoveRoutes(r ...*Route) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRouteIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.Tenant(); ok {
		if err := permission.TenantValidator(v); err != nil {
			return &ValidationError{Name: "tenant", err: fmt.Errorf("ent: validator failed for field \"tenant\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Permission.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Tenant(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldTenant,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldName,
		})
	}
	if value, ok := puo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: permission.FieldData,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
	}
	if puo.mutation.RoutesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRoutesIDs(); len(nodes) > 0 && !puo.mutation.RoutesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RoutesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}