// Code generated by ent, DO NOT EDIT.

package wallettransaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the wallettransaction type in the database.
	Label = "wallet_transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldWalletID holds the string denoting the wallet_id field in the database.
	FieldWalletID = "wallet_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldBalanceBefore holds the string denoting the balance_before field in the database.
	FieldBalanceBefore = "balance_before"
	// FieldBalanceAfter holds the string denoting the balance_after field in the database.
	FieldBalanceAfter = "balance_after"
	// FieldReferenceType holds the string denoting the reference_type field in the database.
	FieldReferenceType = "reference_type"
	// FieldReferenceID holds the string denoting the reference_id field in the database.
	FieldReferenceID = "reference_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldTransactionStatus holds the string denoting the transaction_status field in the database.
	FieldTransactionStatus = "transaction_status"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// Table holds the table name of the wallettransaction in the database.
	Table = "wallet_transactions"
)

// Columns holds all SQL columns for wallettransaction fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldWalletID,
	FieldType,
	FieldAmount,
	FieldBalanceBefore,
	FieldBalanceAfter,
	FieldReferenceType,
	FieldReferenceID,
	FieldDescription,
	FieldMetadata,
	FieldTransactionStatus,
	FieldStatus,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TenantIDValidator is a validator for the "tenant_id" field. It is called by the builders before save.
	TenantIDValidator func(string) error
	// WalletIDValidator is a validator for the "wallet_id" field. It is called by the builders before save.
	WalletIDValidator func(string) error
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultBalanceBefore holds the default value on creation for the "balance_before" field.
	DefaultBalanceBefore decimal.Decimal
	// DefaultBalanceAfter holds the default value on creation for the "balance_after" field.
	DefaultBalanceAfter decimal.Decimal
	// DefaultTransactionStatus holds the default value on creation for the "transaction_status" field.
	DefaultTransactionStatus string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the WalletTransaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByWalletID orders the results by the wallet_id field.
func ByWalletID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWalletID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByBalanceBefore orders the results by the balance_before field.
func ByBalanceBefore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalanceBefore, opts...).ToFunc()
}

// ByBalanceAfter orders the results by the balance_after field.
func ByBalanceAfter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalanceAfter, opts...).ToFunc()
}

// ByReferenceType orders the results by the reference_type field.
func ByReferenceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReferenceType, opts...).ToFunc()
}

// ByReferenceID orders the results by the reference_id field.
func ByReferenceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReferenceID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByTransactionStatus orders the results by the transaction_status field.
func ByTransactionStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionStatus, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}