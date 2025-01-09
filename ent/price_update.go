// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/flexprice/flexprice/ent/predicate"
	"github.com/flexprice/flexprice/ent/price"
	"github.com/flexprice/flexprice/ent/schema"
)

// PriceUpdate is the builder for updating Price entities.
type PriceUpdate struct {
	config
	hooks    []Hook
	mutation *PriceMutation
}

// Where appends a list predicates to the PriceUpdate builder.
func (pu *PriceUpdate) Where(ps ...predicate.Price) *PriceUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PriceUpdate) SetStatus(s string) *PriceUpdate {
	pu.mutation.SetStatus(s)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableStatus(s *string) *PriceUpdate {
	if s != nil {
		pu.SetStatus(*s)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PriceUpdate) SetUpdatedAt(t time.Time) *PriceUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PriceUpdate) SetUpdatedBy(s string) *PriceUpdate {
	pu.mutation.SetUpdatedBy(s)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableUpdatedBy(s *string) *PriceUpdate {
	if s != nil {
		pu.SetUpdatedBy(*s)
	}
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PriceUpdate) ClearUpdatedBy() *PriceUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PriceUpdate) SetAmount(f float64) *PriceUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableAmount(f *float64) *PriceUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PriceUpdate) AddAmount(f float64) *PriceUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetCurrency sets the "currency" field.
func (pu *PriceUpdate) SetCurrency(s string) *PriceUpdate {
	pu.mutation.SetCurrency(s)
	return pu
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableCurrency(s *string) *PriceUpdate {
	if s != nil {
		pu.SetCurrency(*s)
	}
	return pu
}

// SetDisplayAmount sets the "display_amount" field.
func (pu *PriceUpdate) SetDisplayAmount(s string) *PriceUpdate {
	pu.mutation.SetDisplayAmount(s)
	return pu
}

// SetNillableDisplayAmount sets the "display_amount" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableDisplayAmount(s *string) *PriceUpdate {
	if s != nil {
		pu.SetDisplayAmount(*s)
	}
	return pu
}

// SetPlanID sets the "plan_id" field.
func (pu *PriceUpdate) SetPlanID(s string) *PriceUpdate {
	pu.mutation.SetPlanID(s)
	return pu
}

// SetNillablePlanID sets the "plan_id" field if the given value is not nil.
func (pu *PriceUpdate) SetNillablePlanID(s *string) *PriceUpdate {
	if s != nil {
		pu.SetPlanID(*s)
	}
	return pu
}

// SetType sets the "type" field.
func (pu *PriceUpdate) SetType(s string) *PriceUpdate {
	pu.mutation.SetType(s)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableType(s *string) *PriceUpdate {
	if s != nil {
		pu.SetType(*s)
	}
	return pu
}

// SetBillingPeriod sets the "billing_period" field.
func (pu *PriceUpdate) SetBillingPeriod(s string) *PriceUpdate {
	pu.mutation.SetBillingPeriod(s)
	return pu
}

// SetNillableBillingPeriod sets the "billing_period" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableBillingPeriod(s *string) *PriceUpdate {
	if s != nil {
		pu.SetBillingPeriod(*s)
	}
	return pu
}

// SetBillingPeriodCount sets the "billing_period_count" field.
func (pu *PriceUpdate) SetBillingPeriodCount(i int) *PriceUpdate {
	pu.mutation.ResetBillingPeriodCount()
	pu.mutation.SetBillingPeriodCount(i)
	return pu
}

// SetNillableBillingPeriodCount sets the "billing_period_count" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableBillingPeriodCount(i *int) *PriceUpdate {
	if i != nil {
		pu.SetBillingPeriodCount(*i)
	}
	return pu
}

// AddBillingPeriodCount adds i to the "billing_period_count" field.
func (pu *PriceUpdate) AddBillingPeriodCount(i int) *PriceUpdate {
	pu.mutation.AddBillingPeriodCount(i)
	return pu
}

// SetBillingModel sets the "billing_model" field.
func (pu *PriceUpdate) SetBillingModel(s string) *PriceUpdate {
	pu.mutation.SetBillingModel(s)
	return pu
}

// SetNillableBillingModel sets the "billing_model" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableBillingModel(s *string) *PriceUpdate {
	if s != nil {
		pu.SetBillingModel(*s)
	}
	return pu
}

// SetBillingCadence sets the "billing_cadence" field.
func (pu *PriceUpdate) SetBillingCadence(s string) *PriceUpdate {
	pu.mutation.SetBillingCadence(s)
	return pu
}

// SetNillableBillingCadence sets the "billing_cadence" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableBillingCadence(s *string) *PriceUpdate {
	if s != nil {
		pu.SetBillingCadence(*s)
	}
	return pu
}

// SetMeterID sets the "meter_id" field.
func (pu *PriceUpdate) SetMeterID(s string) *PriceUpdate {
	pu.mutation.SetMeterID(s)
	return pu
}

// SetNillableMeterID sets the "meter_id" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableMeterID(s *string) *PriceUpdate {
	if s != nil {
		pu.SetMeterID(*s)
	}
	return pu
}

// ClearMeterID clears the value of the "meter_id" field.
func (pu *PriceUpdate) ClearMeterID() *PriceUpdate {
	pu.mutation.ClearMeterID()
	return pu
}

// SetFilterValues sets the "filter_values" field.
func (pu *PriceUpdate) SetFilterValues(m map[string][]string) *PriceUpdate {
	pu.mutation.SetFilterValues(m)
	return pu
}

// ClearFilterValues clears the value of the "filter_values" field.
func (pu *PriceUpdate) ClearFilterValues() *PriceUpdate {
	pu.mutation.ClearFilterValues()
	return pu
}

// SetTierMode sets the "tier_mode" field.
func (pu *PriceUpdate) SetTierMode(s string) *PriceUpdate {
	pu.mutation.SetTierMode(s)
	return pu
}

// SetNillableTierMode sets the "tier_mode" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableTierMode(s *string) *PriceUpdate {
	if s != nil {
		pu.SetTierMode(*s)
	}
	return pu
}

// ClearTierMode clears the value of the "tier_mode" field.
func (pu *PriceUpdate) ClearTierMode() *PriceUpdate {
	pu.mutation.ClearTierMode()
	return pu
}

// SetTiers sets the "tiers" field.
func (pu *PriceUpdate) SetTiers(st []schema.PriceTier) *PriceUpdate {
	pu.mutation.SetTiers(st)
	return pu
}

// AppendTiers appends st to the "tiers" field.
func (pu *PriceUpdate) AppendTiers(st []schema.PriceTier) *PriceUpdate {
	pu.mutation.AppendTiers(st)
	return pu
}

// ClearTiers clears the value of the "tiers" field.
func (pu *PriceUpdate) ClearTiers() *PriceUpdate {
	pu.mutation.ClearTiers()
	return pu
}

// SetTransformQuantity sets the "transform_quantity" field.
func (pu *PriceUpdate) SetTransformQuantity(sq schema.TransformQuantity) *PriceUpdate {
	pu.mutation.SetTransformQuantity(sq)
	return pu
}

// SetNillableTransformQuantity sets the "transform_quantity" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableTransformQuantity(sq *schema.TransformQuantity) *PriceUpdate {
	if sq != nil {
		pu.SetTransformQuantity(*sq)
	}
	return pu
}

// ClearTransformQuantity clears the value of the "transform_quantity" field.
func (pu *PriceUpdate) ClearTransformQuantity() *PriceUpdate {
	pu.mutation.ClearTransformQuantity()
	return pu
}

// SetLookupKey sets the "lookup_key" field.
func (pu *PriceUpdate) SetLookupKey(s string) *PriceUpdate {
	pu.mutation.SetLookupKey(s)
	return pu
}

// SetNillableLookupKey sets the "lookup_key" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableLookupKey(s *string) *PriceUpdate {
	if s != nil {
		pu.SetLookupKey(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PriceUpdate) SetDescription(s string) *PriceUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PriceUpdate) SetNillableDescription(s *string) *PriceUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PriceUpdate) ClearDescription() *PriceUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetMetadata sets the "metadata" field.
func (pu *PriceUpdate) SetMetadata(m map[string]string) *PriceUpdate {
	pu.mutation.SetMetadata(m)
	return pu
}

// ClearMetadata clears the value of the "metadata" field.
func (pu *PriceUpdate) ClearMetadata() *PriceUpdate {
	pu.mutation.ClearMetadata()
	return pu
}

// Mutation returns the PriceMutation object of the builder.
func (pu *PriceUpdate) Mutation() *PriceMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PriceUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PriceUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PriceUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PriceUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PriceUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := price.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PriceUpdate) check() error {
	if v, ok := pu.mutation.Currency(); ok {
		if err := price.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "Price.currency": %w`, err)}
		}
	}
	if v, ok := pu.mutation.DisplayAmount(); ok {
		if err := price.DisplayAmountValidator(v); err != nil {
			return &ValidationError{Name: "display_amount", err: fmt.Errorf(`ent: validator failed for field "Price.display_amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.PlanID(); ok {
		if err := price.PlanIDValidator(v); err != nil {
			return &ValidationError{Name: "plan_id", err: fmt.Errorf(`ent: validator failed for field "Price.plan_id": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := price.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Price.type": %w`, err)}
		}
	}
	if v, ok := pu.mutation.BillingPeriod(); ok {
		if err := price.BillingPeriodValidator(v); err != nil {
			return &ValidationError{Name: "billing_period", err: fmt.Errorf(`ent: validator failed for field "Price.billing_period": %w`, err)}
		}
	}
	if v, ok := pu.mutation.BillingPeriodCount(); ok {
		if err := price.BillingPeriodCountValidator(v); err != nil {
			return &ValidationError{Name: "billing_period_count", err: fmt.Errorf(`ent: validator failed for field "Price.billing_period_count": %w`, err)}
		}
	}
	if v, ok := pu.mutation.BillingModel(); ok {
		if err := price.BillingModelValidator(v); err != nil {
			return &ValidationError{Name: "billing_model", err: fmt.Errorf(`ent: validator failed for field "Price.billing_model": %w`, err)}
		}
	}
	if v, ok := pu.mutation.BillingCadence(); ok {
		if err := price.BillingCadenceValidator(v); err != nil {
			return &ValidationError{Name: "billing_cadence", err: fmt.Errorf(`ent: validator failed for field "Price.billing_cadence": %w`, err)}
		}
	}
	if v, ok := pu.mutation.LookupKey(); ok {
		if err := price.LookupKeyValidator(v); err != nil {
			return &ValidationError{Name: "lookup_key", err: fmt.Errorf(`ent: validator failed for field "Price.lookup_key": %w`, err)}
		}
	}
	return nil
}

func (pu *PriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(price.Table, price.Columns, sqlgraph.NewFieldSpec(price.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(price.FieldStatus, field.TypeString, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(price.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.CreatedByCleared() {
		_spec.ClearField(price.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(price.FieldUpdatedBy, field.TypeString, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(price.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(price.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(price.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Currency(); ok {
		_spec.SetField(price.FieldCurrency, field.TypeString, value)
	}
	if value, ok := pu.mutation.DisplayAmount(); ok {
		_spec.SetField(price.FieldDisplayAmount, field.TypeString, value)
	}
	if value, ok := pu.mutation.PlanID(); ok {
		_spec.SetField(price.FieldPlanID, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(price.FieldType, field.TypeString, value)
	}
	if value, ok := pu.mutation.BillingPeriod(); ok {
		_spec.SetField(price.FieldBillingPeriod, field.TypeString, value)
	}
	if value, ok := pu.mutation.BillingPeriodCount(); ok {
		_spec.SetField(price.FieldBillingPeriodCount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedBillingPeriodCount(); ok {
		_spec.AddField(price.FieldBillingPeriodCount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.BillingModel(); ok {
		_spec.SetField(price.FieldBillingModel, field.TypeString, value)
	}
	if value, ok := pu.mutation.BillingCadence(); ok {
		_spec.SetField(price.FieldBillingCadence, field.TypeString, value)
	}
	if value, ok := pu.mutation.MeterID(); ok {
		_spec.SetField(price.FieldMeterID, field.TypeString, value)
	}
	if pu.mutation.MeterIDCleared() {
		_spec.ClearField(price.FieldMeterID, field.TypeString)
	}
	if value, ok := pu.mutation.FilterValues(); ok {
		_spec.SetField(price.FieldFilterValues, field.TypeJSON, value)
	}
	if pu.mutation.FilterValuesCleared() {
		_spec.ClearField(price.FieldFilterValues, field.TypeJSON)
	}
	if value, ok := pu.mutation.TierMode(); ok {
		_spec.SetField(price.FieldTierMode, field.TypeString, value)
	}
	if pu.mutation.TierModeCleared() {
		_spec.ClearField(price.FieldTierMode, field.TypeString)
	}
	if value, ok := pu.mutation.Tiers(); ok {
		_spec.SetField(price.FieldTiers, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedTiers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, price.FieldTiers, value)
		})
	}
	if pu.mutation.TiersCleared() {
		_spec.ClearField(price.FieldTiers, field.TypeJSON)
	}
	if value, ok := pu.mutation.TransformQuantity(); ok {
		_spec.SetField(price.FieldTransformQuantity, field.TypeJSON, value)
	}
	if pu.mutation.TransformQuantityCleared() {
		_spec.ClearField(price.FieldTransformQuantity, field.TypeJSON)
	}
	if value, ok := pu.mutation.LookupKey(); ok {
		_spec.SetField(price.FieldLookupKey, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(price.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(price.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Metadata(); ok {
		_spec.SetField(price.FieldMetadata, field.TypeJSON, value)
	}
	if pu.mutation.MetadataCleared() {
		_spec.ClearField(price.FieldMetadata, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{price.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PriceUpdateOne is the builder for updating a single Price entity.
type PriceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PriceMutation
}

// SetStatus sets the "status" field.
func (puo *PriceUpdateOne) SetStatus(s string) *PriceUpdateOne {
	puo.mutation.SetStatus(s)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableStatus(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetStatus(*s)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PriceUpdateOne) SetUpdatedAt(t time.Time) *PriceUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PriceUpdateOne) SetUpdatedBy(s string) *PriceUpdateOne {
	puo.mutation.SetUpdatedBy(s)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableUpdatedBy(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetUpdatedBy(*s)
	}
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PriceUpdateOne) ClearUpdatedBy() *PriceUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PriceUpdateOne) SetAmount(f float64) *PriceUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableAmount(f *float64) *PriceUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PriceUpdateOne) AddAmount(f float64) *PriceUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetCurrency sets the "currency" field.
func (puo *PriceUpdateOne) SetCurrency(s string) *PriceUpdateOne {
	puo.mutation.SetCurrency(s)
	return puo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableCurrency(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetCurrency(*s)
	}
	return puo
}

// SetDisplayAmount sets the "display_amount" field.
func (puo *PriceUpdateOne) SetDisplayAmount(s string) *PriceUpdateOne {
	puo.mutation.SetDisplayAmount(s)
	return puo
}

// SetNillableDisplayAmount sets the "display_amount" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableDisplayAmount(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetDisplayAmount(*s)
	}
	return puo
}

// SetPlanID sets the "plan_id" field.
func (puo *PriceUpdateOne) SetPlanID(s string) *PriceUpdateOne {
	puo.mutation.SetPlanID(s)
	return puo
}

// SetNillablePlanID sets the "plan_id" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillablePlanID(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetPlanID(*s)
	}
	return puo
}

// SetType sets the "type" field.
func (puo *PriceUpdateOne) SetType(s string) *PriceUpdateOne {
	puo.mutation.SetType(s)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableType(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetType(*s)
	}
	return puo
}

// SetBillingPeriod sets the "billing_period" field.
func (puo *PriceUpdateOne) SetBillingPeriod(s string) *PriceUpdateOne {
	puo.mutation.SetBillingPeriod(s)
	return puo
}

// SetNillableBillingPeriod sets the "billing_period" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableBillingPeriod(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetBillingPeriod(*s)
	}
	return puo
}

// SetBillingPeriodCount sets the "billing_period_count" field.
func (puo *PriceUpdateOne) SetBillingPeriodCount(i int) *PriceUpdateOne {
	puo.mutation.ResetBillingPeriodCount()
	puo.mutation.SetBillingPeriodCount(i)
	return puo
}

// SetNillableBillingPeriodCount sets the "billing_period_count" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableBillingPeriodCount(i *int) *PriceUpdateOne {
	if i != nil {
		puo.SetBillingPeriodCount(*i)
	}
	return puo
}

// AddBillingPeriodCount adds i to the "billing_period_count" field.
func (puo *PriceUpdateOne) AddBillingPeriodCount(i int) *PriceUpdateOne {
	puo.mutation.AddBillingPeriodCount(i)
	return puo
}

// SetBillingModel sets the "billing_model" field.
func (puo *PriceUpdateOne) SetBillingModel(s string) *PriceUpdateOne {
	puo.mutation.SetBillingModel(s)
	return puo
}

// SetNillableBillingModel sets the "billing_model" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableBillingModel(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetBillingModel(*s)
	}
	return puo
}

// SetBillingCadence sets the "billing_cadence" field.
func (puo *PriceUpdateOne) SetBillingCadence(s string) *PriceUpdateOne {
	puo.mutation.SetBillingCadence(s)
	return puo
}

// SetNillableBillingCadence sets the "billing_cadence" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableBillingCadence(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetBillingCadence(*s)
	}
	return puo
}

// SetMeterID sets the "meter_id" field.
func (puo *PriceUpdateOne) SetMeterID(s string) *PriceUpdateOne {
	puo.mutation.SetMeterID(s)
	return puo
}

// SetNillableMeterID sets the "meter_id" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableMeterID(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetMeterID(*s)
	}
	return puo
}

// ClearMeterID clears the value of the "meter_id" field.
func (puo *PriceUpdateOne) ClearMeterID() *PriceUpdateOne {
	puo.mutation.ClearMeterID()
	return puo
}

// SetFilterValues sets the "filter_values" field.
func (puo *PriceUpdateOne) SetFilterValues(m map[string][]string) *PriceUpdateOne {
	puo.mutation.SetFilterValues(m)
	return puo
}

// ClearFilterValues clears the value of the "filter_values" field.
func (puo *PriceUpdateOne) ClearFilterValues() *PriceUpdateOne {
	puo.mutation.ClearFilterValues()
	return puo
}

// SetTierMode sets the "tier_mode" field.
func (puo *PriceUpdateOne) SetTierMode(s string) *PriceUpdateOne {
	puo.mutation.SetTierMode(s)
	return puo
}

// SetNillableTierMode sets the "tier_mode" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableTierMode(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetTierMode(*s)
	}
	return puo
}

// ClearTierMode clears the value of the "tier_mode" field.
func (puo *PriceUpdateOne) ClearTierMode() *PriceUpdateOne {
	puo.mutation.ClearTierMode()
	return puo
}

// SetTiers sets the "tiers" field.
func (puo *PriceUpdateOne) SetTiers(st []schema.PriceTier) *PriceUpdateOne {
	puo.mutation.SetTiers(st)
	return puo
}

// AppendTiers appends st to the "tiers" field.
func (puo *PriceUpdateOne) AppendTiers(st []schema.PriceTier) *PriceUpdateOne {
	puo.mutation.AppendTiers(st)
	return puo
}

// ClearTiers clears the value of the "tiers" field.
func (puo *PriceUpdateOne) ClearTiers() *PriceUpdateOne {
	puo.mutation.ClearTiers()
	return puo
}

// SetTransformQuantity sets the "transform_quantity" field.
func (puo *PriceUpdateOne) SetTransformQuantity(sq schema.TransformQuantity) *PriceUpdateOne {
	puo.mutation.SetTransformQuantity(sq)
	return puo
}

// SetNillableTransformQuantity sets the "transform_quantity" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableTransformQuantity(sq *schema.TransformQuantity) *PriceUpdateOne {
	if sq != nil {
		puo.SetTransformQuantity(*sq)
	}
	return puo
}

// ClearTransformQuantity clears the value of the "transform_quantity" field.
func (puo *PriceUpdateOne) ClearTransformQuantity() *PriceUpdateOne {
	puo.mutation.ClearTransformQuantity()
	return puo
}

// SetLookupKey sets the "lookup_key" field.
func (puo *PriceUpdateOne) SetLookupKey(s string) *PriceUpdateOne {
	puo.mutation.SetLookupKey(s)
	return puo
}

// SetNillableLookupKey sets the "lookup_key" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableLookupKey(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetLookupKey(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PriceUpdateOne) SetDescription(s string) *PriceUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PriceUpdateOne) SetNillableDescription(s *string) *PriceUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PriceUpdateOne) ClearDescription() *PriceUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetMetadata sets the "metadata" field.
func (puo *PriceUpdateOne) SetMetadata(m map[string]string) *PriceUpdateOne {
	puo.mutation.SetMetadata(m)
	return puo
}

// ClearMetadata clears the value of the "metadata" field.
func (puo *PriceUpdateOne) ClearMetadata() *PriceUpdateOne {
	puo.mutation.ClearMetadata()
	return puo
}

// Mutation returns the PriceMutation object of the builder.
func (puo *PriceUpdateOne) Mutation() *PriceMutation {
	return puo.mutation
}

// Where appends a list predicates to the PriceUpdate builder.
func (puo *PriceUpdateOne) Where(ps ...predicate.Price) *PriceUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PriceUpdateOne) Select(field string, fields ...string) *PriceUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Price entity.
func (puo *PriceUpdateOne) Save(ctx context.Context) (*Price, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PriceUpdateOne) SaveX(ctx context.Context) *Price {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PriceUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PriceUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PriceUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := price.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PriceUpdateOne) check() error {
	if v, ok := puo.mutation.Currency(); ok {
		if err := price.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "Price.currency": %w`, err)}
		}
	}
	if v, ok := puo.mutation.DisplayAmount(); ok {
		if err := price.DisplayAmountValidator(v); err != nil {
			return &ValidationError{Name: "display_amount", err: fmt.Errorf(`ent: validator failed for field "Price.display_amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.PlanID(); ok {
		if err := price.PlanIDValidator(v); err != nil {
			return &ValidationError{Name: "plan_id", err: fmt.Errorf(`ent: validator failed for field "Price.plan_id": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := price.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Price.type": %w`, err)}
		}
	}
	if v, ok := puo.mutation.BillingPeriod(); ok {
		if err := price.BillingPeriodValidator(v); err != nil {
			return &ValidationError{Name: "billing_period", err: fmt.Errorf(`ent: validator failed for field "Price.billing_period": %w`, err)}
		}
	}
	if v, ok := puo.mutation.BillingPeriodCount(); ok {
		if err := price.BillingPeriodCountValidator(v); err != nil {
			return &ValidationError{Name: "billing_period_count", err: fmt.Errorf(`ent: validator failed for field "Price.billing_period_count": %w`, err)}
		}
	}
	if v, ok := puo.mutation.BillingModel(); ok {
		if err := price.BillingModelValidator(v); err != nil {
			return &ValidationError{Name: "billing_model", err: fmt.Errorf(`ent: validator failed for field "Price.billing_model": %w`, err)}
		}
	}
	if v, ok := puo.mutation.BillingCadence(); ok {
		if err := price.BillingCadenceValidator(v); err != nil {
			return &ValidationError{Name: "billing_cadence", err: fmt.Errorf(`ent: validator failed for field "Price.billing_cadence": %w`, err)}
		}
	}
	if v, ok := puo.mutation.LookupKey(); ok {
		if err := price.LookupKeyValidator(v); err != nil {
			return &ValidationError{Name: "lookup_key", err: fmt.Errorf(`ent: validator failed for field "Price.lookup_key": %w`, err)}
		}
	}
	return nil
}

func (puo *PriceUpdateOne) sqlSave(ctx context.Context) (_node *Price, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(price.Table, price.Columns, sqlgraph.NewFieldSpec(price.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Price.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, price.FieldID)
		for _, f := range fields {
			if !price.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != price.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(price.FieldStatus, field.TypeString, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(price.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.CreatedByCleared() {
		_spec.ClearField(price.FieldCreatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(price.FieldUpdatedBy, field.TypeString, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(price.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(price.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(price.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Currency(); ok {
		_spec.SetField(price.FieldCurrency, field.TypeString, value)
	}
	if value, ok := puo.mutation.DisplayAmount(); ok {
		_spec.SetField(price.FieldDisplayAmount, field.TypeString, value)
	}
	if value, ok := puo.mutation.PlanID(); ok {
		_spec.SetField(price.FieldPlanID, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(price.FieldType, field.TypeString, value)
	}
	if value, ok := puo.mutation.BillingPeriod(); ok {
		_spec.SetField(price.FieldBillingPeriod, field.TypeString, value)
	}
	if value, ok := puo.mutation.BillingPeriodCount(); ok {
		_spec.SetField(price.FieldBillingPeriodCount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedBillingPeriodCount(); ok {
		_spec.AddField(price.FieldBillingPeriodCount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.BillingModel(); ok {
		_spec.SetField(price.FieldBillingModel, field.TypeString, value)
	}
	if value, ok := puo.mutation.BillingCadence(); ok {
		_spec.SetField(price.FieldBillingCadence, field.TypeString, value)
	}
	if value, ok := puo.mutation.MeterID(); ok {
		_spec.SetField(price.FieldMeterID, field.TypeString, value)
	}
	if puo.mutation.MeterIDCleared() {
		_spec.ClearField(price.FieldMeterID, field.TypeString)
	}
	if value, ok := puo.mutation.FilterValues(); ok {
		_spec.SetField(price.FieldFilterValues, field.TypeJSON, value)
	}
	if puo.mutation.FilterValuesCleared() {
		_spec.ClearField(price.FieldFilterValues, field.TypeJSON)
	}
	if value, ok := puo.mutation.TierMode(); ok {
		_spec.SetField(price.FieldTierMode, field.TypeString, value)
	}
	if puo.mutation.TierModeCleared() {
		_spec.ClearField(price.FieldTierMode, field.TypeString)
	}
	if value, ok := puo.mutation.Tiers(); ok {
		_spec.SetField(price.FieldTiers, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedTiers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, price.FieldTiers, value)
		})
	}
	if puo.mutation.TiersCleared() {
		_spec.ClearField(price.FieldTiers, field.TypeJSON)
	}
	if value, ok := puo.mutation.TransformQuantity(); ok {
		_spec.SetField(price.FieldTransformQuantity, field.TypeJSON, value)
	}
	if puo.mutation.TransformQuantityCleared() {
		_spec.ClearField(price.FieldTransformQuantity, field.TypeJSON)
	}
	if value, ok := puo.mutation.LookupKey(); ok {
		_spec.SetField(price.FieldLookupKey, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(price.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(price.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Metadata(); ok {
		_spec.SetField(price.FieldMetadata, field.TypeJSON, value)
	}
	if puo.mutation.MetadataCleared() {
		_spec.ClearField(price.FieldMetadata, field.TypeJSON)
	}
	_node = &Price{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{price.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}