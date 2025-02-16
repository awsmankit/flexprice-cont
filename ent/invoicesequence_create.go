// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flexprice/flexprice/ent/invoicesequence"
)

// InvoiceSequenceCreate is the builder for creating a InvoiceSequence entity.
type InvoiceSequenceCreate struct {
	config
	mutation *InvoiceSequenceMutation
	hooks    []Hook
}

// SetTenantID sets the "tenant_id" field.
func (isc *InvoiceSequenceCreate) SetTenantID(s string) *InvoiceSequenceCreate {
	isc.mutation.SetTenantID(s)
	return isc
}

// SetYearMonth sets the "year_month" field.
func (isc *InvoiceSequenceCreate) SetYearMonth(s string) *InvoiceSequenceCreate {
	isc.mutation.SetYearMonth(s)
	return isc
}

// SetLastValue sets the "last_value" field.
func (isc *InvoiceSequenceCreate) SetLastValue(i int64) *InvoiceSequenceCreate {
	isc.mutation.SetLastValue(i)
	return isc
}

// SetNillableLastValue sets the "last_value" field if the given value is not nil.
func (isc *InvoiceSequenceCreate) SetNillableLastValue(i *int64) *InvoiceSequenceCreate {
	if i != nil {
		isc.SetLastValue(*i)
	}
	return isc
}

// SetCreatedAt sets the "created_at" field.
func (isc *InvoiceSequenceCreate) SetCreatedAt(t time.Time) *InvoiceSequenceCreate {
	isc.mutation.SetCreatedAt(t)
	return isc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (isc *InvoiceSequenceCreate) SetNillableCreatedAt(t *time.Time) *InvoiceSequenceCreate {
	if t != nil {
		isc.SetCreatedAt(*t)
	}
	return isc
}

// SetUpdatedAt sets the "updated_at" field.
func (isc *InvoiceSequenceCreate) SetUpdatedAt(t time.Time) *InvoiceSequenceCreate {
	isc.mutation.SetUpdatedAt(t)
	return isc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (isc *InvoiceSequenceCreate) SetNillableUpdatedAt(t *time.Time) *InvoiceSequenceCreate {
	if t != nil {
		isc.SetUpdatedAt(*t)
	}
	return isc
}

// Mutation returns the InvoiceSequenceMutation object of the builder.
func (isc *InvoiceSequenceCreate) Mutation() *InvoiceSequenceMutation {
	return isc.mutation
}

// Save creates the InvoiceSequence in the database.
func (isc *InvoiceSequenceCreate) Save(ctx context.Context) (*InvoiceSequence, error) {
	isc.defaults()
	return withHooks(ctx, isc.sqlSave, isc.mutation, isc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (isc *InvoiceSequenceCreate) SaveX(ctx context.Context) *InvoiceSequence {
	v, err := isc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (isc *InvoiceSequenceCreate) Exec(ctx context.Context) error {
	_, err := isc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isc *InvoiceSequenceCreate) ExecX(ctx context.Context) {
	if err := isc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (isc *InvoiceSequenceCreate) defaults() {
	if _, ok := isc.mutation.LastValue(); !ok {
		v := invoicesequence.DefaultLastValue
		isc.mutation.SetLastValue(v)
	}
	if _, ok := isc.mutation.CreatedAt(); !ok {
		v := invoicesequence.DefaultCreatedAt()
		isc.mutation.SetCreatedAt(v)
	}
	if _, ok := isc.mutation.UpdatedAt(); !ok {
		v := invoicesequence.DefaultUpdatedAt()
		isc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (isc *InvoiceSequenceCreate) check() error {
	if _, ok := isc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "InvoiceSequence.tenant_id"`)}
	}
	if v, ok := isc.mutation.TenantID(); ok {
		if err := invoicesequence.TenantIDValidator(v); err != nil {
			return &ValidationError{Name: "tenant_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceSequence.tenant_id": %w`, err)}
		}
	}
	if _, ok := isc.mutation.YearMonth(); !ok {
		return &ValidationError{Name: "year_month", err: errors.New(`ent: missing required field "InvoiceSequence.year_month"`)}
	}
	if v, ok := isc.mutation.YearMonth(); ok {
		if err := invoicesequence.YearMonthValidator(v); err != nil {
			return &ValidationError{Name: "year_month", err: fmt.Errorf(`ent: validator failed for field "InvoiceSequence.year_month": %w`, err)}
		}
	}
	if _, ok := isc.mutation.LastValue(); !ok {
		return &ValidationError{Name: "last_value", err: errors.New(`ent: missing required field "InvoiceSequence.last_value"`)}
	}
	if _, ok := isc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "InvoiceSequence.created_at"`)}
	}
	if _, ok := isc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "InvoiceSequence.updated_at"`)}
	}
	return nil
}

func (isc *InvoiceSequenceCreate) sqlSave(ctx context.Context) (*InvoiceSequence, error) {
	if err := isc.check(); err != nil {
		return nil, err
	}
	_node, _spec := isc.createSpec()
	if err := sqlgraph.CreateNode(ctx, isc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	isc.mutation.id = &_node.ID
	isc.mutation.done = true
	return _node, nil
}

func (isc *InvoiceSequenceCreate) createSpec() (*InvoiceSequence, *sqlgraph.CreateSpec) {
	var (
		_node = &InvoiceSequence{config: isc.config}
		_spec = sqlgraph.NewCreateSpec(invoicesequence.Table, sqlgraph.NewFieldSpec(invoicesequence.FieldID, field.TypeInt))
	)
	if value, ok := isc.mutation.TenantID(); ok {
		_spec.SetField(invoicesequence.FieldTenantID, field.TypeString, value)
		_node.TenantID = value
	}
	if value, ok := isc.mutation.YearMonth(); ok {
		_spec.SetField(invoicesequence.FieldYearMonth, field.TypeString, value)
		_node.YearMonth = value
	}
	if value, ok := isc.mutation.LastValue(); ok {
		_spec.SetField(invoicesequence.FieldLastValue, field.TypeInt64, value)
		_node.LastValue = value
	}
	if value, ok := isc.mutation.CreatedAt(); ok {
		_spec.SetField(invoicesequence.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := isc.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicesequence.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// InvoiceSequenceCreateBulk is the builder for creating many InvoiceSequence entities in bulk.
type InvoiceSequenceCreateBulk struct {
	config
	err      error
	builders []*InvoiceSequenceCreate
}

// Save creates the InvoiceSequence entities in the database.
func (iscb *InvoiceSequenceCreateBulk) Save(ctx context.Context) ([]*InvoiceSequence, error) {
	if iscb.err != nil {
		return nil, iscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(iscb.builders))
	nodes := make([]*InvoiceSequence, len(iscb.builders))
	mutators := make([]Mutator, len(iscb.builders))
	for i := range iscb.builders {
		func(i int, root context.Context) {
			builder := iscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceSequenceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, iscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, iscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iscb *InvoiceSequenceCreateBulk) SaveX(ctx context.Context) []*InvoiceSequence {
	v, err := iscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iscb *InvoiceSequenceCreateBulk) Exec(ctx context.Context) error {
	_, err := iscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iscb *InvoiceSequenceCreateBulk) ExecX(ctx context.Context) {
	if err := iscb.Exec(ctx); err != nil {
		panic(err)
	}
}
