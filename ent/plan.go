// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/plan"
)

// Plan is the model entity for the Plan schema.
type Plan struct {
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
	// LookupKey holds the value of the "lookup_key" field.
	LookupKey string `json:"lookup_key,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// InvoiceCadence holds the value of the "invoice_cadence" field.
	InvoiceCadence string `json:"invoice_cadence,omitempty"`
	// TrialPeriod holds the value of the "trial_period" field.
	TrialPeriod  int `json:"trial_period,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Plan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case plan.FieldTrialPeriod:
			values[i] = new(sql.NullInt64)
		case plan.FieldID, plan.FieldTenantID, plan.FieldStatus, plan.FieldCreatedBy, plan.FieldUpdatedBy, plan.FieldLookupKey, plan.FieldName, plan.FieldDescription, plan.FieldInvoiceCadence:
			values[i] = new(sql.NullString)
		case plan.FieldCreatedAt, plan.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Plan fields.
func (pl *Plan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plan.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pl.ID = value.String
			}
		case plan.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				pl.TenantID = value.String
			}
		case plan.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pl.Status = value.String
			}
		case plan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		case plan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case plan.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pl.CreatedBy = value.String
			}
		case plan.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pl.UpdatedBy = value.String
			}
		case plan.FieldLookupKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lookup_key", values[i])
			} else if value.Valid {
				pl.LookupKey = value.String
			}
		case plan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case plan.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pl.Description = value.String
			}
		case plan.FieldInvoiceCadence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_cadence", values[i])
			} else if value.Valid {
				pl.InvoiceCadence = value.String
			}
		case plan.FieldTrialPeriod:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trial_period", values[i])
			} else if value.Valid {
				pl.TrialPeriod = int(value.Int64)
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Plan.
// This includes values selected through modifiers, order, etc.
func (pl *Plan) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// Update returns a builder for updating this Plan.
// Note that you need to call Plan.Unwrap() before calling this method if this Plan
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Plan) Update() *PlanUpdateOne {
	return NewPlanClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Plan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Plan) Unwrap() *Plan {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Plan is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Plan) String() string {
	var builder strings.Builder
	builder.WriteString("Plan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(pl.TenantID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(pl.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pl.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pl.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("lookup_key=")
	builder.WriteString(pl.LookupKey)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pl.Description)
	builder.WriteString(", ")
	builder.WriteString("invoice_cadence=")
	builder.WriteString(pl.InvoiceCadence)
	builder.WriteString(", ")
	builder.WriteString("trial_period=")
	builder.WriteString(fmt.Sprintf("%v", pl.TrialPeriod))
	builder.WriteByte(')')
	return builder.String()
}

// Plans is a parsable slice of Plan.
type Plans []*Plan