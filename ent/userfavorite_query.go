// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/user"
	"github.com/emoss08/trenova/ent/userfavorite"
	"github.com/google/uuid"
)

// UserFavoriteQuery is the builder for querying UserFavorite entities.
type UserFavoriteQuery struct {
	config
	ctx              *QueryContext
	order            []userfavorite.OrderOption
	inters           []Interceptor
	predicates       []predicate.UserFavorite
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withUser         *UserQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserFavoriteQuery builder.
func (ufq *UserFavoriteQuery) Where(ps ...predicate.UserFavorite) *UserFavoriteQuery {
	ufq.predicates = append(ufq.predicates, ps...)
	return ufq
}

// Limit the number of records to be returned by this query.
func (ufq *UserFavoriteQuery) Limit(limit int) *UserFavoriteQuery {
	ufq.ctx.Limit = &limit
	return ufq
}

// Offset to start from.
func (ufq *UserFavoriteQuery) Offset(offset int) *UserFavoriteQuery {
	ufq.ctx.Offset = &offset
	return ufq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufq *UserFavoriteQuery) Unique(unique bool) *UserFavoriteQuery {
	ufq.ctx.Unique = &unique
	return ufq
}

// Order specifies how the records should be ordered.
func (ufq *UserFavoriteQuery) Order(o ...userfavorite.OrderOption) *UserFavoriteQuery {
	ufq.order = append(ufq.order, o...)
	return ufq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (ufq *UserFavoriteQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userfavorite.Table, userfavorite.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userfavorite.BusinessUnitTable, userfavorite.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (ufq *UserFavoriteQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userfavorite.Table, userfavorite.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userfavorite.OrganizationTable, userfavorite.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ufq *UserFavoriteQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userfavorite.Table, userfavorite.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userfavorite.UserTable, userfavorite.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserFavorite entity from the query.
// Returns a *NotFoundError when no UserFavorite was found.
func (ufq *UserFavoriteQuery) First(ctx context.Context) (*UserFavorite, error) {
	nodes, err := ufq.Limit(1).All(setContextOp(ctx, ufq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userfavorite.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufq *UserFavoriteQuery) FirstX(ctx context.Context) *UserFavorite {
	node, err := ufq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserFavorite ID from the query.
// Returns a *NotFoundError when no UserFavorite ID was found.
func (ufq *UserFavoriteQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ufq.Limit(1).IDs(setContextOp(ctx, ufq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userfavorite.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufq *UserFavoriteQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ufq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserFavorite entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserFavorite entity is found.
// Returns a *NotFoundError when no UserFavorite entities are found.
func (ufq *UserFavoriteQuery) Only(ctx context.Context) (*UserFavorite, error) {
	nodes, err := ufq.Limit(2).All(setContextOp(ctx, ufq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userfavorite.Label}
	default:
		return nil, &NotSingularError{userfavorite.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufq *UserFavoriteQuery) OnlyX(ctx context.Context) *UserFavorite {
	node, err := ufq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserFavorite ID in the query.
// Returns a *NotSingularError when more than one UserFavorite ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufq *UserFavoriteQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ufq.Limit(2).IDs(setContextOp(ctx, ufq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userfavorite.Label}
	default:
		err = &NotSingularError{userfavorite.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufq *UserFavoriteQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ufq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserFavorites.
func (ufq *UserFavoriteQuery) All(ctx context.Context) ([]*UserFavorite, error) {
	ctx = setContextOp(ctx, ufq.ctx, "All")
	if err := ufq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserFavorite, *UserFavoriteQuery]()
	return withInterceptors[[]*UserFavorite](ctx, ufq, qr, ufq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufq *UserFavoriteQuery) AllX(ctx context.Context) []*UserFavorite {
	nodes, err := ufq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserFavorite IDs.
func (ufq *UserFavoriteQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ufq.ctx.Unique == nil && ufq.path != nil {
		ufq.Unique(true)
	}
	ctx = setContextOp(ctx, ufq.ctx, "IDs")
	if err = ufq.Select(userfavorite.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufq *UserFavoriteQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ufq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufq *UserFavoriteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufq.ctx, "Count")
	if err := ufq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufq, querierCount[*UserFavoriteQuery](), ufq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufq *UserFavoriteQuery) CountX(ctx context.Context) int {
	count, err := ufq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufq *UserFavoriteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufq.ctx, "Exist")
	switch _, err := ufq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufq *UserFavoriteQuery) ExistX(ctx context.Context) bool {
	exist, err := ufq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserFavoriteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufq *UserFavoriteQuery) Clone() *UserFavoriteQuery {
	if ufq == nil {
		return nil
	}
	return &UserFavoriteQuery{
		config:           ufq.config,
		ctx:              ufq.ctx.Clone(),
		order:            append([]userfavorite.OrderOption{}, ufq.order...),
		inters:           append([]Interceptor{}, ufq.inters...),
		predicates:       append([]predicate.UserFavorite{}, ufq.predicates...),
		withBusinessUnit: ufq.withBusinessUnit.Clone(),
		withOrganization: ufq.withOrganization.Clone(),
		withUser:         ufq.withUser.Clone(),
		// clone intermediate query.
		sql:  ufq.sql.Clone(),
		path: ufq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UserFavoriteQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *UserFavoriteQuery {
	query := (&BusinessUnitClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withBusinessUnit = query
	return ufq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UserFavoriteQuery) WithOrganization(opts ...func(*OrganizationQuery)) *UserFavoriteQuery {
	query := (&OrganizationClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withOrganization = query
	return ufq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UserFavoriteQuery) WithUser(opts ...func(*UserQuery)) *UserFavoriteQuery {
	query := (&UserClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withUser = query
	return ufq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserFavorite.Query().
//		GroupBy(userfavorite.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufq *UserFavoriteQuery) GroupBy(field string, fields ...string) *UserFavoriteGroupBy {
	ufq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserFavoriteGroupBy{build: ufq}
	grbuild.flds = &ufq.ctx.Fields
	grbuild.label = userfavorite.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//	}
//
//	client.UserFavorite.Query().
//		Select(userfavorite.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (ufq *UserFavoriteQuery) Select(fields ...string) *UserFavoriteSelect {
	ufq.ctx.Fields = append(ufq.ctx.Fields, fields...)
	sbuild := &UserFavoriteSelect{UserFavoriteQuery: ufq}
	sbuild.label = userfavorite.Label
	sbuild.flds, sbuild.scan = &ufq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserFavoriteSelect configured with the given aggregations.
func (ufq *UserFavoriteQuery) Aggregate(fns ...AggregateFunc) *UserFavoriteSelect {
	return ufq.Select().Aggregate(fns...)
}

func (ufq *UserFavoriteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufq.ctx.Fields {
		if !userfavorite.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufq.path != nil {
		prev, err := ufq.path(ctx)
		if err != nil {
			return err
		}
		ufq.sql = prev
	}
	return nil
}

func (ufq *UserFavoriteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserFavorite, error) {
	var (
		nodes       = []*UserFavorite{}
		_spec       = ufq.querySpec()
		loadedTypes = [3]bool{
			ufq.withBusinessUnit != nil,
			ufq.withOrganization != nil,
			ufq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserFavorite).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserFavorite{config: ufq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ufq.modifiers) > 0 {
		_spec.Modifiers = ufq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ufq.withBusinessUnit; query != nil {
		if err := ufq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *UserFavorite, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := ufq.withOrganization; query != nil {
		if err := ufq.loadOrganization(ctx, query, nodes, nil,
			func(n *UserFavorite, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := ufq.withUser; query != nil {
		if err := ufq.loadUser(ctx, query, nodes, nil,
			func(n *UserFavorite, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ufq *UserFavoriteQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*UserFavorite, init func(*UserFavorite), assign func(*UserFavorite, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserFavorite)
	for i := range nodes {
		fk := nodes[i].BusinessUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(businessunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ufq *UserFavoriteQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*UserFavorite, init func(*UserFavorite), assign func(*UserFavorite, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserFavorite)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ufq *UserFavoriteQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserFavorite, init func(*UserFavorite), assign func(*UserFavorite, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserFavorite)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ufq *UserFavoriteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufq.querySpec()
	if len(ufq.modifiers) > 0 {
		_spec.Modifiers = ufq.modifiers
	}
	_spec.Node.Columns = ufq.ctx.Fields
	if len(ufq.ctx.Fields) > 0 {
		_spec.Unique = ufq.ctx.Unique != nil && *ufq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ufq.driver, _spec)
}

func (ufq *UserFavoriteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userfavorite.Table, userfavorite.Columns, sqlgraph.NewFieldSpec(userfavorite.FieldID, field.TypeUUID))
	_spec.From = ufq.sql
	if unique := ufq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufq.path != nil {
		_spec.Unique = true
	}
	if fields := ufq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfavorite.FieldID)
		for i := range fields {
			if fields[i] != userfavorite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ufq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(userfavorite.FieldBusinessUnitID)
		}
		if ufq.withOrganization != nil {
			_spec.Node.AddColumnOnce(userfavorite.FieldOrganizationID)
		}
		if ufq.withUser != nil {
			_spec.Node.AddColumnOnce(userfavorite.FieldUserID)
		}
	}
	if ps := ufq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufq *UserFavoriteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufq.driver.Dialect())
	t1 := builder.Table(userfavorite.Table)
	columns := ufq.ctx.Fields
	if len(columns) == 0 {
		columns = userfavorite.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufq.sql != nil {
		selector = ufq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufq.ctx.Unique != nil && *ufq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ufq.modifiers {
		m(selector)
	}
	for _, p := range ufq.predicates {
		p(selector)
	}
	for _, p := range ufq.order {
		p(selector)
	}
	if offset := ufq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ufq *UserFavoriteQuery) Modify(modifiers ...func(s *sql.Selector)) *UserFavoriteSelect {
	ufq.modifiers = append(ufq.modifiers, modifiers...)
	return ufq.Select()
}

// UserFavoriteGroupBy is the group-by builder for UserFavorite entities.
type UserFavoriteGroupBy struct {
	selector
	build *UserFavoriteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufgb *UserFavoriteGroupBy) Aggregate(fns ...AggregateFunc) *UserFavoriteGroupBy {
	ufgb.fns = append(ufgb.fns, fns...)
	return ufgb
}

// Scan applies the selector query and scans the result into the given value.
func (ufgb *UserFavoriteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufgb.build.ctx, "GroupBy")
	if err := ufgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserFavoriteQuery, *UserFavoriteGroupBy](ctx, ufgb.build, ufgb, ufgb.build.inters, v)
}

func (ufgb *UserFavoriteGroupBy) sqlScan(ctx context.Context, root *UserFavoriteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufgb.fns))
	for _, fn := range ufgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufgb.flds)+len(ufgb.fns))
		for _, f := range *ufgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserFavoriteSelect is the builder for selecting fields of UserFavorite entities.
type UserFavoriteSelect struct {
	*UserFavoriteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufs *UserFavoriteSelect) Aggregate(fns ...AggregateFunc) *UserFavoriteSelect {
	ufs.fns = append(ufs.fns, fns...)
	return ufs
}

// Scan applies the selector query and scans the result into the given value.
func (ufs *UserFavoriteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufs.ctx, "Select")
	if err := ufs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserFavoriteQuery, *UserFavoriteSelect](ctx, ufs.UserFavoriteQuery, ufs, ufs.inters, v)
}

func (ufs *UserFavoriteSelect) sqlScan(ctx context.Context, root *UserFavoriteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufs.fns))
	for _, fn := range ufs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ufs *UserFavoriteSelect) Modify(modifiers ...func(s *sql.Selector)) *UserFavoriteSelect {
	ufs.modifiers = append(ufs.modifiers, modifiers...)
	return ufs
}
