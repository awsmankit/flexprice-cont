// Code generated by ent, DO NOT EDIT.

package wallet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/predicate"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldID, id))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldTenantID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedBy, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCustomerID, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCurrency, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldDescription, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldBalance, v))
}

// WalletStatus applies equality check predicate on the "wallet_status" field. It's identical to WalletStatusEQ.
func WalletStatus(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldWalletStatus, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldTenantID, v))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldTenantID, v))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldTenantID, v))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldTenantID, v))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldTenantID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldCustomerID, vs...))
}

// CustomerIDGT applies the GT predicate on the "customer_id" field.
func CustomerIDGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldCustomerID, v))
}

// CustomerIDGTE applies the GTE predicate on the "customer_id" field.
func CustomerIDGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldCustomerID, v))
}

// CustomerIDLT applies the LT predicate on the "customer_id" field.
func CustomerIDLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldCustomerID, v))
}

// CustomerIDLTE applies the LTE predicate on the "customer_id" field.
func CustomerIDLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldCustomerID, v))
}

// CustomerIDContains applies the Contains predicate on the "customer_id" field.
func CustomerIDContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldCustomerID, v))
}

// CustomerIDHasPrefix applies the HasPrefix predicate on the "customer_id" field.
func CustomerIDHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldCustomerID, v))
}

// CustomerIDHasSuffix applies the HasSuffix predicate on the "customer_id" field.
func CustomerIDHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldCustomerID, v))
}

// CustomerIDEqualFold applies the EqualFold predicate on the "customer_id" field.
func CustomerIDEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldCustomerID, v))
}

// CustomerIDContainsFold applies the ContainsFold predicate on the "customer_id" field.
func CustomerIDContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldCustomerID, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldCurrency, v))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldCurrency, v))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldCurrency, v))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldCurrency, v))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldCurrency, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldDescription, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Wallet {
	return predicate.Wallet(sql.FieldNotNull(FieldMetadata))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v decimal.Decimal) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldBalance, v))
}

// WalletStatusEQ applies the EQ predicate on the "wallet_status" field.
func WalletStatusEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldWalletStatus, v))
}

// WalletStatusNEQ applies the NEQ predicate on the "wallet_status" field.
func WalletStatusNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldWalletStatus, v))
}

// WalletStatusIn applies the In predicate on the "wallet_status" field.
func WalletStatusIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldWalletStatus, vs...))
}

// WalletStatusNotIn applies the NotIn predicate on the "wallet_status" field.
func WalletStatusNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldWalletStatus, vs...))
}

// WalletStatusGT applies the GT predicate on the "wallet_status" field.
func WalletStatusGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldWalletStatus, v))
}

// WalletStatusGTE applies the GTE predicate on the "wallet_status" field.
func WalletStatusGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldWalletStatus, v))
}

// WalletStatusLT applies the LT predicate on the "wallet_status" field.
func WalletStatusLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldWalletStatus, v))
}

// WalletStatusLTE applies the LTE predicate on the "wallet_status" field.
func WalletStatusLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldWalletStatus, v))
}

// WalletStatusContains applies the Contains predicate on the "wallet_status" field.
func WalletStatusContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldWalletStatus, v))
}

// WalletStatusHasPrefix applies the HasPrefix predicate on the "wallet_status" field.
func WalletStatusHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldWalletStatus, v))
}

// WalletStatusHasSuffix applies the HasSuffix predicate on the "wallet_status" field.
func WalletStatusHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldWalletStatus, v))
}

// WalletStatusEqualFold applies the EqualFold predicate on the "wallet_status" field.
func WalletStatusEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldWalletStatus, v))
}

// WalletStatusContainsFold applies the ContainsFold predicate on the "wallet_status" field.
func WalletStatusContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldWalletStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.NotPredicates(p))
}
