// Code generated by ent, DO NOT EDIT.

package plan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldID, id))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldTenantID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldUpdatedBy, v))
}

// LookupKey applies equality check predicate on the "lookup_key" field. It's identical to LookupKeyEQ.
func LookupKey(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldLookupKey, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldDescription, v))
}

// InvoiceCadence applies equality check predicate on the "invoice_cadence" field. It's identical to InvoiceCadenceEQ.
func InvoiceCadence(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldInvoiceCadence, v))
}

// TrialPeriod applies equality check predicate on the "trial_period" field. It's identical to TrialPeriodEQ.
func TrialPeriod(v int) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldTrialPeriod, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldTenantID, v))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldTenantID, v))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldTenantID, v))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldTenantID, v))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldTenantID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// LookupKeyEQ applies the EQ predicate on the "lookup_key" field.
func LookupKeyEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldLookupKey, v))
}

// LookupKeyNEQ applies the NEQ predicate on the "lookup_key" field.
func LookupKeyNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldLookupKey, v))
}

// LookupKeyIn applies the In predicate on the "lookup_key" field.
func LookupKeyIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldLookupKey, vs...))
}

// LookupKeyNotIn applies the NotIn predicate on the "lookup_key" field.
func LookupKeyNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldLookupKey, vs...))
}

// LookupKeyGT applies the GT predicate on the "lookup_key" field.
func LookupKeyGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldLookupKey, v))
}

// LookupKeyGTE applies the GTE predicate on the "lookup_key" field.
func LookupKeyGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldLookupKey, v))
}

// LookupKeyLT applies the LT predicate on the "lookup_key" field.
func LookupKeyLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldLookupKey, v))
}

// LookupKeyLTE applies the LTE predicate on the "lookup_key" field.
func LookupKeyLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldLookupKey, v))
}

// LookupKeyContains applies the Contains predicate on the "lookup_key" field.
func LookupKeyContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldLookupKey, v))
}

// LookupKeyHasPrefix applies the HasPrefix predicate on the "lookup_key" field.
func LookupKeyHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldLookupKey, v))
}

// LookupKeyHasSuffix applies the HasSuffix predicate on the "lookup_key" field.
func LookupKeyHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldLookupKey, v))
}

// LookupKeyEqualFold applies the EqualFold predicate on the "lookup_key" field.
func LookupKeyEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldLookupKey, v))
}

// LookupKeyContainsFold applies the ContainsFold predicate on the "lookup_key" field.
func LookupKeyContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldLookupKey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldDescription, v))
}

// InvoiceCadenceEQ applies the EQ predicate on the "invoice_cadence" field.
func InvoiceCadenceEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldInvoiceCadence, v))
}

// InvoiceCadenceNEQ applies the NEQ predicate on the "invoice_cadence" field.
func InvoiceCadenceNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldInvoiceCadence, v))
}

// InvoiceCadenceIn applies the In predicate on the "invoice_cadence" field.
func InvoiceCadenceIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldInvoiceCadence, vs...))
}

// InvoiceCadenceNotIn applies the NotIn predicate on the "invoice_cadence" field.
func InvoiceCadenceNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldInvoiceCadence, vs...))
}

// InvoiceCadenceGT applies the GT predicate on the "invoice_cadence" field.
func InvoiceCadenceGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldInvoiceCadence, v))
}

// InvoiceCadenceGTE applies the GTE predicate on the "invoice_cadence" field.
func InvoiceCadenceGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldInvoiceCadence, v))
}

// InvoiceCadenceLT applies the LT predicate on the "invoice_cadence" field.
func InvoiceCadenceLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldInvoiceCadence, v))
}

// InvoiceCadenceLTE applies the LTE predicate on the "invoice_cadence" field.
func InvoiceCadenceLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldInvoiceCadence, v))
}

// InvoiceCadenceContains applies the Contains predicate on the "invoice_cadence" field.
func InvoiceCadenceContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldInvoiceCadence, v))
}

// InvoiceCadenceHasPrefix applies the HasPrefix predicate on the "invoice_cadence" field.
func InvoiceCadenceHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldInvoiceCadence, v))
}

// InvoiceCadenceHasSuffix applies the HasSuffix predicate on the "invoice_cadence" field.
func InvoiceCadenceHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldInvoiceCadence, v))
}

// InvoiceCadenceEqualFold applies the EqualFold predicate on the "invoice_cadence" field.
func InvoiceCadenceEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldInvoiceCadence, v))
}

// InvoiceCadenceContainsFold applies the ContainsFold predicate on the "invoice_cadence" field.
func InvoiceCadenceContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldInvoiceCadence, v))
}

// TrialPeriodEQ applies the EQ predicate on the "trial_period" field.
func TrialPeriodEQ(v int) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldTrialPeriod, v))
}

// TrialPeriodNEQ applies the NEQ predicate on the "trial_period" field.
func TrialPeriodNEQ(v int) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldTrialPeriod, v))
}

// TrialPeriodIn applies the In predicate on the "trial_period" field.
func TrialPeriodIn(vs ...int) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldTrialPeriod, vs...))
}

// TrialPeriodNotIn applies the NotIn predicate on the "trial_period" field.
func TrialPeriodNotIn(vs ...int) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldTrialPeriod, vs...))
}

// TrialPeriodGT applies the GT predicate on the "trial_period" field.
func TrialPeriodGT(v int) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldTrialPeriod, v))
}

// TrialPeriodGTE applies the GTE predicate on the "trial_period" field.
func TrialPeriodGTE(v int) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldTrialPeriod, v))
}

// TrialPeriodLT applies the LT predicate on the "trial_period" field.
func TrialPeriodLT(v int) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldTrialPeriod, v))
}

// TrialPeriodLTE applies the LTE predicate on the "trial_period" field.
func TrialPeriodLTE(v int) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldTrialPeriod, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Plan) predicate.Plan {
	return predicate.Plan(sql.NotPredicates(p))
}