// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flexprice/flexprice/ent/invoicesequence"
	"github.com/flexprice/flexprice/ent/predicate"
)

// InvoiceSequenceDelete is the builder for deleting a InvoiceSequence entity.
type InvoiceSequenceDelete struct {
	config
	hooks    []Hook
	mutation *InvoiceSequenceMutation
}

// Where appends a list predicates to the InvoiceSequenceDelete builder.
func (isd *InvoiceSequenceDelete) Where(ps ...predicate.InvoiceSequence) *InvoiceSequenceDelete {
	isd.mutation.Where(ps...)
	return isd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (isd *InvoiceSequenceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, isd.sqlExec, isd.mutation, isd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (isd *InvoiceSequenceDelete) ExecX(ctx context.Context) int {
	n, err := isd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (isd *InvoiceSequenceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(invoicesequence.Table, sqlgraph.NewFieldSpec(invoicesequence.FieldID, field.TypeInt))
	if ps := isd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, isd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	isd.mutation.done = true
	return affected, err
}

// InvoiceSequenceDeleteOne is the builder for deleting a single InvoiceSequence entity.
type InvoiceSequenceDeleteOne struct {
	isd *InvoiceSequenceDelete
}

// Where appends a list predicates to the InvoiceSequenceDelete builder.
func (isdo *InvoiceSequenceDeleteOne) Where(ps ...predicate.InvoiceSequence) *InvoiceSequenceDeleteOne {
	isdo.isd.mutation.Where(ps...)
	return isdo
}

// Exec executes the deletion query.
func (isdo *InvoiceSequenceDeleteOne) Exec(ctx context.Context) error {
	n, err := isdo.isd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invoicesequence.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (isdo *InvoiceSequenceDeleteOne) ExecX(ctx context.Context) {
	if err := isdo.Exec(ctx); err != nil {
		panic(err)
	}
}