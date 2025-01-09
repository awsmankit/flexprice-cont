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
	"github.com/flexprice/flexprice/ent/invoicesequence"
	"github.com/flexprice/flexprice/ent/predicate"
)

// InvoiceSequenceUpdate is the builder for updating InvoiceSequence entities.
type InvoiceSequenceUpdate struct {
	config
	hooks    []Hook
	mutation *InvoiceSequenceMutation
}

// Where appends a list predicates to the InvoiceSequenceUpdate builder.
func (isu *InvoiceSequenceUpdate) Where(ps ...predicate.InvoiceSequence) *InvoiceSequenceUpdate {
	isu.mutation.Where(ps...)
	return isu
}

// SetTenantID sets the "tenant_id" field.
func (isu *InvoiceSequenceUpdate) SetTenantID(s string) *InvoiceSequenceUpdate {
	isu.mutation.SetTenantID(s)
	return isu
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (isu *InvoiceSequenceUpdate) SetNillableTenantID(s *string) *InvoiceSequenceUpdate {
	if s != nil {
		isu.SetTenantID(*s)
	}
	return isu
}

// SetYearMonth sets the "year_month" field.
func (isu *InvoiceSequenceUpdate) SetYearMonth(s string) *InvoiceSequenceUpdate {
	isu.mutation.SetYearMonth(s)
	return isu
}

// SetNillableYearMonth sets the "year_month" field if the given value is not nil.
func (isu *InvoiceSequenceUpdate) SetNillableYearMonth(s *string) *InvoiceSequenceUpdate {
	if s != nil {
		isu.SetYearMonth(*s)
	}
	return isu
}

// SetLastValue sets the "last_value" field.
func (isu *InvoiceSequenceUpdate) SetLastValue(i int64) *InvoiceSequenceUpdate {
	isu.mutation.ResetLastValue()
	isu.mutation.SetLastValue(i)
	return isu
}

// SetNillableLastValue sets the "last_value" field if the given value is not nil.
func (isu *InvoiceSequenceUpdate) SetNillableLastValue(i *int64) *InvoiceSequenceUpdate {
	if i != nil {
		isu.SetLastValue(*i)
	}
	return isu
}

// AddLastValue adds i to the "last_value" field.
func (isu *InvoiceSequenceUpdate) AddLastValue(i int64) *InvoiceSequenceUpdate {
	isu.mutation.AddLastValue(i)
	return isu
}

// SetUpdatedAt sets the "updated_at" field.
func (isu *InvoiceSequenceUpdate) SetUpdatedAt(t time.Time) *InvoiceSequenceUpdate {
	isu.mutation.SetUpdatedAt(t)
	return isu
}

// Mutation returns the InvoiceSequenceMutation object of the builder.
func (isu *InvoiceSequenceUpdate) Mutation() *InvoiceSequenceMutation {
	return isu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (isu *InvoiceSequenceUpdate) Save(ctx context.Context) (int, error) {
	isu.defaults()
	return withHooks(ctx, isu.sqlSave, isu.mutation, isu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (isu *InvoiceSequenceUpdate) SaveX(ctx context.Context) int {
	affected, err := isu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (isu *InvoiceSequenceUpdate) Exec(ctx context.Context) error {
	_, err := isu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isu *InvoiceSequenceUpdate) ExecX(ctx context.Context) {
	if err := isu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (isu *InvoiceSequenceUpdate) defaults() {
	if _, ok := isu.mutation.UpdatedAt(); !ok {
		v := invoicesequence.UpdateDefaultUpdatedAt()
		isu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (isu *InvoiceSequenceUpdate) check() error {
	if v, ok := isu.mutation.TenantID(); ok {
		if err := invoicesequence.TenantIDValidator(v); err != nil {
			return &ValidationError{Name: "tenant_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceSequence.tenant_id": %w`, err)}
		}
	}
	if v, ok := isu.mutation.YearMonth(); ok {
		if err := invoicesequence.YearMonthValidator(v); err != nil {
			return &ValidationError{Name: "year_month", err: fmt.Errorf(`ent: validator failed for field "InvoiceSequence.year_month": %w`, err)}
		}
	}
	return nil
}

func (isu *InvoiceSequenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := isu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicesequence.Table, invoicesequence.Columns, sqlgraph.NewFieldSpec(invoicesequence.FieldID, field.TypeInt))
	if ps := isu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := isu.mutation.TenantID(); ok {
		_spec.SetField(invoicesequence.FieldTenantID, field.TypeString, value)
	}
	if value, ok := isu.mutation.YearMonth(); ok {
		_spec.SetField(invoicesequence.FieldYearMonth, field.TypeString, value)
	}
	if value, ok := isu.mutation.LastValue(); ok {
		_spec.SetField(invoicesequence.FieldLastValue, field.TypeInt64, value)
	}
	if value, ok := isu.mutation.AddedLastValue(); ok {
		_spec.AddField(invoicesequence.FieldLastValue, field.TypeInt64, value)
	}
	if value, ok := isu.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicesequence.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, isu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicesequence.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	isu.mutation.done = true
	return n, nil
}

// InvoiceSequenceUpdateOne is the builder for updating a single InvoiceSequence entity.
type InvoiceSequenceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvoiceSequenceMutation
}

// SetTenantID sets the "tenant_id" field.
func (isuo *InvoiceSequenceUpdateOne) SetTenantID(s string) *InvoiceSequenceUpdateOne {
	isuo.mutation.SetTenantID(s)
	return isuo
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (isuo *InvoiceSequenceUpdateOne) SetNillableTenantID(s *string) *InvoiceSequenceUpdateOne {
	if s != nil {
		isuo.SetTenantID(*s)
	}
	return isuo
}

// SetYearMonth sets the "year_month" field.
func (isuo *InvoiceSequenceUpdateOne) SetYearMonth(s string) *InvoiceSequenceUpdateOne {
	isuo.mutation.SetYearMonth(s)
	return isuo
}

// SetNillableYearMonth sets the "year_month" field if the given value is not nil.
func (isuo *InvoiceSequenceUpdateOne) SetNillableYearMonth(s *string) *InvoiceSequenceUpdateOne {
	if s != nil {
		isuo.SetYearMonth(*s)
	}
	return isuo
}

// SetLastValue sets the "last_value" field.
func (isuo *InvoiceSequenceUpdateOne) SetLastValue(i int64) *InvoiceSequenceUpdateOne {
	isuo.mutation.ResetLastValue()
	isuo.mutation.SetLastValue(i)
	return isuo
}

// SetNillableLastValue sets the "last_value" field if the given value is not nil.
func (isuo *InvoiceSequenceUpdateOne) SetNillableLastValue(i *int64) *InvoiceSequenceUpdateOne {
	if i != nil {
		isuo.SetLastValue(*i)
	}
	return isuo
}

// AddLastValue adds i to the "last_value" field.
func (isuo *InvoiceSequenceUpdateOne) AddLastValue(i int64) *InvoiceSequenceUpdateOne {
	isuo.mutation.AddLastValue(i)
	return isuo
}

// SetUpdatedAt sets the "updated_at" field.
func (isuo *InvoiceSequenceUpdateOne) SetUpdatedAt(t time.Time) *InvoiceSequenceUpdateOne {
	isuo.mutation.SetUpdatedAt(t)
	return isuo
}

// Mutation returns the InvoiceSequenceMutation object of the builder.
func (isuo *InvoiceSequenceUpdateOne) Mutation() *InvoiceSequenceMutation {
	return isuo.mutation
}

// Where appends a list predicates to the InvoiceSequenceUpdate builder.
func (isuo *InvoiceSequenceUpdateOne) Where(ps ...predicate.InvoiceSequence) *InvoiceSequenceUpdateOne {
	isuo.mutation.Where(ps...)
	return isuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (isuo *InvoiceSequenceUpdateOne) Select(field string, fields ...string) *InvoiceSequenceUpdateOne {
	isuo.fields = append([]string{field}, fields...)
	return isuo
}

// Save executes the query and returns the updated InvoiceSequence entity.
func (isuo *InvoiceSequenceUpdateOne) Save(ctx context.Context) (*InvoiceSequence, error) {
	isuo.defaults()
	return withHooks(ctx, isuo.sqlSave, isuo.mutation, isuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (isuo *InvoiceSequenceUpdateOne) SaveX(ctx context.Context) *InvoiceSequence {
	node, err := isuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (isuo *InvoiceSequenceUpdateOne) Exec(ctx context.Context) error {
	_, err := isuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isuo *InvoiceSequenceUpdateOne) ExecX(ctx context.Context) {
	if err := isuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (isuo *InvoiceSequenceUpdateOne) defaults() {
	if _, ok := isuo.mutation.UpdatedAt(); !ok {
		v := invoicesequence.UpdateDefaultUpdatedAt()
		isuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (isuo *InvoiceSequenceUpdateOne) check() error {
	if v, ok := isuo.mutation.TenantID(); ok {
		if err := invoicesequence.TenantIDValidator(v); err != nil {
			return &ValidationError{Name: "tenant_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceSequence.tenant_id": %w`, err)}
		}
	}
	if v, ok := isuo.mutation.YearMonth(); ok {
		if err := invoicesequence.YearMonthValidator(v); err != nil {
			return &ValidationError{Name: "year_month", err: fmt.Errorf(`ent: validator failed for field "InvoiceSequence.year_month": %w`, err)}
		}
	}
	return nil
}

func (isuo *InvoiceSequenceUpdateOne) sqlSave(ctx context.Context) (_node *InvoiceSequence, err error) {
	if err := isuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicesequence.Table, invoicesequence.Columns, sqlgraph.NewFieldSpec(invoicesequence.FieldID, field.TypeInt))
	id, ok := isuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InvoiceSequence.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := isuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicesequence.FieldID)
		for _, f := range fields {
			if !invoicesequence.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoicesequence.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := isuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := isuo.mutation.TenantID(); ok {
		_spec.SetField(invoicesequence.FieldTenantID, field.TypeString, value)
	}
	if value, ok := isuo.mutation.YearMonth(); ok {
		_spec.SetField(invoicesequence.FieldYearMonth, field.TypeString, value)
	}
	if value, ok := isuo.mutation.LastValue(); ok {
		_spec.SetField(invoicesequence.FieldLastValue, field.TypeInt64, value)
	}
	if value, ok := isuo.mutation.AddedLastValue(); ok {
		_spec.AddField(invoicesequence.FieldLastValue, field.TypeInt64, value)
	}
	if value, ok := isuo.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicesequence.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &InvoiceSequence{config: isuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, isuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicesequence.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	isuo.mutation.done = true
	return _node, nil
}