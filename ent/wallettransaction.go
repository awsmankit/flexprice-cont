// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/wallettransaction"
	"github.com/shopspring/decimal"
)

// WalletTransaction is the model entity for the WalletTransaction schema.
type WalletTransaction struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// WalletID holds the value of the "wallet_id" field.
	WalletID string `json:"wallet_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// BalanceBefore holds the value of the "balance_before" field.
	BalanceBefore decimal.Decimal `json:"balance_before,omitempty"`
	// BalanceAfter holds the value of the "balance_after" field.
	BalanceAfter decimal.Decimal `json:"balance_after,omitempty"`
	// ReferenceType holds the value of the "reference_type" field.
	ReferenceType string `json:"reference_type,omitempty"`
	// ReferenceID holds the value of the "reference_id" field.
	ReferenceID string `json:"reference_id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]string `json:"metadata,omitempty"`
	// TransactionStatus holds the value of the "transaction_status" field.
	TransactionStatus string `json:"transaction_status,omitempty"`
	selectValues      sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WalletTransaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wallettransaction.FieldMetadata:
			values[i] = new([]byte)
		case wallettransaction.FieldAmount, wallettransaction.FieldBalanceBefore, wallettransaction.FieldBalanceAfter:
			values[i] = new(decimal.Decimal)
		case wallettransaction.FieldID, wallettransaction.FieldTenantID, wallettransaction.FieldStatus, wallettransaction.FieldCreatedBy, wallettransaction.FieldUpdatedBy, wallettransaction.FieldWalletID, wallettransaction.FieldType, wallettransaction.FieldReferenceType, wallettransaction.FieldReferenceID, wallettransaction.FieldDescription, wallettransaction.FieldTransactionStatus:
			values[i] = new(sql.NullString)
		case wallettransaction.FieldCreatedAt, wallettransaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WalletTransaction fields.
func (wt *WalletTransaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wallettransaction.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				wt.ID = value.String
			}
		case wallettransaction.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				wt.TenantID = value.String
			}
		case wallettransaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				wt.Status = value.String
			}
		case wallettransaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wt.CreatedAt = value.Time
			}
		case wallettransaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wt.UpdatedAt = value.Time
			}
		case wallettransaction.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				wt.CreatedBy = value.String
			}
		case wallettransaction.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				wt.UpdatedBy = value.String
			}
		case wallettransaction.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_id", values[i])
			} else if value.Valid {
				wt.WalletID = value.String
			}
		case wallettransaction.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				wt.Type = value.String
			}
		case wallettransaction.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				wt.Amount = *value
			}
		case wallettransaction.FieldBalanceBefore:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field balance_before", values[i])
			} else if value != nil {
				wt.BalanceBefore = *value
			}
		case wallettransaction.FieldBalanceAfter:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field balance_after", values[i])
			} else if value != nil {
				wt.BalanceAfter = *value
			}
		case wallettransaction.FieldReferenceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_type", values[i])
			} else if value.Valid {
				wt.ReferenceType = value.String
			}
		case wallettransaction.FieldReferenceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_id", values[i])
			} else if value.Valid {
				wt.ReferenceID = value.String
			}
		case wallettransaction.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				wt.Description = value.String
			}
		case wallettransaction.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wt.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case wallettransaction.FieldTransactionStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_status", values[i])
			} else if value.Valid {
				wt.TransactionStatus = value.String
			}
		default:
			wt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WalletTransaction.
// This includes values selected through modifiers, order, etc.
func (wt *WalletTransaction) Value(name string) (ent.Value, error) {
	return wt.selectValues.Get(name)
}

// Update returns a builder for updating this WalletTransaction.
// Note that you need to call WalletTransaction.Unwrap() before calling this method if this WalletTransaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (wt *WalletTransaction) Update() *WalletTransactionUpdateOne {
	return NewWalletTransactionClient(wt.config).UpdateOne(wt)
}

// Unwrap unwraps the WalletTransaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wt *WalletTransaction) Unwrap() *WalletTransaction {
	_tx, ok := wt.config.driver.(*txDriver)
	if !ok {
		panic("ent: WalletTransaction is not a transactional entity")
	}
	wt.config.driver = _tx.drv
	return wt
}

// String implements the fmt.Stringer.
func (wt *WalletTransaction) String() string {
	var builder strings.Builder
	builder.WriteString("WalletTransaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wt.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(wt.TenantID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(wt.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(wt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(wt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(wt.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(wt.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("wallet_id=")
	builder.WriteString(wt.WalletID)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(wt.Type)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", wt.Amount))
	builder.WriteString(", ")
	builder.WriteString("balance_before=")
	builder.WriteString(fmt.Sprintf("%v", wt.BalanceBefore))
	builder.WriteString(", ")
	builder.WriteString("balance_after=")
	builder.WriteString(fmt.Sprintf("%v", wt.BalanceAfter))
	builder.WriteString(", ")
	builder.WriteString("reference_type=")
	builder.WriteString(wt.ReferenceType)
	builder.WriteString(", ")
	builder.WriteString("reference_id=")
	builder.WriteString(wt.ReferenceID)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(wt.Description)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", wt.Metadata))
	builder.WriteString(", ")
	builder.WriteString("transaction_status=")
	builder.WriteString(wt.TransactionStatus)
	builder.WriteByte(')')
	return builder.String()
}

// WalletTransactions is a parsable slice of WalletTransaction.
type WalletTransactions []*WalletTransaction
