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
	"github.com/flexprice/flexprice/ent/invoice"
	"github.com/flexprice/flexprice/ent/predicate"
	"github.com/shopspring/decimal"
)

// InvoiceUpdate is the builder for updating Invoice entities.
type InvoiceUpdate struct {
	config
	hooks    []Hook
	mutation *InvoiceMutation
}

// Where appends a list predicates to the InvoiceUpdate builder.
func (iu *InvoiceUpdate) Where(ps ...predicate.Invoice) *InvoiceUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetStatus sets the "status" field.
func (iu *InvoiceUpdate) SetStatus(s string) *InvoiceUpdate {
	iu.mutation.SetStatus(s)
	return iu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableStatus(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetStatus(*s)
	}
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *InvoiceUpdate) SetUpdatedAt(t time.Time) *InvoiceUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetUpdatedBy sets the "updated_by" field.
func (iu *InvoiceUpdate) SetUpdatedBy(s string) *InvoiceUpdate {
	iu.mutation.SetUpdatedBy(s)
	return iu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableUpdatedBy(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetUpdatedBy(*s)
	}
	return iu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iu *InvoiceUpdate) ClearUpdatedBy() *InvoiceUpdate {
	iu.mutation.ClearUpdatedBy()
	return iu
}

// SetSubscriptionID sets the "subscription_id" field.
func (iu *InvoiceUpdate) SetSubscriptionID(s string) *InvoiceUpdate {
	iu.mutation.SetSubscriptionID(s)
	return iu
}

// SetNillableSubscriptionID sets the "subscription_id" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableSubscriptionID(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetSubscriptionID(*s)
	}
	return iu
}

// ClearSubscriptionID clears the value of the "subscription_id" field.
func (iu *InvoiceUpdate) ClearSubscriptionID() *InvoiceUpdate {
	iu.mutation.ClearSubscriptionID()
	return iu
}

// SetInvoiceStatus sets the "invoice_status" field.
func (iu *InvoiceUpdate) SetInvoiceStatus(s string) *InvoiceUpdate {
	iu.mutation.SetInvoiceStatus(s)
	return iu
}

// SetNillableInvoiceStatus sets the "invoice_status" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableInvoiceStatus(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetInvoiceStatus(*s)
	}
	return iu
}

// SetPaymentStatus sets the "payment_status" field.
func (iu *InvoiceUpdate) SetPaymentStatus(s string) *InvoiceUpdate {
	iu.mutation.SetPaymentStatus(s)
	return iu
}

// SetNillablePaymentStatus sets the "payment_status" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillablePaymentStatus(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetPaymentStatus(*s)
	}
	return iu
}

// SetAmountDue sets the "amount_due" field.
func (iu *InvoiceUpdate) SetAmountDue(d decimal.Decimal) *InvoiceUpdate {
	iu.mutation.SetAmountDue(d)
	return iu
}

// SetNillableAmountDue sets the "amount_due" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableAmountDue(d *decimal.Decimal) *InvoiceUpdate {
	if d != nil {
		iu.SetAmountDue(*d)
	}
	return iu
}

// SetAmountPaid sets the "amount_paid" field.
func (iu *InvoiceUpdate) SetAmountPaid(d decimal.Decimal) *InvoiceUpdate {
	iu.mutation.SetAmountPaid(d)
	return iu
}

// SetNillableAmountPaid sets the "amount_paid" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableAmountPaid(d *decimal.Decimal) *InvoiceUpdate {
	if d != nil {
		iu.SetAmountPaid(*d)
	}
	return iu
}

// SetAmountRemaining sets the "amount_remaining" field.
func (iu *InvoiceUpdate) SetAmountRemaining(d decimal.Decimal) *InvoiceUpdate {
	iu.mutation.SetAmountRemaining(d)
	return iu
}

// SetNillableAmountRemaining sets the "amount_remaining" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableAmountRemaining(d *decimal.Decimal) *InvoiceUpdate {
	if d != nil {
		iu.SetAmountRemaining(*d)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *InvoiceUpdate) SetDescription(s string) *InvoiceUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableDescription(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// ClearDescription clears the value of the "description" field.
func (iu *InvoiceUpdate) ClearDescription() *InvoiceUpdate {
	iu.mutation.ClearDescription()
	return iu
}

// SetDueDate sets the "due_date" field.
func (iu *InvoiceUpdate) SetDueDate(t time.Time) *InvoiceUpdate {
	iu.mutation.SetDueDate(t)
	return iu
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableDueDate(t *time.Time) *InvoiceUpdate {
	if t != nil {
		iu.SetDueDate(*t)
	}
	return iu
}

// ClearDueDate clears the value of the "due_date" field.
func (iu *InvoiceUpdate) ClearDueDate() *InvoiceUpdate {
	iu.mutation.ClearDueDate()
	return iu
}

// SetPaidAt sets the "paid_at" field.
func (iu *InvoiceUpdate) SetPaidAt(t time.Time) *InvoiceUpdate {
	iu.mutation.SetPaidAt(t)
	return iu
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillablePaidAt(t *time.Time) *InvoiceUpdate {
	if t != nil {
		iu.SetPaidAt(*t)
	}
	return iu
}

// ClearPaidAt clears the value of the "paid_at" field.
func (iu *InvoiceUpdate) ClearPaidAt() *InvoiceUpdate {
	iu.mutation.ClearPaidAt()
	return iu
}

// SetVoidedAt sets the "voided_at" field.
func (iu *InvoiceUpdate) SetVoidedAt(t time.Time) *InvoiceUpdate {
	iu.mutation.SetVoidedAt(t)
	return iu
}

// SetNillableVoidedAt sets the "voided_at" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableVoidedAt(t *time.Time) *InvoiceUpdate {
	if t != nil {
		iu.SetVoidedAt(*t)
	}
	return iu
}

// ClearVoidedAt clears the value of the "voided_at" field.
func (iu *InvoiceUpdate) ClearVoidedAt() *InvoiceUpdate {
	iu.mutation.ClearVoidedAt()
	return iu
}

// SetFinalizedAt sets the "finalized_at" field.
func (iu *InvoiceUpdate) SetFinalizedAt(t time.Time) *InvoiceUpdate {
	iu.mutation.SetFinalizedAt(t)
	return iu
}

// SetNillableFinalizedAt sets the "finalized_at" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableFinalizedAt(t *time.Time) *InvoiceUpdate {
	if t != nil {
		iu.SetFinalizedAt(*t)
	}
	return iu
}

// ClearFinalizedAt clears the value of the "finalized_at" field.
func (iu *InvoiceUpdate) ClearFinalizedAt() *InvoiceUpdate {
	iu.mutation.ClearFinalizedAt()
	return iu
}

// SetInvoicePdfURL sets the "invoice_pdf_url" field.
func (iu *InvoiceUpdate) SetInvoicePdfURL(s string) *InvoiceUpdate {
	iu.mutation.SetInvoicePdfURL(s)
	return iu
}

// SetNillableInvoicePdfURL sets the "invoice_pdf_url" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableInvoicePdfURL(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetInvoicePdfURL(*s)
	}
	return iu
}

// ClearInvoicePdfURL clears the value of the "invoice_pdf_url" field.
func (iu *InvoiceUpdate) ClearInvoicePdfURL() *InvoiceUpdate {
	iu.mutation.ClearInvoicePdfURL()
	return iu
}

// SetBillingReason sets the "billing_reason" field.
func (iu *InvoiceUpdate) SetBillingReason(s string) *InvoiceUpdate {
	iu.mutation.SetBillingReason(s)
	return iu
}

// SetNillableBillingReason sets the "billing_reason" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableBillingReason(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetBillingReason(*s)
	}
	return iu
}

// ClearBillingReason clears the value of the "billing_reason" field.
func (iu *InvoiceUpdate) ClearBillingReason() *InvoiceUpdate {
	iu.mutation.ClearBillingReason()
	return iu
}

// SetMetadata sets the "metadata" field.
func (iu *InvoiceUpdate) SetMetadata(m map[string]interface{}) *InvoiceUpdate {
	iu.mutation.SetMetadata(m)
	return iu
}

// ClearMetadata clears the value of the "metadata" field.
func (iu *InvoiceUpdate) ClearMetadata() *InvoiceUpdate {
	iu.mutation.ClearMetadata()
	return iu
}

// SetVersion sets the "version" field.
func (iu *InvoiceUpdate) SetVersion(i int) *InvoiceUpdate {
	iu.mutation.ResetVersion()
	iu.mutation.SetVersion(i)
	return iu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableVersion(i *int) *InvoiceUpdate {
	if i != nil {
		iu.SetVersion(*i)
	}
	return iu
}

// AddVersion adds i to the "version" field.
func (iu *InvoiceUpdate) AddVersion(i int) *InvoiceUpdate {
	iu.mutation.AddVersion(i)
	return iu
}

// Mutation returns the InvoiceMutation object of the builder.
func (iu *InvoiceUpdate) Mutation() *InvoiceMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InvoiceUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InvoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InvoiceUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InvoiceUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InvoiceUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := invoice.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

func (iu *InvoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(invoice.Table, invoice.Columns, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeString))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Status(); ok {
		_spec.SetField(invoice.FieldStatus, field.TypeString, value)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(invoice.FieldUpdatedAt, field.TypeTime, value)
	}
	if iu.mutation.CreatedByCleared() {
		_spec.ClearField(invoice.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iu.mutation.UpdatedBy(); ok {
		_spec.SetField(invoice.FieldUpdatedBy, field.TypeString, value)
	}
	if iu.mutation.UpdatedByCleared() {
		_spec.ClearField(invoice.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iu.mutation.SubscriptionID(); ok {
		_spec.SetField(invoice.FieldSubscriptionID, field.TypeString, value)
	}
	if iu.mutation.SubscriptionIDCleared() {
		_spec.ClearField(invoice.FieldSubscriptionID, field.TypeString)
	}
	if value, ok := iu.mutation.InvoiceStatus(); ok {
		_spec.SetField(invoice.FieldInvoiceStatus, field.TypeString, value)
	}
	if value, ok := iu.mutation.PaymentStatus(); ok {
		_spec.SetField(invoice.FieldPaymentStatus, field.TypeString, value)
	}
	if value, ok := iu.mutation.AmountDue(); ok {
		_spec.SetField(invoice.FieldAmountDue, field.TypeOther, value)
	}
	if value, ok := iu.mutation.AmountPaid(); ok {
		_spec.SetField(invoice.FieldAmountPaid, field.TypeOther, value)
	}
	if value, ok := iu.mutation.AmountRemaining(); ok {
		_spec.SetField(invoice.FieldAmountRemaining, field.TypeOther, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(invoice.FieldDescription, field.TypeString, value)
	}
	if iu.mutation.DescriptionCleared() {
		_spec.ClearField(invoice.FieldDescription, field.TypeString)
	}
	if value, ok := iu.mutation.DueDate(); ok {
		_spec.SetField(invoice.FieldDueDate, field.TypeTime, value)
	}
	if iu.mutation.DueDateCleared() {
		_spec.ClearField(invoice.FieldDueDate, field.TypeTime)
	}
	if value, ok := iu.mutation.PaidAt(); ok {
		_spec.SetField(invoice.FieldPaidAt, field.TypeTime, value)
	}
	if iu.mutation.PaidAtCleared() {
		_spec.ClearField(invoice.FieldPaidAt, field.TypeTime)
	}
	if value, ok := iu.mutation.VoidedAt(); ok {
		_spec.SetField(invoice.FieldVoidedAt, field.TypeTime, value)
	}
	if iu.mutation.VoidedAtCleared() {
		_spec.ClearField(invoice.FieldVoidedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.FinalizedAt(); ok {
		_spec.SetField(invoice.FieldFinalizedAt, field.TypeTime, value)
	}
	if iu.mutation.FinalizedAtCleared() {
		_spec.ClearField(invoice.FieldFinalizedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.InvoicePdfURL(); ok {
		_spec.SetField(invoice.FieldInvoicePdfURL, field.TypeString, value)
	}
	if iu.mutation.InvoicePdfURLCleared() {
		_spec.ClearField(invoice.FieldInvoicePdfURL, field.TypeString)
	}
	if value, ok := iu.mutation.BillingReason(); ok {
		_spec.SetField(invoice.FieldBillingReason, field.TypeString, value)
	}
	if iu.mutation.BillingReasonCleared() {
		_spec.ClearField(invoice.FieldBillingReason, field.TypeString)
	}
	if value, ok := iu.mutation.Metadata(); ok {
		_spec.SetField(invoice.FieldMetadata, field.TypeJSON, value)
	}
	if iu.mutation.MetadataCleared() {
		_spec.ClearField(invoice.FieldMetadata, field.TypeJSON)
	}
	if value, ok := iu.mutation.Version(); ok {
		_spec.SetField(invoice.FieldVersion, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedVersion(); ok {
		_spec.AddField(invoice.FieldVersion, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InvoiceUpdateOne is the builder for updating a single Invoice entity.
type InvoiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvoiceMutation
}

// SetStatus sets the "status" field.
func (iuo *InvoiceUpdateOne) SetStatus(s string) *InvoiceUpdateOne {
	iuo.mutation.SetStatus(s)
	return iuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableStatus(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetStatus(*s)
	}
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *InvoiceUpdateOne) SetUpdatedAt(t time.Time) *InvoiceUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetUpdatedBy sets the "updated_by" field.
func (iuo *InvoiceUpdateOne) SetUpdatedBy(s string) *InvoiceUpdateOne {
	iuo.mutation.SetUpdatedBy(s)
	return iuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableUpdatedBy(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetUpdatedBy(*s)
	}
	return iuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iuo *InvoiceUpdateOne) ClearUpdatedBy() *InvoiceUpdateOne {
	iuo.mutation.ClearUpdatedBy()
	return iuo
}

// SetSubscriptionID sets the "subscription_id" field.
func (iuo *InvoiceUpdateOne) SetSubscriptionID(s string) *InvoiceUpdateOne {
	iuo.mutation.SetSubscriptionID(s)
	return iuo
}

// SetNillableSubscriptionID sets the "subscription_id" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableSubscriptionID(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetSubscriptionID(*s)
	}
	return iuo
}

// ClearSubscriptionID clears the value of the "subscription_id" field.
func (iuo *InvoiceUpdateOne) ClearSubscriptionID() *InvoiceUpdateOne {
	iuo.mutation.ClearSubscriptionID()
	return iuo
}

// SetInvoiceStatus sets the "invoice_status" field.
func (iuo *InvoiceUpdateOne) SetInvoiceStatus(s string) *InvoiceUpdateOne {
	iuo.mutation.SetInvoiceStatus(s)
	return iuo
}

// SetNillableInvoiceStatus sets the "invoice_status" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableInvoiceStatus(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetInvoiceStatus(*s)
	}
	return iuo
}

// SetPaymentStatus sets the "payment_status" field.
func (iuo *InvoiceUpdateOne) SetPaymentStatus(s string) *InvoiceUpdateOne {
	iuo.mutation.SetPaymentStatus(s)
	return iuo
}

// SetNillablePaymentStatus sets the "payment_status" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillablePaymentStatus(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetPaymentStatus(*s)
	}
	return iuo
}

// SetAmountDue sets the "amount_due" field.
func (iuo *InvoiceUpdateOne) SetAmountDue(d decimal.Decimal) *InvoiceUpdateOne {
	iuo.mutation.SetAmountDue(d)
	return iuo
}

// SetNillableAmountDue sets the "amount_due" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableAmountDue(d *decimal.Decimal) *InvoiceUpdateOne {
	if d != nil {
		iuo.SetAmountDue(*d)
	}
	return iuo
}

// SetAmountPaid sets the "amount_paid" field.
func (iuo *InvoiceUpdateOne) SetAmountPaid(d decimal.Decimal) *InvoiceUpdateOne {
	iuo.mutation.SetAmountPaid(d)
	return iuo
}

// SetNillableAmountPaid sets the "amount_paid" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableAmountPaid(d *decimal.Decimal) *InvoiceUpdateOne {
	if d != nil {
		iuo.SetAmountPaid(*d)
	}
	return iuo
}

// SetAmountRemaining sets the "amount_remaining" field.
func (iuo *InvoiceUpdateOne) SetAmountRemaining(d decimal.Decimal) *InvoiceUpdateOne {
	iuo.mutation.SetAmountRemaining(d)
	return iuo
}

// SetNillableAmountRemaining sets the "amount_remaining" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableAmountRemaining(d *decimal.Decimal) *InvoiceUpdateOne {
	if d != nil {
		iuo.SetAmountRemaining(*d)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *InvoiceUpdateOne) SetDescription(s string) *InvoiceUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableDescription(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// ClearDescription clears the value of the "description" field.
func (iuo *InvoiceUpdateOne) ClearDescription() *InvoiceUpdateOne {
	iuo.mutation.ClearDescription()
	return iuo
}

// SetDueDate sets the "due_date" field.
func (iuo *InvoiceUpdateOne) SetDueDate(t time.Time) *InvoiceUpdateOne {
	iuo.mutation.SetDueDate(t)
	return iuo
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableDueDate(t *time.Time) *InvoiceUpdateOne {
	if t != nil {
		iuo.SetDueDate(*t)
	}
	return iuo
}

// ClearDueDate clears the value of the "due_date" field.
func (iuo *InvoiceUpdateOne) ClearDueDate() *InvoiceUpdateOne {
	iuo.mutation.ClearDueDate()
	return iuo
}

// SetPaidAt sets the "paid_at" field.
func (iuo *InvoiceUpdateOne) SetPaidAt(t time.Time) *InvoiceUpdateOne {
	iuo.mutation.SetPaidAt(t)
	return iuo
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillablePaidAt(t *time.Time) *InvoiceUpdateOne {
	if t != nil {
		iuo.SetPaidAt(*t)
	}
	return iuo
}

// ClearPaidAt clears the value of the "paid_at" field.
func (iuo *InvoiceUpdateOne) ClearPaidAt() *InvoiceUpdateOne {
	iuo.mutation.ClearPaidAt()
	return iuo
}

// SetVoidedAt sets the "voided_at" field.
func (iuo *InvoiceUpdateOne) SetVoidedAt(t time.Time) *InvoiceUpdateOne {
	iuo.mutation.SetVoidedAt(t)
	return iuo
}

// SetNillableVoidedAt sets the "voided_at" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableVoidedAt(t *time.Time) *InvoiceUpdateOne {
	if t != nil {
		iuo.SetVoidedAt(*t)
	}
	return iuo
}

// ClearVoidedAt clears the value of the "voided_at" field.
func (iuo *InvoiceUpdateOne) ClearVoidedAt() *InvoiceUpdateOne {
	iuo.mutation.ClearVoidedAt()
	return iuo
}

// SetFinalizedAt sets the "finalized_at" field.
func (iuo *InvoiceUpdateOne) SetFinalizedAt(t time.Time) *InvoiceUpdateOne {
	iuo.mutation.SetFinalizedAt(t)
	return iuo
}

// SetNillableFinalizedAt sets the "finalized_at" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableFinalizedAt(t *time.Time) *InvoiceUpdateOne {
	if t != nil {
		iuo.SetFinalizedAt(*t)
	}
	return iuo
}

// ClearFinalizedAt clears the value of the "finalized_at" field.
func (iuo *InvoiceUpdateOne) ClearFinalizedAt() *InvoiceUpdateOne {
	iuo.mutation.ClearFinalizedAt()
	return iuo
}

// SetInvoicePdfURL sets the "invoice_pdf_url" field.
func (iuo *InvoiceUpdateOne) SetInvoicePdfURL(s string) *InvoiceUpdateOne {
	iuo.mutation.SetInvoicePdfURL(s)
	return iuo
}

// SetNillableInvoicePdfURL sets the "invoice_pdf_url" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableInvoicePdfURL(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetInvoicePdfURL(*s)
	}
	return iuo
}

// ClearInvoicePdfURL clears the value of the "invoice_pdf_url" field.
func (iuo *InvoiceUpdateOne) ClearInvoicePdfURL() *InvoiceUpdateOne {
	iuo.mutation.ClearInvoicePdfURL()
	return iuo
}

// SetBillingReason sets the "billing_reason" field.
func (iuo *InvoiceUpdateOne) SetBillingReason(s string) *InvoiceUpdateOne {
	iuo.mutation.SetBillingReason(s)
	return iuo
}

// SetNillableBillingReason sets the "billing_reason" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableBillingReason(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetBillingReason(*s)
	}
	return iuo
}

// ClearBillingReason clears the value of the "billing_reason" field.
func (iuo *InvoiceUpdateOne) ClearBillingReason() *InvoiceUpdateOne {
	iuo.mutation.ClearBillingReason()
	return iuo
}

// SetMetadata sets the "metadata" field.
func (iuo *InvoiceUpdateOne) SetMetadata(m map[string]interface{}) *InvoiceUpdateOne {
	iuo.mutation.SetMetadata(m)
	return iuo
}

// ClearMetadata clears the value of the "metadata" field.
func (iuo *InvoiceUpdateOne) ClearMetadata() *InvoiceUpdateOne {
	iuo.mutation.ClearMetadata()
	return iuo
}

// SetVersion sets the "version" field.
func (iuo *InvoiceUpdateOne) SetVersion(i int) *InvoiceUpdateOne {
	iuo.mutation.ResetVersion()
	iuo.mutation.SetVersion(i)
	return iuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableVersion(i *int) *InvoiceUpdateOne {
	if i != nil {
		iuo.SetVersion(*i)
	}
	return iuo
}

// AddVersion adds i to the "version" field.
func (iuo *InvoiceUpdateOne) AddVersion(i int) *InvoiceUpdateOne {
	iuo.mutation.AddVersion(i)
	return iuo
}

// Mutation returns the InvoiceMutation object of the builder.
func (iuo *InvoiceUpdateOne) Mutation() *InvoiceMutation {
	return iuo.mutation
}

// Where appends a list predicates to the InvoiceUpdate builder.
func (iuo *InvoiceUpdateOne) Where(ps ...predicate.Invoice) *InvoiceUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InvoiceUpdateOne) Select(field string, fields ...string) *InvoiceUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Invoice entity.
func (iuo *InvoiceUpdateOne) Save(ctx context.Context) (*Invoice, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InvoiceUpdateOne) SaveX(ctx context.Context) *Invoice {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InvoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InvoiceUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InvoiceUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := invoice.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

func (iuo *InvoiceUpdateOne) sqlSave(ctx context.Context) (_node *Invoice, err error) {
	_spec := sqlgraph.NewUpdateSpec(invoice.Table, invoice.Columns, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeString))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Invoice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoice.FieldID)
		for _, f := range fields {
			if !invoice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Status(); ok {
		_spec.SetField(invoice.FieldStatus, field.TypeString, value)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(invoice.FieldUpdatedAt, field.TypeTime, value)
	}
	if iuo.mutation.CreatedByCleared() {
		_spec.ClearField(invoice.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.UpdatedBy(); ok {
		_spec.SetField(invoice.FieldUpdatedBy, field.TypeString, value)
	}
	if iuo.mutation.UpdatedByCleared() {
		_spec.ClearField(invoice.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.SubscriptionID(); ok {
		_spec.SetField(invoice.FieldSubscriptionID, field.TypeString, value)
	}
	if iuo.mutation.SubscriptionIDCleared() {
		_spec.ClearField(invoice.FieldSubscriptionID, field.TypeString)
	}
	if value, ok := iuo.mutation.InvoiceStatus(); ok {
		_spec.SetField(invoice.FieldInvoiceStatus, field.TypeString, value)
	}
	if value, ok := iuo.mutation.PaymentStatus(); ok {
		_spec.SetField(invoice.FieldPaymentStatus, field.TypeString, value)
	}
	if value, ok := iuo.mutation.AmountDue(); ok {
		_spec.SetField(invoice.FieldAmountDue, field.TypeOther, value)
	}
	if value, ok := iuo.mutation.AmountPaid(); ok {
		_spec.SetField(invoice.FieldAmountPaid, field.TypeOther, value)
	}
	if value, ok := iuo.mutation.AmountRemaining(); ok {
		_spec.SetField(invoice.FieldAmountRemaining, field.TypeOther, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(invoice.FieldDescription, field.TypeString, value)
	}
	if iuo.mutation.DescriptionCleared() {
		_spec.ClearField(invoice.FieldDescription, field.TypeString)
	}
	if value, ok := iuo.mutation.DueDate(); ok {
		_spec.SetField(invoice.FieldDueDate, field.TypeTime, value)
	}
	if iuo.mutation.DueDateCleared() {
		_spec.ClearField(invoice.FieldDueDate, field.TypeTime)
	}
	if value, ok := iuo.mutation.PaidAt(); ok {
		_spec.SetField(invoice.FieldPaidAt, field.TypeTime, value)
	}
	if iuo.mutation.PaidAtCleared() {
		_spec.ClearField(invoice.FieldPaidAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.VoidedAt(); ok {
		_spec.SetField(invoice.FieldVoidedAt, field.TypeTime, value)
	}
	if iuo.mutation.VoidedAtCleared() {
		_spec.ClearField(invoice.FieldVoidedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.FinalizedAt(); ok {
		_spec.SetField(invoice.FieldFinalizedAt, field.TypeTime, value)
	}
	if iuo.mutation.FinalizedAtCleared() {
		_spec.ClearField(invoice.FieldFinalizedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.InvoicePdfURL(); ok {
		_spec.SetField(invoice.FieldInvoicePdfURL, field.TypeString, value)
	}
	if iuo.mutation.InvoicePdfURLCleared() {
		_spec.ClearField(invoice.FieldInvoicePdfURL, field.TypeString)
	}
	if value, ok := iuo.mutation.BillingReason(); ok {
		_spec.SetField(invoice.FieldBillingReason, field.TypeString, value)
	}
	if iuo.mutation.BillingReasonCleared() {
		_spec.ClearField(invoice.FieldBillingReason, field.TypeString)
	}
	if value, ok := iuo.mutation.Metadata(); ok {
		_spec.SetField(invoice.FieldMetadata, field.TypeJSON, value)
	}
	if iuo.mutation.MetadataCleared() {
		_spec.ClearField(invoice.FieldMetadata, field.TypeJSON)
	}
	if value, ok := iuo.mutation.Version(); ok {
		_spec.SetField(invoice.FieldVersion, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedVersion(); ok {
		_spec.AddField(invoice.FieldVersion, field.TypeInt, value)
	}
	_node = &Invoice{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}