// Code generated by ent, DO NOT EDIT.

package domain

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/devcui/ncepu-cs-project/domain/role"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoleName holds the value of the "role_name" field.
	RoleName string `json:"role_name,omitempty"`
	// RoleValue holds the value of the "role_value" field.
	RoleValue string `json:"role_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Resource holds the value of the resource edge.
	Resource []*Resource `json:"resource,omitempty"`
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) ResourceOrErr() ([]*Resource, error) {
	if e.loadedTypes[0] {
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			values[i] = new(sql.NullInt64)
		case role.FieldRoleName, role.FieldRoleValue:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case role.FieldRoleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_name", values[i])
			} else if value.Valid {
				r.RoleName = value.String
			}
		case role.FieldRoleValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_value", values[i])
			} else if value.Valid {
				r.RoleValue = value.String
			}
		}
	}
	return nil
}

// QueryResource queries the "resource" edge of the Role entity.
func (r *Role) QueryResource() *ResourceQuery {
	return NewRoleClient(r.config).QueryResource(r)
}

// QueryUser queries the "user" edge of the Role entity.
func (r *Role) QueryUser() *UserQuery {
	return NewRoleClient(r.config).QueryUser(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("domain: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("role_name=")
	builder.WriteString(r.RoleName)
	builder.WriteString(", ")
	builder.WriteString("role_value=")
	builder.WriteString(r.RoleValue)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
