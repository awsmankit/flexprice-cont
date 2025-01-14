// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/billingsequence"
)

// BillingSequence is the model entity for the BillingSequence schema.
type BillingSequence struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// SubscriptionID holds the value of the "subscription_id" field.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// LastSequence holds the value of the "last_sequence" field.
	LastSequence int `json:"last_sequence,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillingSequence) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billingsequence.FieldID, billingsequence.FieldLastSequence:
			values[i] = new(sql.NullInt64)
		case billingsequence.FieldTenantID, billingsequence.FieldSubscriptionID:
			values[i] = new(sql.NullString)
		case billingsequence.FieldCreatedAt, billingsequence.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillingSequence fields.
func (bs *BillingSequence) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billingsequence.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bs.ID = int(value.Int64)
		case billingsequence.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				bs.TenantID = value.String
			}
		case billingsequence.FieldSubscriptionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subscription_id", values[i])
			} else if value.Valid {
				bs.SubscriptionID = value.String
			}
		case billingsequence.FieldLastSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_sequence", values[i])
			} else if value.Valid {
				bs.LastSequence = int(value.Int64)
			}
		case billingsequence.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bs.CreatedAt = value.Time
			}
		case billingsequence.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bs.UpdatedAt = value.Time
			}
		default:
			bs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillingSequence.
// This includes values selected through modifiers, order, etc.
func (bs *BillingSequence) Value(name string) (ent.Value, error) {
	return bs.selectValues.Get(name)
}

// Update returns a builder for updating this BillingSequence.
// Note that you need to call BillingSequence.Unwrap() before calling this method if this BillingSequence
// was returned from a transaction, and the transaction was committed or rolled back.
func (bs *BillingSequence) Update() *BillingSequenceUpdateOne {
	return NewBillingSequenceClient(bs.config).UpdateOne(bs)
}

// Unwrap unwraps the BillingSequence entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bs *BillingSequence) Unwrap() *BillingSequence {
	_tx, ok := bs.config.driver.(*txDriver)
	if !ok {
		panic("ent: BillingSequence is not a transactional entity")
	}
	bs.config.driver = _tx.drv
	return bs
}

// String implements the fmt.Stringer.
func (bs *BillingSequence) String() string {
	var builder strings.Builder
	builder.WriteString("BillingSequence(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bs.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(bs.TenantID)
	builder.WriteString(", ")
	builder.WriteString("subscription_id=")
	builder.WriteString(bs.SubscriptionID)
	builder.WriteString(", ")
	builder.WriteString("last_sequence=")
	builder.WriteString(fmt.Sprintf("%v", bs.LastSequence))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bs.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BillingSequences is a parsable slice of BillingSequence.
type BillingSequences []*BillingSequence
