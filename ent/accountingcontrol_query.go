// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/accountingcontrol"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// AccountingControlQuery is the builder for querying AccountingControl entities.
type AccountingControlQuery struct {
	config
	ctx                   *QueryContext
	order                 []accountingcontrol.OrderOption
	inters                []Interceptor
	predicates            []predicate.AccountingControl
	withOrganization      *OrganizationQuery
	withBusinessUnit      *BusinessUnitQuery
	withDefaultRevAccount *GeneralLedgerAccountQuery
	withDefaultExpAccount *GeneralLedgerAccountQuery
	withFKs               bool
	modifiers             []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountingControlQuery builder.
func (acq *AccountingControlQuery) Where(ps ...predicate.AccountingControl) *AccountingControlQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AccountingControlQuery) Limit(limit int) *AccountingControlQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AccountingControlQuery) Offset(offset int) *AccountingControlQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AccountingControlQuery) Unique(unique bool) *AccountingControlQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AccountingControlQuery) Order(o ...accountingcontrol.OrderOption) *AccountingControlQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryOrganization chains the current query on the "organization" edge.
func (acq *AccountingControlQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountingcontrol.Table, accountingcontrol.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, accountingcontrol.OrganizationTable, accountingcontrol.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (acq *AccountingControlQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountingcontrol.Table, accountingcontrol.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, accountingcontrol.BusinessUnitTable, accountingcontrol.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDefaultRevAccount chains the current query on the "default_rev_account" edge.
func (acq *AccountingControlQuery) QueryDefaultRevAccount() *GeneralLedgerAccountQuery {
	query := (&GeneralLedgerAccountClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountingcontrol.Table, accountingcontrol.FieldID, selector),
			sqlgraph.To(generalledgeraccount.Table, generalledgeraccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, accountingcontrol.DefaultRevAccountTable, accountingcontrol.DefaultRevAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDefaultExpAccount chains the current query on the "default_exp_account" edge.
func (acq *AccountingControlQuery) QueryDefaultExpAccount() *GeneralLedgerAccountQuery {
	query := (&GeneralLedgerAccountClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountingcontrol.Table, accountingcontrol.FieldID, selector),
			sqlgraph.To(generalledgeraccount.Table, generalledgeraccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, accountingcontrol.DefaultExpAccountTable, accountingcontrol.DefaultExpAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccountingControl entity from the query.
// Returns a *NotFoundError when no AccountingControl was found.
func (acq *AccountingControlQuery) First(ctx context.Context) (*AccountingControl, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountingcontrol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AccountingControlQuery) FirstX(ctx context.Context) *AccountingControl {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountingControl ID from the query.
// Returns a *NotFoundError when no AccountingControl ID was found.
func (acq *AccountingControlQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountingcontrol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AccountingControlQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountingControl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountingControl entity is found.
// Returns a *NotFoundError when no AccountingControl entities are found.
func (acq *AccountingControlQuery) Only(ctx context.Context) (*AccountingControl, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountingcontrol.Label}
	default:
		return nil, &NotSingularError{accountingcontrol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AccountingControlQuery) OnlyX(ctx context.Context) *AccountingControl {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountingControl ID in the query.
// Returns a *NotSingularError when more than one AccountingControl ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AccountingControlQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountingcontrol.Label}
	default:
		err = &NotSingularError{accountingcontrol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AccountingControlQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountingControls.
func (acq *AccountingControlQuery) All(ctx context.Context) ([]*AccountingControl, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccountingControl, *AccountingControlQuery]()
	return withInterceptors[[]*AccountingControl](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AccountingControlQuery) AllX(ctx context.Context) []*AccountingControl {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountingControl IDs.
func (acq *AccountingControlQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(accountingcontrol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AccountingControlQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AccountingControlQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AccountingControlQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AccountingControlQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AccountingControlQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AccountingControlQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountingControlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AccountingControlQuery) Clone() *AccountingControlQuery {
	if acq == nil {
		return nil
	}
	return &AccountingControlQuery{
		config:                acq.config,
		ctx:                   acq.ctx.Clone(),
		order:                 append([]accountingcontrol.OrderOption{}, acq.order...),
		inters:                append([]Interceptor{}, acq.inters...),
		predicates:            append([]predicate.AccountingControl{}, acq.predicates...),
		withOrganization:      acq.withOrganization.Clone(),
		withBusinessUnit:      acq.withBusinessUnit.Clone(),
		withDefaultRevAccount: acq.withDefaultRevAccount.Clone(),
		withDefaultExpAccount: acq.withDefaultExpAccount.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccountingControlQuery) WithOrganization(opts ...func(*OrganizationQuery)) *AccountingControlQuery {
	query := (&OrganizationClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withOrganization = query
	return acq
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccountingControlQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *AccountingControlQuery {
	query := (&BusinessUnitClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withBusinessUnit = query
	return acq
}

// WithDefaultRevAccount tells the query-builder to eager-load the nodes that are connected to
// the "default_rev_account" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccountingControlQuery) WithDefaultRevAccount(opts ...func(*GeneralLedgerAccountQuery)) *AccountingControlQuery {
	query := (&GeneralLedgerAccountClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withDefaultRevAccount = query
	return acq
}

// WithDefaultExpAccount tells the query-builder to eager-load the nodes that are connected to
// the "default_exp_account" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccountingControlQuery) WithDefaultExpAccount(opts ...func(*GeneralLedgerAccountQuery)) *AccountingControlQuery {
	query := (&GeneralLedgerAccountClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withDefaultExpAccount = query
	return acq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccountingControl.Query().
//		GroupBy(accountingcontrol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *AccountingControlQuery) GroupBy(field string, fields ...string) *AccountingControlGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountingControlGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = accountingcontrol.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//	}
//
//	client.AccountingControl.Query().
//		Select(accountingcontrol.FieldCreatedAt).
//		Scan(ctx, &v)
func (acq *AccountingControlQuery) Select(fields ...string) *AccountingControlSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AccountingControlSelect{AccountingControlQuery: acq}
	sbuild.label = accountingcontrol.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountingControlSelect configured with the given aggregations.
func (acq *AccountingControlQuery) Aggregate(fns ...AggregateFunc) *AccountingControlSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AccountingControlQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !accountingcontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AccountingControlQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountingControl, error) {
	var (
		nodes       = []*AccountingControl{}
		withFKs     = acq.withFKs
		_spec       = acq.querySpec()
		loadedTypes = [4]bool{
			acq.withOrganization != nil,
			acq.withBusinessUnit != nil,
			acq.withDefaultRevAccount != nil,
			acq.withDefaultExpAccount != nil,
		}
	)
	if acq.withOrganization != nil || acq.withBusinessUnit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, accountingcontrol.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountingControl).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountingControl{config: acq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acq.withOrganization; query != nil {
		if err := acq.loadOrganization(ctx, query, nodes, nil,
			func(n *AccountingControl, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := acq.withBusinessUnit; query != nil {
		if err := acq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *AccountingControl, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := acq.withDefaultRevAccount; query != nil {
		if err := acq.loadDefaultRevAccount(ctx, query, nodes, nil,
			func(n *AccountingControl, e *GeneralLedgerAccount) { n.Edges.DefaultRevAccount = e }); err != nil {
			return nil, err
		}
	}
	if query := acq.withDefaultExpAccount; query != nil {
		if err := acq.loadDefaultExpAccount(ctx, query, nodes, nil,
			func(n *AccountingControl, e *GeneralLedgerAccount) { n.Edges.DefaultExpAccount = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *AccountingControlQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*AccountingControl, init func(*AccountingControl), assign func(*AccountingControl, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AccountingControl)
	for i := range nodes {
		if nodes[i].organization_id == nil {
			continue
		}
		fk := *nodes[i].organization_id
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
func (acq *AccountingControlQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*AccountingControl, init func(*AccountingControl), assign func(*AccountingControl, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AccountingControl)
	for i := range nodes {
		if nodes[i].business_unit_id == nil {
			continue
		}
		fk := *nodes[i].business_unit_id
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
func (acq *AccountingControlQuery) loadDefaultRevAccount(ctx context.Context, query *GeneralLedgerAccountQuery, nodes []*AccountingControl, init func(*AccountingControl), assign func(*AccountingControl, *GeneralLedgerAccount)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AccountingControl)
	for i := range nodes {
		if nodes[i].DefaultRevAccountID == nil {
			continue
		}
		fk := *nodes[i].DefaultRevAccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(generalledgeraccount.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "default_rev_account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (acq *AccountingControlQuery) loadDefaultExpAccount(ctx context.Context, query *GeneralLedgerAccountQuery, nodes []*AccountingControl, init func(*AccountingControl), assign func(*AccountingControl, *GeneralLedgerAccount)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AccountingControl)
	for i := range nodes {
		if nodes[i].DefaultExpAccountID == nil {
			continue
		}
		fk := *nodes[i].DefaultExpAccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(generalledgeraccount.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "default_exp_account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (acq *AccountingControlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AccountingControlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accountingcontrol.Table, accountingcontrol.Columns, sqlgraph.NewFieldSpec(accountingcontrol.FieldID, field.TypeUUID))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountingcontrol.FieldID)
		for i := range fields {
			if fields[i] != accountingcontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if acq.withDefaultRevAccount != nil {
			_spec.Node.AddColumnOnce(accountingcontrol.FieldDefaultRevAccountID)
		}
		if acq.withDefaultExpAccount != nil {
			_spec.Node.AddColumnOnce(accountingcontrol.FieldDefaultExpAccountID)
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AccountingControlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(accountingcontrol.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = accountingcontrol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range acq.modifiers {
		m(selector)
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acq *AccountingControlQuery) Modify(modifiers ...func(s *sql.Selector)) *AccountingControlSelect {
	acq.modifiers = append(acq.modifiers, modifiers...)
	return acq.Select()
}

// AccountingControlGroupBy is the group-by builder for AccountingControl entities.
type AccountingControlGroupBy struct {
	selector
	build *AccountingControlQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AccountingControlGroupBy) Aggregate(fns ...AggregateFunc) *AccountingControlGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AccountingControlGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountingControlQuery, *AccountingControlGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AccountingControlGroupBy) sqlScan(ctx context.Context, root *AccountingControlQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountingControlSelect is the builder for selecting fields of AccountingControl entities.
type AccountingControlSelect struct {
	*AccountingControlQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AccountingControlSelect) Aggregate(fns ...AggregateFunc) *AccountingControlSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AccountingControlSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountingControlQuery, *AccountingControlSelect](ctx, acs.AccountingControlQuery, acs, acs.inters, v)
}

func (acs *AccountingControlSelect) sqlScan(ctx context.Context, root *AccountingControlQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acs *AccountingControlSelect) Modify(modifiers ...func(s *sql.Selector)) *AccountingControlSelect {
	acs.modifiers = append(acs.modifiers, modifiers...)
	return acs
}
