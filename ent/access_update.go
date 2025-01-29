// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"spl-access/ent/access"
	"spl-access/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessUpdate is the builder for updating Access entities.
type AccessUpdate struct {
	config
	hooks    []Hook
	mutation *AccessMutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (au *AccessUpdate) Where(ps ...predicate.Access) *AccessUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetRun sets the "run" field.
func (au *AccessUpdate) SetRun(s string) *AccessUpdate {
	au.mutation.SetRun(s)
	return au
}

// SetNillableRun sets the "run" field if the given value is not nil.
func (au *AccessUpdate) SetNillableRun(s *string) *AccessUpdate {
	if s != nil {
		au.SetRun(*s)
	}
	return au
}

// SetLocation sets the "location" field.
func (au *AccessUpdate) SetLocation(a access.Location) *AccessUpdate {
	au.mutation.SetLocation(a)
	return au
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (au *AccessUpdate) SetNillableLocation(a *access.Location) *AccessUpdate {
	if a != nil {
		au.SetLocation(*a)
	}
	return au
}

// SetEntryAt sets the "entry_at" field.
func (au *AccessUpdate) SetEntryAt(t time.Time) *AccessUpdate {
	au.mutation.SetEntryAt(t)
	return au
}

// SetNillableEntryAt sets the "entry_at" field if the given value is not nil.
func (au *AccessUpdate) SetNillableEntryAt(t *time.Time) *AccessUpdate {
	if t != nil {
		au.SetEntryAt(*t)
	}
	return au
}

// SetExitAt sets the "exit_at" field.
func (au *AccessUpdate) SetExitAt(t time.Time) *AccessUpdate {
	au.mutation.SetExitAt(t)
	return au
}

// SetNillableExitAt sets the "exit_at" field if the given value is not nil.
func (au *AccessUpdate) SetNillableExitAt(t *time.Time) *AccessUpdate {
	if t != nil {
		au.SetExitAt(*t)
	}
	return au
}

// ClearExitAt clears the value of the "exit_at" field.
func (au *AccessUpdate) ClearExitAt() *AccessUpdate {
	au.mutation.ClearExitAt()
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccessUpdate) SetCreatedAt(t time.Time) *AccessUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccessUpdate) SetNillableCreatedAt(t *time.Time) *AccessUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccessUpdate) SetUpdatedAt(t time.Time) *AccessUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the AccessMutation object of the builder.
func (au *AccessUpdate) Mutation() *AccessMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccessUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccessUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccessUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccessUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccessUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := access.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccessUpdate) check() error {
	if v, ok := au.mutation.Location(); ok {
		if err := access.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Access.location": %w`, err)}
		}
	}
	return nil
}

func (au *AccessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Run(); ok {
		_spec.SetField(access.FieldRun, field.TypeString, value)
	}
	if value, ok := au.mutation.Location(); ok {
		_spec.SetField(access.FieldLocation, field.TypeEnum, value)
	}
	if value, ok := au.mutation.EntryAt(); ok {
		_spec.SetField(access.FieldEntryAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.ExitAt(); ok {
		_spec.SetField(access.FieldExitAt, field.TypeTime, value)
	}
	if au.mutation.ExitAtCleared() {
		_spec.ClearField(access.FieldExitAt, field.TypeTime)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(access.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(access.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccessUpdateOne is the builder for updating a single Access entity.
type AccessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessMutation
}

// SetRun sets the "run" field.
func (auo *AccessUpdateOne) SetRun(s string) *AccessUpdateOne {
	auo.mutation.SetRun(s)
	return auo
}

// SetNillableRun sets the "run" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableRun(s *string) *AccessUpdateOne {
	if s != nil {
		auo.SetRun(*s)
	}
	return auo
}

// SetLocation sets the "location" field.
func (auo *AccessUpdateOne) SetLocation(a access.Location) *AccessUpdateOne {
	auo.mutation.SetLocation(a)
	return auo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableLocation(a *access.Location) *AccessUpdateOne {
	if a != nil {
		auo.SetLocation(*a)
	}
	return auo
}

// SetEntryAt sets the "entry_at" field.
func (auo *AccessUpdateOne) SetEntryAt(t time.Time) *AccessUpdateOne {
	auo.mutation.SetEntryAt(t)
	return auo
}

// SetNillableEntryAt sets the "entry_at" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableEntryAt(t *time.Time) *AccessUpdateOne {
	if t != nil {
		auo.SetEntryAt(*t)
	}
	return auo
}

// SetExitAt sets the "exit_at" field.
func (auo *AccessUpdateOne) SetExitAt(t time.Time) *AccessUpdateOne {
	auo.mutation.SetExitAt(t)
	return auo
}

// SetNillableExitAt sets the "exit_at" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableExitAt(t *time.Time) *AccessUpdateOne {
	if t != nil {
		auo.SetExitAt(*t)
	}
	return auo
}

// ClearExitAt clears the value of the "exit_at" field.
func (auo *AccessUpdateOne) ClearExitAt() *AccessUpdateOne {
	auo.mutation.ClearExitAt()
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccessUpdateOne) SetCreatedAt(t time.Time) *AccessUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableCreatedAt(t *time.Time) *AccessUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccessUpdateOne) SetUpdatedAt(t time.Time) *AccessUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the AccessMutation object of the builder.
func (auo *AccessUpdateOne) Mutation() *AccessMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (auo *AccessUpdateOne) Where(ps ...predicate.Access) *AccessUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccessUpdateOne) Select(field string, fields ...string) *AccessUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Access entity.
func (auo *AccessUpdateOne) Save(ctx context.Context) (*Access, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccessUpdateOne) SaveX(ctx context.Context) *Access {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccessUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccessUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccessUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := access.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccessUpdateOne) check() error {
	if v, ok := auo.mutation.Location(); ok {
		if err := access.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Access.location": %w`, err)}
		}
	}
	return nil
}

func (auo *AccessUpdateOne) sqlSave(ctx context.Context) (_node *Access, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Access.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, access.FieldID)
		for _, f := range fields {
			if !access.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != access.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Run(); ok {
		_spec.SetField(access.FieldRun, field.TypeString, value)
	}
	if value, ok := auo.mutation.Location(); ok {
		_spec.SetField(access.FieldLocation, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.EntryAt(); ok {
		_spec.SetField(access.FieldEntryAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.ExitAt(); ok {
		_spec.SetField(access.FieldExitAt, field.TypeTime, value)
	}
	if auo.mutation.ExitAtCleared() {
		_spec.ClearField(access.FieldExitAt, field.TypeTime)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(access.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(access.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Access{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
