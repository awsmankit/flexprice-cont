// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flexprice/flexprice/ent/billingsequence"
)

// BillingSequenceCreate is the builder for creating a BillingSequence entity.
type BillingSequenceCreate struct {
	config
	mutation *BillingSequenceMutation
	hooks    []Hook
}

// SetTenantID sets the "tenant_id" field.
func (bsc *BillingSequenceCreate) SetTenantID(s string) *BillingSequenceCreate {
	bsc.mutation.SetTenantID(s)
	return bsc
}

// SetSubscriptionID sets the "subscription_id" field.
func (bsc *BillingSequenceCreate) SetSubscriptionID(s string) *BillingSequenceCreate {
	bsc.mutation.SetSubscriptionID(s)
	return bsc
}

// SetLastSequence sets the "last_sequence" field.
func (bsc *BillingSequenceCreate) SetLastSequence(i int) *BillingSequenceCreate {
	bsc.mutation.SetLastSequence(i)
	return bsc
}

// SetNillableLastSequence sets the "last_sequence" field if the given value is not nil.
func (bsc *BillingSequenceCreate) SetNillableLastSequence(i *int) *BillingSequenceCreate {
	if i != nil {
		bsc.SetLastSequence(*i)
	}
	return bsc
}

// SetCreatedAt sets the "created_at" field.
func (bsc *BillingSequenceCreate) SetCreatedAt(t time.Time) *BillingSequenceCreate {
	bsc.mutation.SetCreatedAt(t)
	return bsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bsc *BillingSequenceCreate) SetNillableCreatedAt(t *time.Time) *BillingSequenceCreate {
	if t != nil {
		bsc.SetCreatedAt(*t)
	}
	return bsc
}

// SetUpdatedAt sets the "updated_at" field.
func (bsc *BillingSequenceCreate) SetUpdatedAt(t time.Time) *BillingSequenceCreate {
	bsc.mutation.SetUpdatedAt(t)
	return bsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bsc *BillingSequenceCreate) SetNillableUpdatedAt(t *time.Time) *BillingSequenceCreate {
	if t != nil {
		bsc.SetUpdatedAt(*t)
	}
	return bsc
}

// Mutation returns the BillingSequenceMutation object of the builder.
func (bsc *BillingSequenceCreate) Mutation() *BillingSequenceMutation {
	return bsc.mutation
}

// Save creates the BillingSequence in the database.
func (bsc *BillingSequenceCreate) Save(ctx context.Context) (*BillingSequence, error) {
	bsc.defaults()
	return withHooks(ctx, bsc.sqlSave, bsc.mutation, bsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bsc *BillingSequenceCreate) SaveX(ctx context.Context) *BillingSequence {
	v, err := bsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bsc *BillingSequenceCreate) Exec(ctx context.Context) error {
	_, err := bsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsc *BillingSequenceCreate) ExecX(ctx context.Context) {
	if err := bsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bsc *BillingSequenceCreate) defaults() {
	if _, ok := bsc.mutation.LastSequence(); !ok {
		v := billingsequence.DefaultLastSequence
		bsc.mutation.SetLastSequence(v)
	}
	if _, ok := bsc.mutation.CreatedAt(); !ok {
		v := billingsequence.DefaultCreatedAt()
		bsc.mutation.SetCreatedAt(v)
	}
	if _, ok := bsc.mutation.UpdatedAt(); !ok {
		v := billingsequence.DefaultUpdatedAt()
		bsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bsc *BillingSequenceCreate) check() error {
	if _, ok := bsc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "BillingSequence.tenant_id"`)}
	}
	if v, ok := bsc.mutation.TenantID(); ok {
		if err := billingsequence.TenantIDValidator(v); err != nil {
			return &ValidationError{Name: "tenant_id", err: fmt.Errorf(`ent: validator failed for field "BillingSequence.tenant_id": %w`, err)}
		}
	}
	if _, ok := bsc.mutation.SubscriptionID(); !ok {
		return &ValidationError{Name: "subscription_id", err: errors.New(`ent: missing required field "BillingSequence.subscription_id"`)}
	}
	if v, ok := bsc.mutation.SubscriptionID(); ok {
		if err := billingsequence.SubscriptionIDValidator(v); err != nil {
			return &ValidationError{Name: "subscription_id", err: fmt.Errorf(`ent: validator failed for field "BillingSequence.subscription_id": %w`, err)}
		}
	}
	if _, ok := bsc.mutation.LastSequence(); !ok {
		return &ValidationError{Name: "last_sequence", err: errors.New(`ent: missing required field "BillingSequence.last_sequence"`)}
	}
	if _, ok := bsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BillingSequence.created_at"`)}
	}
	if _, ok := bsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BillingSequence.updated_at"`)}
	}
	return nil
}

func (bsc *BillingSequenceCreate) sqlSave(ctx context.Context) (*BillingSequence, error) {
	if err := bsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bsc.mutation.id = &_node.ID
	bsc.mutation.done = true
	return _node, nil
}

func (bsc *BillingSequenceCreate) createSpec() (*BillingSequence, *sqlgraph.CreateSpec) {
	var (
		_node = &BillingSequence{config: bsc.config}
		_spec = sqlgraph.NewCreateSpec(billingsequence.Table, sqlgraph.NewFieldSpec(billingsequence.FieldID, field.TypeInt))
	)
	if value, ok := bsc.mutation.TenantID(); ok {
		_spec.SetField(billingsequence.FieldTenantID, field.TypeString, value)
		_node.TenantID = value
	}
	if value, ok := bsc.mutation.SubscriptionID(); ok {
		_spec.SetField(billingsequence.FieldSubscriptionID, field.TypeString, value)
		_node.SubscriptionID = value
	}
	if value, ok := bsc.mutation.LastSequence(); ok {
		_spec.SetField(billingsequence.FieldLastSequence, field.TypeInt, value)
		_node.LastSequence = value
	}
	if value, ok := bsc.mutation.CreatedAt(); ok {
		_spec.SetField(billingsequence.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bsc.mutation.UpdatedAt(); ok {
		_spec.SetField(billingsequence.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// BillingSequenceCreateBulk is the builder for creating many BillingSequence entities in bulk.
type BillingSequenceCreateBulk struct {
	config
	err      error
	builders []*BillingSequenceCreate
}

// Save creates the BillingSequence entities in the database.
func (bscb *BillingSequenceCreateBulk) Save(ctx context.Context) ([]*BillingSequence, error) {
	if bscb.err != nil {
		return nil, bscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bscb.builders))
	nodes := make([]*BillingSequence, len(bscb.builders))
	mutators := make([]Mutator, len(bscb.builders))
	for i := range bscb.builders {
		func(i int, root context.Context) {
			builder := bscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillingSequenceMutation)
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
					_, err = mutators[i+1].Mutate(root, bscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bscb *BillingSequenceCreateBulk) SaveX(ctx context.Context) []*BillingSequence {
	v, err := bscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bscb *BillingSequenceCreateBulk) Exec(ctx context.Context) error {
	_, err := bscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bscb *BillingSequenceCreateBulk) ExecX(ctx context.Context) {
	if err := bscb.Exec(ctx); err != nil {
		panic(err)
	}
}
