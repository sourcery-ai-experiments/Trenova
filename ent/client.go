// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/emoss08/trenova/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// BusinessUnit is the client for interacting with the BusinessUnit builders.
	BusinessUnit *BusinessUnitClient
	// Organization is the client for interacting with the Organization builders.
	Organization *OrganizationClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.BusinessUnit = NewBusinessUnitClient(c.config)
	c.Organization = NewOrganizationClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		BusinessUnit: NewBusinessUnitClient(cfg),
		Organization: NewOrganizationClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:          ctx,
		config:       cfg,
		BusinessUnit: NewBusinessUnitClient(cfg),
		Organization: NewOrganizationClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		BusinessUnit.
//		Query().
//		Count(ctx)
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
	c.BusinessUnit.Use(hooks...)
	c.Organization.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.BusinessUnit.Intercept(interceptors...)
	c.Organization.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BusinessUnitMutation:
		return c.BusinessUnit.mutate(ctx, m)
	case *OrganizationMutation:
		return c.Organization.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BusinessUnitClient is a client for the BusinessUnit schema.
type BusinessUnitClient struct {
	config
}

// NewBusinessUnitClient returns a client for the BusinessUnit from the given config.
func NewBusinessUnitClient(c config) *BusinessUnitClient {
	return &BusinessUnitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `businessunit.Hooks(f(g(h())))`.
func (c *BusinessUnitClient) Use(hooks ...Hook) {
	c.hooks.BusinessUnit = append(c.hooks.BusinessUnit, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `businessunit.Intercept(f(g(h())))`.
func (c *BusinessUnitClient) Intercept(interceptors ...Interceptor) {
	c.inters.BusinessUnit = append(c.inters.BusinessUnit, interceptors...)
}

// Create returns a builder for creating a BusinessUnit entity.
func (c *BusinessUnitClient) Create() *BusinessUnitCreate {
	mutation := newBusinessUnitMutation(c.config, OpCreate)
	return &BusinessUnitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BusinessUnit entities.
func (c *BusinessUnitClient) CreateBulk(builders ...*BusinessUnitCreate) *BusinessUnitCreateBulk {
	return &BusinessUnitCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BusinessUnitClient) MapCreateBulk(slice any, setFunc func(*BusinessUnitCreate, int)) *BusinessUnitCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BusinessUnitCreateBulk{err: fmt.Errorf("calling to BusinessUnitClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BusinessUnitCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BusinessUnitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BusinessUnit.
func (c *BusinessUnitClient) Update() *BusinessUnitUpdate {
	mutation := newBusinessUnitMutation(c.config, OpUpdate)
	return &BusinessUnitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BusinessUnitClient) UpdateOne(bu *BusinessUnit) *BusinessUnitUpdateOne {
	mutation := newBusinessUnitMutation(c.config, OpUpdateOne, withBusinessUnit(bu))
	return &BusinessUnitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BusinessUnitClient) UpdateOneID(id uuid.UUID) *BusinessUnitUpdateOne {
	mutation := newBusinessUnitMutation(c.config, OpUpdateOne, withBusinessUnitID(id))
	return &BusinessUnitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BusinessUnit.
func (c *BusinessUnitClient) Delete() *BusinessUnitDelete {
	mutation := newBusinessUnitMutation(c.config, OpDelete)
	return &BusinessUnitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BusinessUnitClient) DeleteOne(bu *BusinessUnit) *BusinessUnitDeleteOne {
	return c.DeleteOneID(bu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BusinessUnitClient) DeleteOneID(id uuid.UUID) *BusinessUnitDeleteOne {
	builder := c.Delete().Where(businessunit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BusinessUnitDeleteOne{builder}
}

// Query returns a query builder for BusinessUnit.
func (c *BusinessUnitClient) Query() *BusinessUnitQuery {
	return &BusinessUnitQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBusinessUnit},
		inters: c.Interceptors(),
	}
}

// Get returns a BusinessUnit entity by its id.
func (c *BusinessUnitClient) Get(ctx context.Context, id uuid.UUID) (*BusinessUnit, error) {
	return c.Query().Where(businessunit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BusinessUnitClient) GetX(ctx context.Context, id uuid.UUID) *BusinessUnit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BusinessUnit.
func (c *BusinessUnitClient) QueryParent(bu *BusinessUnit) *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := bu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(businessunit.Table, businessunit.FieldID, id),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, businessunit.ParentTable, businessunit.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BusinessUnitClient) Hooks() []Hook {
	return c.hooks.BusinessUnit
}

// Interceptors returns the client interceptors.
func (c *BusinessUnitClient) Interceptors() []Interceptor {
	return c.inters.BusinessUnit
}

func (c *BusinessUnitClient) mutate(ctx context.Context, m *BusinessUnitMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BusinessUnitCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BusinessUnitUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BusinessUnitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BusinessUnitDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown BusinessUnit mutation op: %q", m.Op())
	}
}

// OrganizationClient is a client for the Organization schema.
type OrganizationClient struct {
	config
}

// NewOrganizationClient returns a client for the Organization from the given config.
func NewOrganizationClient(c config) *OrganizationClient {
	return &OrganizationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `organization.Hooks(f(g(h())))`.
func (c *OrganizationClient) Use(hooks ...Hook) {
	c.hooks.Organization = append(c.hooks.Organization, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `organization.Intercept(f(g(h())))`.
func (c *OrganizationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Organization = append(c.inters.Organization, interceptors...)
}

// Create returns a builder for creating a Organization entity.
func (c *OrganizationClient) Create() *OrganizationCreate {
	mutation := newOrganizationMutation(c.config, OpCreate)
	return &OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Organization entities.
func (c *OrganizationClient) CreateBulk(builders ...*OrganizationCreate) *OrganizationCreateBulk {
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrganizationClient) MapCreateBulk(slice any, setFunc func(*OrganizationCreate, int)) *OrganizationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrganizationCreateBulk{err: fmt.Errorf("calling to OrganizationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrganizationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Organization.
func (c *OrganizationClient) Update() *OrganizationUpdate {
	mutation := newOrganizationMutation(c.config, OpUpdate)
	return &OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrganizationClient) UpdateOne(o *Organization) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganization(o))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrganizationClient) UpdateOneID(id uuid.UUID) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganizationID(id))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Organization.
func (c *OrganizationClient) Delete() *OrganizationDelete {
	mutation := newOrganizationMutation(c.config, OpDelete)
	return &OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrganizationClient) DeleteOne(o *Organization) *OrganizationDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrganizationClient) DeleteOneID(id uuid.UUID) *OrganizationDeleteOne {
	builder := c.Delete().Where(organization.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrganizationDeleteOne{builder}
}

// Query returns a query builder for Organization.
func (c *OrganizationClient) Query() *OrganizationQuery {
	return &OrganizationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrganization},
		inters: c.Interceptors(),
	}
}

// Get returns a Organization entity by its id.
func (c *OrganizationClient) Get(ctx context.Context, id uuid.UUID) (*Organization, error) {
	return c.Query().Where(organization.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrganizationClient) GetX(ctx context.Context, id uuid.UUID) *Organization {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OrganizationClient) Hooks() []Hook {
	return c.hooks.Organization
}

// Interceptors returns the client interceptors.
func (c *OrganizationClient) Interceptors() []Interceptor {
	return c.inters.Organization
}

func (c *OrganizationClient) mutate(ctx context.Context, m *OrganizationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Organization mutation op: %q", m.Op())
	}
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

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
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

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
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
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
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

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		BusinessUnit, Organization, User []ent.Hook
	}
	inters struct {
		BusinessUnit, Organization, User []ent.Interceptor
	}
)
