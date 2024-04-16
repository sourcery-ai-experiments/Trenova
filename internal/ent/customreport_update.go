// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/customreport"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// CustomReportUpdate is the builder for updating CustomReport entities.
type CustomReportUpdate struct {
	config
	hooks     []Hook
	mutation  *CustomReportMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CustomReportUpdate builder.
func (cru *CustomReportUpdate) Where(ps ...predicate.CustomReport) *CustomReportUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetUpdatedAt sets the "updated_at" field.
func (cru *CustomReportUpdate) SetUpdatedAt(t time.Time) *CustomReportUpdate {
	cru.mutation.SetUpdatedAt(t)
	return cru
}

// SetVersion sets the "version" field.
func (cru *CustomReportUpdate) SetVersion(i int) *CustomReportUpdate {
	cru.mutation.ResetVersion()
	cru.mutation.SetVersion(i)
	return cru
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cru *CustomReportUpdate) SetNillableVersion(i *int) *CustomReportUpdate {
	if i != nil {
		cru.SetVersion(*i)
	}
	return cru
}

// AddVersion adds i to the "version" field.
func (cru *CustomReportUpdate) AddVersion(i int) *CustomReportUpdate {
	cru.mutation.AddVersion(i)
	return cru
}

// SetName sets the "name" field.
func (cru *CustomReportUpdate) SetName(s string) *CustomReportUpdate {
	cru.mutation.SetName(s)
	return cru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cru *CustomReportUpdate) SetNillableName(s *string) *CustomReportUpdate {
	if s != nil {
		cru.SetName(*s)
	}
	return cru
}

// SetDescription sets the "description" field.
func (cru *CustomReportUpdate) SetDescription(s string) *CustomReportUpdate {
	cru.mutation.SetDescription(s)
	return cru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cru *CustomReportUpdate) SetNillableDescription(s *string) *CustomReportUpdate {
	if s != nil {
		cru.SetDescription(*s)
	}
	return cru
}

// ClearDescription clears the value of the "description" field.
func (cru *CustomReportUpdate) ClearDescription() *CustomReportUpdate {
	cru.mutation.ClearDescription()
	return cru
}

// SetTable sets the "table" field.
func (cru *CustomReportUpdate) SetTable(s string) *CustomReportUpdate {
	cru.mutation.SetTable(s)
	return cru
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cru *CustomReportUpdate) SetNillableTable(s *string) *CustomReportUpdate {
	if s != nil {
		cru.SetTable(*s)
	}
	return cru
}

// ClearTable clears the value of the "table" field.
func (cru *CustomReportUpdate) ClearTable() *CustomReportUpdate {
	cru.mutation.ClearTable()
	return cru
}

// Mutation returns the CustomReportMutation object of the builder.
func (cru *CustomReportUpdate) Mutation() *CustomReportMutation {
	return cru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CustomReportUpdate) Save(ctx context.Context) (int, error) {
	cru.defaults()
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CustomReportUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CustomReportUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CustomReportUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cru *CustomReportUpdate) defaults() {
	if _, ok := cru.mutation.UpdatedAt(); !ok {
		v := customreport.UpdateDefaultUpdatedAt()
		cru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *CustomReportUpdate) check() error {
	if v, ok := cru.mutation.Name(); ok {
		if err := customreport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CustomReport.name": %w`, err)}
		}
	}
	if _, ok := cru.mutation.BusinessUnitID(); cru.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CustomReport.business_unit"`)
	}
	if _, ok := cru.mutation.OrganizationID(); cru.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CustomReport.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cru *CustomReportUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomReportUpdate {
	cru.modifiers = append(cru.modifiers, modifiers...)
	return cru
}

func (cru *CustomReportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customreport.Table, customreport.Columns, sqlgraph.NewFieldSpec(customreport.FieldID, field.TypeUUID))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.UpdatedAt(); ok {
		_spec.SetField(customreport.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cru.mutation.Version(); ok {
		_spec.SetField(customreport.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedVersion(); ok {
		_spec.AddField(customreport.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cru.mutation.Name(); ok {
		_spec.SetField(customreport.FieldName, field.TypeString, value)
	}
	if value, ok := cru.mutation.Description(); ok {
		_spec.SetField(customreport.FieldDescription, field.TypeString, value)
	}
	if cru.mutation.DescriptionCleared() {
		_spec.ClearField(customreport.FieldDescription, field.TypeString)
	}
	if value, ok := cru.mutation.Table(); ok {
		_spec.SetField(customreport.FieldTable, field.TypeString, value)
	}
	if cru.mutation.TableCleared() {
		_spec.ClearField(customreport.FieldTable, field.TypeString)
	}
	_spec.AddModifiers(cru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customreport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CustomReportUpdateOne is the builder for updating a single CustomReport entity.
type CustomReportUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CustomReportMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cruo *CustomReportUpdateOne) SetUpdatedAt(t time.Time) *CustomReportUpdateOne {
	cruo.mutation.SetUpdatedAt(t)
	return cruo
}

// SetVersion sets the "version" field.
func (cruo *CustomReportUpdateOne) SetVersion(i int) *CustomReportUpdateOne {
	cruo.mutation.ResetVersion()
	cruo.mutation.SetVersion(i)
	return cruo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cruo *CustomReportUpdateOne) SetNillableVersion(i *int) *CustomReportUpdateOne {
	if i != nil {
		cruo.SetVersion(*i)
	}
	return cruo
}

// AddVersion adds i to the "version" field.
func (cruo *CustomReportUpdateOne) AddVersion(i int) *CustomReportUpdateOne {
	cruo.mutation.AddVersion(i)
	return cruo
}

// SetName sets the "name" field.
func (cruo *CustomReportUpdateOne) SetName(s string) *CustomReportUpdateOne {
	cruo.mutation.SetName(s)
	return cruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cruo *CustomReportUpdateOne) SetNillableName(s *string) *CustomReportUpdateOne {
	if s != nil {
		cruo.SetName(*s)
	}
	return cruo
}

// SetDescription sets the "description" field.
func (cruo *CustomReportUpdateOne) SetDescription(s string) *CustomReportUpdateOne {
	cruo.mutation.SetDescription(s)
	return cruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cruo *CustomReportUpdateOne) SetNillableDescription(s *string) *CustomReportUpdateOne {
	if s != nil {
		cruo.SetDescription(*s)
	}
	return cruo
}

// ClearDescription clears the value of the "description" field.
func (cruo *CustomReportUpdateOne) ClearDescription() *CustomReportUpdateOne {
	cruo.mutation.ClearDescription()
	return cruo
}

// SetTable sets the "table" field.
func (cruo *CustomReportUpdateOne) SetTable(s string) *CustomReportUpdateOne {
	cruo.mutation.SetTable(s)
	return cruo
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cruo *CustomReportUpdateOne) SetNillableTable(s *string) *CustomReportUpdateOne {
	if s != nil {
		cruo.SetTable(*s)
	}
	return cruo
}

// ClearTable clears the value of the "table" field.
func (cruo *CustomReportUpdateOne) ClearTable() *CustomReportUpdateOne {
	cruo.mutation.ClearTable()
	return cruo
}

// Mutation returns the CustomReportMutation object of the builder.
func (cruo *CustomReportUpdateOne) Mutation() *CustomReportMutation {
	return cruo.mutation
}

// Where appends a list predicates to the CustomReportUpdate builder.
func (cruo *CustomReportUpdateOne) Where(ps ...predicate.CustomReport) *CustomReportUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CustomReportUpdateOne) Select(field string, fields ...string) *CustomReportUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CustomReport entity.
func (cruo *CustomReportUpdateOne) Save(ctx context.Context) (*CustomReport, error) {
	cruo.defaults()
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CustomReportUpdateOne) SaveX(ctx context.Context) *CustomReport {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CustomReportUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CustomReportUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cruo *CustomReportUpdateOne) defaults() {
	if _, ok := cruo.mutation.UpdatedAt(); !ok {
		v := customreport.UpdateDefaultUpdatedAt()
		cruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *CustomReportUpdateOne) check() error {
	if v, ok := cruo.mutation.Name(); ok {
		if err := customreport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CustomReport.name": %w`, err)}
		}
	}
	if _, ok := cruo.mutation.BusinessUnitID(); cruo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CustomReport.business_unit"`)
	}
	if _, ok := cruo.mutation.OrganizationID(); cruo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CustomReport.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cruo *CustomReportUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomReportUpdateOne {
	cruo.modifiers = append(cruo.modifiers, modifiers...)
	return cruo
}

func (cruo *CustomReportUpdateOne) sqlSave(ctx context.Context) (_node *CustomReport, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customreport.Table, customreport.Columns, sqlgraph.NewFieldSpec(customreport.FieldID, field.TypeUUID))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CustomReport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customreport.FieldID)
		for _, f := range fields {
			if !customreport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customreport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.UpdatedAt(); ok {
		_spec.SetField(customreport.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.Version(); ok {
		_spec.SetField(customreport.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedVersion(); ok {
		_spec.AddField(customreport.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.Name(); ok {
		_spec.SetField(customreport.FieldName, field.TypeString, value)
	}
	if value, ok := cruo.mutation.Description(); ok {
		_spec.SetField(customreport.FieldDescription, field.TypeString, value)
	}
	if cruo.mutation.DescriptionCleared() {
		_spec.ClearField(customreport.FieldDescription, field.TypeString)
	}
	if value, ok := cruo.mutation.Table(); ok {
		_spec.SetField(customreport.FieldTable, field.TypeString, value)
	}
	if cruo.mutation.TableCleared() {
		_spec.ClearField(customreport.FieldTable, field.TypeString)
	}
	_spec.AddModifiers(cruo.modifiers...)
	_node = &CustomReport{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customreport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
