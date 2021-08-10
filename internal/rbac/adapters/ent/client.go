// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/migrate"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/group"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/node"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/route"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Node is the client for interacting with the Node builders.
	Node *NodeClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Route is the client for interacting with the Route builders.
	Route *RouteClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Group = NewGroupClient(c.config)
	c.Node = NewNodeClient(c.config)
	c.Permission = NewPermissionClient(c.config)
	c.Route = NewRouteClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Group:      NewGroupClient(cfg),
		Node:       NewNodeClient(cfg),
		Permission: NewPermissionClient(cfg),
		Route:      NewRouteClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Group:      NewGroupClient(cfg),
		Node:       NewNodeClient(cfg),
		Permission: NewPermissionClient(cfg),
		Route:      NewRouteClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Group.Use(hooks...)
	c.Node.Use(hooks...)
	c.Permission.Use(hooks...)
	c.Route.Use(hooks...)
	c.User.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Group.
func (c *GroupClient) QueryParent(gr *Group) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, group.ParentTable, group.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Group.
func (c *GroupClient) QueryChildren(gr *Group) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.ChildrenTable, group.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNodes queries the nodes edge of a Group.
func (c *GroupClient) QueryNodes(gr *Group) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.NodesTable, group.NodesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// NodeClient is a client for the Node schema.
type NodeClient struct {
	config
}

// NewNodeClient returns a client for the Node from the given config.
func NewNodeClient(c config) *NodeClient {
	return &NodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `node.Hooks(f(g(h())))`.
func (c *NodeClient) Use(hooks ...Hook) {
	c.hooks.Node = append(c.hooks.Node, hooks...)
}

// Create returns a create builder for Node.
func (c *NodeClient) Create() *NodeCreate {
	mutation := newNodeMutation(c.config, OpCreate)
	return &NodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Node entities.
func (c *NodeClient) CreateBulk(builders ...*NodeCreate) *NodeCreateBulk {
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Node.
func (c *NodeClient) Update() *NodeUpdate {
	mutation := newNodeMutation(c.config, OpUpdate)
	return &NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeClient) UpdateOne(n *Node) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNode(n))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeClient) UpdateOneID(id int) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNodeID(id))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Node.
func (c *NodeClient) Delete() *NodeDelete {
	mutation := newNodeMutation(c.config, OpDelete)
	return &NodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NodeClient) DeleteOne(n *Node) *NodeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NodeClient) DeleteOneID(id int) *NodeDeleteOne {
	builder := c.Delete().Where(node.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeDeleteOne{builder}
}

// Query returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{
		config: c.config,
	}
}

// Get returns a Node entity by its id.
func (c *NodeClient) Get(ctx context.Context, id int) (*Node, error) {
	return c.Query().Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeClient) GetX(ctx context.Context, id int) *Node {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Node.
func (c *NodeClient) QueryParent(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, node.ParentTable, node.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Node.
func (c *NodeClient) QueryChildren(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, node.ChildrenTable, node.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a Node.
func (c *NodeClient) QueryGroups(n *Node) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, node.GroupsTable, node.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermissions queries the permissions edge of a Node.
func (c *NodeClient) QueryPermissions(n *Node) *PermissionQuery {
	query := &PermissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, node.PermissionsTable, node.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Node.
func (c *NodeClient) QueryUsers(n *Node) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, node.UsersTable, node.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeClient) Hooks() []Hook {
	return c.hooks.Node
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Create returns a create builder for Permission.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id int) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PermissionClient) DeleteOneID(id int) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{
		config: c.config,
	}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id int) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id int) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoutes queries the routes edge of a Permission.
func (c *PermissionClient) QueryRoutes(pe *Permission) *RouteQuery {
	query := &RouteQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(route.Table, route.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, permission.RoutesTable, permission.RoutesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNodes queries the nodes edge of a Permission.
func (c *PermissionClient) QueryNodes(pe *Permission) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.NodesTable, permission.NodesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Permission.
func (c *PermissionClient) QueryUsers(pe *Permission) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, permission.UsersTable, permission.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// RouteClient is a client for the Route schema.
type RouteClient struct {
	config
}

// NewRouteClient returns a client for the Route from the given config.
func NewRouteClient(c config) *RouteClient {
	return &RouteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `route.Hooks(f(g(h())))`.
func (c *RouteClient) Use(hooks ...Hook) {
	c.hooks.Route = append(c.hooks.Route, hooks...)
}

// Create returns a create builder for Route.
func (c *RouteClient) Create() *RouteCreate {
	mutation := newRouteMutation(c.config, OpCreate)
	return &RouteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Route entities.
func (c *RouteClient) CreateBulk(builders ...*RouteCreate) *RouteCreateBulk {
	return &RouteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Route.
func (c *RouteClient) Update() *RouteUpdate {
	mutation := newRouteMutation(c.config, OpUpdate)
	return &RouteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RouteClient) UpdateOne(r *Route) *RouteUpdateOne {
	mutation := newRouteMutation(c.config, OpUpdateOne, withRoute(r))
	return &RouteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RouteClient) UpdateOneID(id int) *RouteUpdateOne {
	mutation := newRouteMutation(c.config, OpUpdateOne, withRouteID(id))
	return &RouteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Route.
func (c *RouteClient) Delete() *RouteDelete {
	mutation := newRouteMutation(c.config, OpDelete)
	return &RouteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RouteClient) DeleteOne(r *Route) *RouteDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RouteClient) DeleteOneID(id int) *RouteDeleteOne {
	builder := c.Delete().Where(route.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RouteDeleteOne{builder}
}

// Query returns a query builder for Route.
func (c *RouteClient) Query() *RouteQuery {
	return &RouteQuery{
		config: c.config,
	}
}

// Get returns a Route entity by its id.
func (c *RouteClient) Get(ctx context.Context, id int) (*Route, error) {
	return c.Query().Where(route.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RouteClient) GetX(ctx context.Context, id int) *Route {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPermissions queries the permissions edge of a Route.
func (c *RouteClient) QueryPermissions(r *Route) *PermissionQuery {
	query := &PermissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(route.Table, route.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, route.PermissionsTable, route.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RouteClient) Hooks() []Hook {
	return c.hooks.Route
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a User.
func (c *UserClient) QueryParent(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.ParentTable, user.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a User.
func (c *UserClient) QueryChildren(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ChildrenTable, user.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.GroupsTable, user.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNodes queries the nodes edge of a User.
func (c *UserClient) QueryNodes(u *User) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.NodesTable, user.NodesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermissions queries the permissions edge of a User.
func (c *UserClient) QueryPermissions(u *User) *PermissionQuery {
	query := &PermissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.PermissionsTable, user.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
