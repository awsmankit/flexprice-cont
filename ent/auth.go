// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/auth"
)

// Auth is the model entity for the Auth schema.
type Auth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Auth) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			values[i] = new(sql.NullInt64)
		case auth.FieldUserID, auth.FieldProvider, auth.FieldToken, auth.FieldStatus:
			values[i] = new(sql.NullString)
		case auth.FieldCreatedAt, auth.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Auth fields.
func (a *Auth) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case auth.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = value.String
			}
		case auth.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				a.Provider = value.String
			}
		case auth.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				a.Token = value.String
			}
		case auth.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = value.String
			}
		case auth.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case auth.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Auth.
// This includes values selected through modifiers, order, etc.
func (a *Auth) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Auth.
// Note that you need to call Auth.Unwrap() before calling this method if this Auth
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Auth) Update() *AuthUpdateOne {
	return NewAuthClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Auth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Auth) Unwrap() *Auth {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Auth is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Auth) String() string {
	var builder strings.Builder
	builder.WriteString("Auth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("user_id=")
	builder.WriteString(a.UserID)
	builder.WriteString(", ")
	builder.WriteString("provider=")
	builder.WriteString(a.Provider)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(a.Token)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(a.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Auths is a parsable slice of Auth.
type Auths []*Auth