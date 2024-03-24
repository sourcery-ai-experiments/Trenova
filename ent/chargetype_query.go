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
	"github.com/emoss08/trenova/ent/chargetype"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ChargeTypeQuery is the builder for querying ChargeType entities.
type ChargeTypeQuery struct {
	config
	ctx              *QueryContext
	order            []chargetype.OrderOption
	inters           []Interceptor
	predicates       []predicate.ChargeType
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChargeTypeQuery builder.
func (ctq *ChargeTypeQuery) Where(ps ...predicate.ChargeType) *ChargeTypeQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit the number of records to be returned by this query.
func (ctq *ChargeTypeQuery) Limit(limit int) *ChargeTypeQuery {
	ctq.ctx.Limit = &limit
	return ctq
}

// Offset to start from.
func (ctq *ChargeTypeQuery) Offset(offset int) *ChargeTypeQuery {
	ctq.ctx.Offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *ChargeTypeQuery) Unique(unique bool) *ChargeTypeQuery {
	ctq.ctx.Unique = &unique
	return ctq
}

// Order specifies how the records should be ordered.
func (ctq *ChargeTypeQuery) Order(o ...chargetype.OrderOption) *ChargeTypeQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (ctq *ChargeTypeQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chargetype.Table, chargetype.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, chargetype.BusinessUnitTable, chargetype.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (ctq *ChargeTypeQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chargetype.Table, chargetype.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, chargetype.OrganizationTable, chargetype.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChargeType entity from the query.
// Returns a *NotFoundError when no ChargeType was found.
func (ctq *ChargeTypeQuery) First(ctx context.Context) (*ChargeType, error) {
	nodes, err := ctq.Limit(1).All(setContextOp(ctx, ctq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chargetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *ChargeTypeQuery) FirstX(ctx context.Context) *ChargeType {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChargeType ID from the query.
// Returns a *NotFoundError when no ChargeType ID was found.
func (ctq *ChargeTypeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ctq.Limit(1).IDs(setContextOp(ctx, ctq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chargetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *ChargeTypeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChargeType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChargeType entity is found.
// Returns a *NotFoundError when no ChargeType entities are found.
func (ctq *ChargeTypeQuery) Only(ctx context.Context) (*ChargeType, error) {
	nodes, err := ctq.Limit(2).All(setContextOp(ctx, ctq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chargetype.Label}
	default:
		return nil, &NotSingularError{chargetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *ChargeTypeQuery) OnlyX(ctx context.Context) *ChargeType {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChargeType ID in the query.
// Returns a *NotSingularError when more than one ChargeType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *ChargeTypeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ctq.Limit(2).IDs(setContextOp(ctx, ctq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chargetype.Label}
	default:
		err = &NotSingularError{chargetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *ChargeTypeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChargeTypes.
func (ctq *ChargeTypeQuery) All(ctx context.Context) ([]*ChargeType, error) {
	ctx = setContextOp(ctx, ctq.ctx, "All")
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChargeType, *ChargeTypeQuery]()
	return withInterceptors[[]*ChargeType](ctx, ctq, qr, ctq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ctq *ChargeTypeQuery) AllX(ctx context.Context) []*ChargeType {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChargeType IDs.
func (ctq *ChargeTypeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ctq.ctx.Unique == nil && ctq.path != nil {
		ctq.Unique(true)
	}
	ctx = setContextOp(ctx, ctq.ctx, "IDs")
	if err = ctq.Select(chargetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *ChargeTypeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *ChargeTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Count")
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ctq, querierCount[*ChargeTypeQuery](), ctq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *ChargeTypeQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *ChargeTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Exist")
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *ChargeTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChargeTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *ChargeTypeQuery) Clone() *ChargeTypeQuery {
	if ctq == nil {
		return nil
	}
	return &ChargeTypeQuery{
		config:           ctq.config,
		ctx:              ctq.ctx.Clone(),
		order:            append([]chargetype.OrderOption{}, ctq.order...),
		inters:           append([]Interceptor{}, ctq.inters...),
		predicates:       append([]predicate.ChargeType{}, ctq.predicates...),
		withBusinessUnit: ctq.withBusinessUnit.Clone(),
		withOrganization: ctq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ChargeTypeQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *ChargeTypeQuery {
	query := (&BusinessUnitClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withBusinessUnit = query
	return ctq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ChargeTypeQuery) WithOrganization(opts ...func(*OrganizationQuery)) *ChargeTypeQuery {
	query := (&OrganizationClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withOrganization = query
	return ctq
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
//	client.ChargeType.Query().
//		GroupBy(chargetype.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctq *ChargeTypeQuery) GroupBy(field string, fields ...string) *ChargeTypeGroupBy {
	ctq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChargeTypeGroupBy{build: ctq}
	grbuild.flds = &ctq.ctx.Fields
	grbuild.label = chargetype.Label
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
//	client.ChargeType.Query().
//		Select(chargetype.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (ctq *ChargeTypeQuery) Select(fields ...string) *ChargeTypeSelect {
	ctq.ctx.Fields = append(ctq.ctx.Fields, fields...)
	sbuild := &ChargeTypeSelect{ChargeTypeQuery: ctq}
	sbuild.label = chargetype.Label
	sbuild.flds, sbuild.scan = &ctq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChargeTypeSelect configured with the given aggregations.
func (ctq *ChargeTypeQuery) Aggregate(fns ...AggregateFunc) *ChargeTypeSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *ChargeTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ctq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ctq); err != nil {
				return err
			}
		}
	}
	for _, f := range ctq.ctx.Fields {
		if !chargetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *ChargeTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChargeType, error) {
	var (
		nodes       = []*ChargeType{}
		_spec       = ctq.querySpec()
		loadedTypes = [2]bool{
			ctq.withBusinessUnit != nil,
			ctq.withOrganization != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChargeType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChargeType{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctq.withBusinessUnit; query != nil {
		if err := ctq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *ChargeType, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := ctq.withOrganization; query != nil {
		if err := ctq.loadOrganization(ctx, query, nodes, nil,
			func(n *ChargeType, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctq *ChargeTypeQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*ChargeType, init func(*ChargeType), assign func(*ChargeType, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ChargeType)
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
func (ctq *ChargeTypeQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*ChargeType, init func(*ChargeType), assign func(*ChargeType, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ChargeType)
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

func (ctq *ChargeTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	_spec.Node.Columns = ctq.ctx.Fields
	if len(ctq.ctx.Fields) > 0 {
		_spec.Unique = ctq.ctx.Unique != nil && *ctq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *ChargeTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chargetype.Table, chargetype.Columns, sqlgraph.NewFieldSpec(chargetype.FieldID, field.TypeUUID))
	_spec.From = ctq.sql
	if unique := ctq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ctq.path != nil {
		_spec.Unique = true
	}
	if fields := ctq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chargetype.FieldID)
		for i := range fields {
			if fields[i] != chargetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ctq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(chargetype.FieldBusinessUnitID)
		}
		if ctq.withOrganization != nil {
			_spec.Node.AddColumnOnce(chargetype.FieldOrganizationID)
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *ChargeTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(chargetype.Table)
	columns := ctq.ctx.Fields
	if len(columns) == 0 {
		columns = chargetype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.ctx.Unique != nil && *ctq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ctq.modifiers {
		m(selector)
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ctq *ChargeTypeQuery) Modify(modifiers ...func(s *sql.Selector)) *ChargeTypeSelect {
	ctq.modifiers = append(ctq.modifiers, modifiers...)
	return ctq.Select()
}

// ChargeTypeGroupBy is the group-by builder for ChargeType entities.
type ChargeTypeGroupBy struct {
	selector
	build *ChargeTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *ChargeTypeGroupBy) Aggregate(fns ...AggregateFunc) *ChargeTypeGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the selector query and scans the result into the given value.
func (ctgb *ChargeTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ctgb.build.ctx, "GroupBy")
	if err := ctgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChargeTypeQuery, *ChargeTypeGroupBy](ctx, ctgb.build, ctgb, ctgb.build.inters, v)
}

func (ctgb *ChargeTypeGroupBy) sqlScan(ctx context.Context, root *ChargeTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ctgb.flds)+len(ctgb.fns))
		for _, f := range *ctgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ctgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChargeTypeSelect is the builder for selecting fields of ChargeType entities.
type ChargeTypeSelect struct {
	*ChargeTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *ChargeTypeSelect) Aggregate(fns ...AggregateFunc) *ChargeTypeSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *ChargeTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cts.ctx, "Select")
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChargeTypeQuery, *ChargeTypeSelect](ctx, cts.ChargeTypeQuery, cts, cts.inters, v)
}

func (cts *ChargeTypeSelect) sqlScan(ctx context.Context, root *ChargeTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cts *ChargeTypeSelect) Modify(modifiers ...func(s *sql.Selector)) *ChargeTypeSelect {
	cts.modifiers = append(cts.modifiers, modifiers...)
	return cts
}
