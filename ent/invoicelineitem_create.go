// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flexprice/flexprice/ent/invoice"
	"github.com/flexprice/flexprice/ent/invoicelineitem"
	"github.com/shopspring/decimal"
)

// InvoiceLineItemCreate is the builder for creating a InvoiceLineItem entity.
type InvoiceLineItemCreate struct {
	config
	mutation *InvoiceLineItemMutation
	hooks    []Hook
}

// SetTenantID sets the "tenant_id" field.
func (ilic *InvoiceLineItemCreate) SetTenantID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetTenantID(s)
	return ilic
}

// SetStatus sets the "status" field.
func (ilic *InvoiceLineItemCreate) SetStatus(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetStatus(s)
	return ilic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableStatus(s *string) *InvoiceLineItemCreate {
	if s != nil {
		ilic.SetStatus(*s)
	}
	return ilic
}

// SetCreatedAt sets the "created_at" field.
func (ilic *InvoiceLineItemCreate) SetCreatedAt(t time.Time) *InvoiceLineItemCreate {
	ilic.mutation.SetCreatedAt(t)
	return ilic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableCreatedAt(t *time.Time) *InvoiceLineItemCreate {
	if t != nil {
		ilic.SetCreatedAt(*t)
	}
	return ilic
}

// SetUpdatedAt sets the "updated_at" field.
func (ilic *InvoiceLineItemCreate) SetUpdatedAt(t time.Time) *InvoiceLineItemCreate {
	ilic.mutation.SetUpdatedAt(t)
	return ilic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableUpdatedAt(t *time.Time) *InvoiceLineItemCreate {
	if t != nil {
		ilic.SetUpdatedAt(*t)
	}
	return ilic
}

// SetCreatedBy sets the "created_by" field.
func (ilic *InvoiceLineItemCreate) SetCreatedBy(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetCreatedBy(s)
	return ilic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableCreatedBy(s *string) *InvoiceLineItemCreate {
	if s != nil {
		ilic.SetCreatedBy(*s)
	}
	return ilic
}

// SetUpdatedBy sets the "updated_by" field.
func (ilic *InvoiceLineItemCreate) SetUpdatedBy(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetUpdatedBy(s)
	return ilic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableUpdatedBy(s *string) *InvoiceLineItemCreate {
	if s != nil {
		ilic.SetUpdatedBy(*s)
	}
	return ilic
}

// SetInvoiceID sets the "invoice_id" field.
func (ilic *InvoiceLineItemCreate) SetInvoiceID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetInvoiceID(s)
	return ilic
}

// SetCustomerID sets the "customer_id" field.
func (ilic *InvoiceLineItemCreate) SetCustomerID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetCustomerID(s)
	return ilic
}

// SetSubscriptionID sets the "subscription_id" field.
func (ilic *InvoiceLineItemCreate) SetSubscriptionID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetSubscriptionID(s)
	return ilic
}

// SetNillableSubscriptionID sets the "subscription_id" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableSubscriptionID(s *string) *InvoiceLineItemCreate {
	if s != nil {
		ilic.SetSubscriptionID(*s)
	}
	return ilic
}

// SetPriceID sets the "price_id" field.
func (ilic *InvoiceLineItemCreate) SetPriceID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetPriceID(s)
	return ilic
}

// SetMeterID sets the "meter_id" field.
func (ilic *InvoiceLineItemCreate) SetMeterID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetMeterID(s)
	return ilic
}

// SetNillableMeterID sets the "meter_id" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableMeterID(s *string) *InvoiceLineItemCreate {
	if s != nil {
		ilic.SetMeterID(*s)
	}
	return ilic
}

// SetAmount sets the "amount" field.
func (ilic *InvoiceLineItemCreate) SetAmount(d decimal.Decimal) *InvoiceLineItemCreate {
	ilic.mutation.SetAmount(d)
	return ilic
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableAmount(d *decimal.Decimal) *InvoiceLineItemCreate {
	if d != nil {
		ilic.SetAmount(*d)
	}
	return ilic
}

// SetQuantity sets the "quantity" field.
func (ilic *InvoiceLineItemCreate) SetQuantity(d decimal.Decimal) *InvoiceLineItemCreate {
	ilic.mutation.SetQuantity(d)
	return ilic
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillableQuantity(d *decimal.Decimal) *InvoiceLineItemCreate {
	if d != nil {
		ilic.SetQuantity(*d)
	}
	return ilic
}

// SetCurrency sets the "currency" field.
func (ilic *InvoiceLineItemCreate) SetCurrency(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetCurrency(s)
	return ilic
}

// SetPeriodStart sets the "period_start" field.
func (ilic *InvoiceLineItemCreate) SetPeriodStart(t time.Time) *InvoiceLineItemCreate {
	ilic.mutation.SetPeriodStart(t)
	return ilic
}

// SetNillablePeriodStart sets the "period_start" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillablePeriodStart(t *time.Time) *InvoiceLineItemCreate {
	if t != nil {
		ilic.SetPeriodStart(*t)
	}
	return ilic
}

// SetPeriodEnd sets the "period_end" field.
func (ilic *InvoiceLineItemCreate) SetPeriodEnd(t time.Time) *InvoiceLineItemCreate {
	ilic.mutation.SetPeriodEnd(t)
	return ilic
}

// SetNillablePeriodEnd sets the "period_end" field if the given value is not nil.
func (ilic *InvoiceLineItemCreate) SetNillablePeriodEnd(t *time.Time) *InvoiceLineItemCreate {
	if t != nil {
		ilic.SetPeriodEnd(*t)
	}
	return ilic
}

// SetMetadata sets the "metadata" field.
func (ilic *InvoiceLineItemCreate) SetMetadata(m map[string]string) *InvoiceLineItemCreate {
	ilic.mutation.SetMetadata(m)
	return ilic
}

// SetID sets the "id" field.
func (ilic *InvoiceLineItemCreate) SetID(s string) *InvoiceLineItemCreate {
	ilic.mutation.SetID(s)
	return ilic
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (ilic *InvoiceLineItemCreate) SetInvoice(i *Invoice) *InvoiceLineItemCreate {
	return ilic.SetInvoiceID(i.ID)
}

// Mutation returns the InvoiceLineItemMutation object of the builder.
func (ilic *InvoiceLineItemCreate) Mutation() *InvoiceLineItemMutation {
	return ilic.mutation
}

// Save creates the InvoiceLineItem in the database.
func (ilic *InvoiceLineItemCreate) Save(ctx context.Context) (*InvoiceLineItem, error) {
	ilic.defaults()
	return withHooks(ctx, ilic.sqlSave, ilic.mutation, ilic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ilic *InvoiceLineItemCreate) SaveX(ctx context.Context) *InvoiceLineItem {
	v, err := ilic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilic *InvoiceLineItemCreate) Exec(ctx context.Context) error {
	_, err := ilic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilic *InvoiceLineItemCreate) ExecX(ctx context.Context) {
	if err := ilic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ilic *InvoiceLineItemCreate) defaults() {
	if _, ok := ilic.mutation.Status(); !ok {
		v := invoicelineitem.DefaultStatus
		ilic.mutation.SetStatus(v)
	}
	if _, ok := ilic.mutation.CreatedAt(); !ok {
		v := invoicelineitem.DefaultCreatedAt()
		ilic.mutation.SetCreatedAt(v)
	}
	if _, ok := ilic.mutation.UpdatedAt(); !ok {
		v := invoicelineitem.DefaultUpdatedAt()
		ilic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ilic.mutation.Amount(); !ok {
		v := invoicelineitem.DefaultAmount
		ilic.mutation.SetAmount(v)
	}
	if _, ok := ilic.mutation.Quantity(); !ok {
		v := invoicelineitem.DefaultQuantity
		ilic.mutation.SetQuantity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ilic *InvoiceLineItemCreate) check() error {
	if _, ok := ilic.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "InvoiceLineItem.tenant_id"`)}
	}
	if v, ok := ilic.mutation.TenantID(); ok {
		if err := invoicelineitem.TenantIDValidator(v); err != nil {
			return &ValidationError{Name: "tenant_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceLineItem.tenant_id": %w`, err)}
		}
	}
	if _, ok := ilic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "InvoiceLineItem.status"`)}
	}
	if _, ok := ilic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "InvoiceLineItem.created_at"`)}
	}
	if _, ok := ilic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "InvoiceLineItem.updated_at"`)}
	}
	if _, ok := ilic.mutation.InvoiceID(); !ok {
		return &ValidationError{Name: "invoice_id", err: errors.New(`ent: missing required field "InvoiceLineItem.invoice_id"`)}
	}
	if v, ok := ilic.mutation.InvoiceID(); ok {
		if err := invoicelineitem.InvoiceIDValidator(v); err != nil {
			return &ValidationError{Name: "invoice_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceLineItem.invoice_id": %w`, err)}
		}
	}
	if _, ok := ilic.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`ent: missing required field "InvoiceLineItem.customer_id"`)}
	}
	if v, ok := ilic.mutation.CustomerID(); ok {
		if err := invoicelineitem.CustomerIDValidator(v); err != nil {
			return &ValidationError{Name: "customer_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceLineItem.customer_id": %w`, err)}
		}
	}
	if _, ok := ilic.mutation.PriceID(); !ok {
		return &ValidationError{Name: "price_id", err: errors.New(`ent: missing required field "InvoiceLineItem.price_id"`)}
	}
	if v, ok := ilic.mutation.PriceID(); ok {
		if err := invoicelineitem.PriceIDValidator(v); err != nil {
			return &ValidationError{Name: "price_id", err: fmt.Errorf(`ent: validator failed for field "InvoiceLineItem.price_id": %w`, err)}
		}
	}
	if _, ok := ilic.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "InvoiceLineItem.amount"`)}
	}
	if _, ok := ilic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "InvoiceLineItem.quantity"`)}
	}
	if _, ok := ilic.mutation.Currency(); !ok {
		return &ValidationError{Name: "currency", err: errors.New(`ent: missing required field "InvoiceLineItem.currency"`)}
	}
	if v, ok := ilic.mutation.Currency(); ok {
		if err := invoicelineitem.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "InvoiceLineItem.currency": %w`, err)}
		}
	}
	if len(ilic.mutation.InvoiceIDs()) == 0 {
		return &ValidationError{Name: "invoice", err: errors.New(`ent: missing required edge "InvoiceLineItem.invoice"`)}
	}
	return nil
}

func (ilic *InvoiceLineItemCreate) sqlSave(ctx context.Context) (*InvoiceLineItem, error) {
	if err := ilic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ilic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ilic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected InvoiceLineItem.ID type: %T", _spec.ID.Value)
		}
	}
	ilic.mutation.id = &_node.ID
	ilic.mutation.done = true
	return _node, nil
}

func (ilic *InvoiceLineItemCreate) createSpec() (*InvoiceLineItem, *sqlgraph.CreateSpec) {
	var (
		_node = &InvoiceLineItem{config: ilic.config}
		_spec = sqlgraph.NewCreateSpec(invoicelineitem.Table, sqlgraph.NewFieldSpec(invoicelineitem.FieldID, field.TypeString))
	)
	if id, ok := ilic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ilic.mutation.TenantID(); ok {
		_spec.SetField(invoicelineitem.FieldTenantID, field.TypeString, value)
		_node.TenantID = value
	}
	if value, ok := ilic.mutation.Status(); ok {
		_spec.SetField(invoicelineitem.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := ilic.mutation.CreatedAt(); ok {
		_spec.SetField(invoicelineitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ilic.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicelineitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ilic.mutation.CreatedBy(); ok {
		_spec.SetField(invoicelineitem.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ilic.mutation.UpdatedBy(); ok {
		_spec.SetField(invoicelineitem.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ilic.mutation.CustomerID(); ok {
		_spec.SetField(invoicelineitem.FieldCustomerID, field.TypeString, value)
		_node.CustomerID = value
	}
	if value, ok := ilic.mutation.SubscriptionID(); ok {
		_spec.SetField(invoicelineitem.FieldSubscriptionID, field.TypeString, value)
		_node.SubscriptionID = &value
	}
	if value, ok := ilic.mutation.PriceID(); ok {
		_spec.SetField(invoicelineitem.FieldPriceID, field.TypeString, value)
		_node.PriceID = value
	}
	if value, ok := ilic.mutation.MeterID(); ok {
		_spec.SetField(invoicelineitem.FieldMeterID, field.TypeString, value)
		_node.MeterID = &value
	}
	if value, ok := ilic.mutation.Amount(); ok {
		_spec.SetField(invoicelineitem.FieldAmount, field.TypeOther, value)
		_node.Amount = value
	}
	if value, ok := ilic.mutation.Quantity(); ok {
		_spec.SetField(invoicelineitem.FieldQuantity, field.TypeOther, value)
		_node.Quantity = value
	}
	if value, ok := ilic.mutation.Currency(); ok {
		_spec.SetField(invoicelineitem.FieldCurrency, field.TypeString, value)
		_node.Currency = value
	}
	if value, ok := ilic.mutation.PeriodStart(); ok {
		_spec.SetField(invoicelineitem.FieldPeriodStart, field.TypeTime, value)
		_node.PeriodStart = &value
	}
	if value, ok := ilic.mutation.PeriodEnd(); ok {
		_spec.SetField(invoicelineitem.FieldPeriodEnd, field.TypeTime, value)
		_node.PeriodEnd = &value
	}
	if value, ok := ilic.mutation.Metadata(); ok {
		_spec.SetField(invoicelineitem.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := ilic.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invoicelineitem.InvoiceTable,
			Columns: []string{invoicelineitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InvoiceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvoiceLineItemCreateBulk is the builder for creating many InvoiceLineItem entities in bulk.
type InvoiceLineItemCreateBulk struct {
	config
	err      error
	builders []*InvoiceLineItemCreate
}

// Save creates the InvoiceLineItem entities in the database.
func (ilicb *InvoiceLineItemCreateBulk) Save(ctx context.Context) ([]*InvoiceLineItem, error) {
	if ilicb.err != nil {
		return nil, ilicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ilicb.builders))
	nodes := make([]*InvoiceLineItem, len(ilicb.builders))
	mutators := make([]Mutator, len(ilicb.builders))
	for i := range ilicb.builders {
		func(i int, root context.Context) {
			builder := ilicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceLineItemMutation)
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
					_, err = mutators[i+1].Mutate(root, ilicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ilicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ilicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ilicb *InvoiceLineItemCreateBulk) SaveX(ctx context.Context) []*InvoiceLineItem {
	v, err := ilicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilicb *InvoiceLineItemCreateBulk) Exec(ctx context.Context) error {
	_, err := ilicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilicb *InvoiceLineItemCreateBulk) ExecX(ctx context.Context) {
	if err := ilicb.Exec(ctx); err != nil {
		panic(err)
	}
}