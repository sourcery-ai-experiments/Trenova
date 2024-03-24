// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/session"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks     []Hook
	mutation  *SessionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetData sets the "data" field.
func (su *SessionUpdate) SetData(s string) *SessionUpdate {
	su.mutation.SetData(s)
	return su
}

// SetNillableData sets the "data" field if the given value is not nil.
func (su *SessionUpdate) SetNillableData(s *string) *SessionUpdate {
	if s != nil {
		su.SetData(*s)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SessionUpdate) SetUpdatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUpdatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// SetExpiresAt sets the "expires_at" field.
func (su *SessionUpdate) SetExpiresAt(t time.Time) *SessionUpdate {
	su.mutation.SetExpiresAt(t)
	return su
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableExpiresAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetExpiresAt(*t)
	}
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SessionUpdate) check() error {
	if v, ok := su.mutation.Data(); ok {
		if err := session.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "Session.data": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SessionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SessionUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Data(); ok {
		_spec.SetField(session.FieldData, field.TypeString, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SessionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetData sets the "data" field.
func (suo *SessionUpdateOne) SetData(s string) *SessionUpdateOne {
	suo.mutation.SetData(s)
	return suo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableData(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetData(*s)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SessionUpdateOne) SetUpdatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUpdatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// SetExpiresAt sets the "expires_at" field.
func (suo *SessionUpdateOne) SetExpiresAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetExpiresAt(t)
	return suo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableExpiresAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetExpiresAt(*t)
	}
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SessionUpdateOne) check() error {
	if v, ok := suo.mutation.Data(); ok {
		if err := session.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "Session.data": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SessionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SessionUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Data(); ok {
		_spec.SetField(session.FieldData, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
