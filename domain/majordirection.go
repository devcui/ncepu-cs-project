// Code generated by ent, DO NOT EDIT.

package domain

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/devcui/ncepu-cs-project/domain/class"
	"github.com/devcui/ncepu-cs-project/domain/majordirection"
)

// MajorDirection is the model entity for the MajorDirection schema.
type MajorDirection struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 专业方向名称
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MajorDirectionQuery when eager-loading is set.
	Edges MajorDirectionEdges `json:"edges"`
}

// MajorDirectionEdges holds the relations/edges for other nodes in the graph.
type MajorDirectionEdges struct {
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MajorDirectionEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[0] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MajorDirection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case majordirection.FieldID:
			values[i] = new(sql.NullInt64)
		case majordirection.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MajorDirection", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MajorDirection fields.
func (md *MajorDirection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case majordirection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			md.ID = int(value.Int64)
		case majordirection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				md.Name = value.String
			}
		}
	}
	return nil
}

// QueryClass queries the "class" edge of the MajorDirection entity.
func (md *MajorDirection) QueryClass() *ClassQuery {
	return NewMajorDirectionClient(md.config).QueryClass(md)
}

// Update returns a builder for updating this MajorDirection.
// Note that you need to call MajorDirection.Unwrap() before calling this method if this MajorDirection
// was returned from a transaction, and the transaction was committed or rolled back.
func (md *MajorDirection) Update() *MajorDirectionUpdateOne {
	return NewMajorDirectionClient(md.config).UpdateOne(md)
}

// Unwrap unwraps the MajorDirection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (md *MajorDirection) Unwrap() *MajorDirection {
	_tx, ok := md.config.driver.(*txDriver)
	if !ok {
		panic("domain: MajorDirection is not a transactional entity")
	}
	md.config.driver = _tx.drv
	return md
}

// String implements the fmt.Stringer.
func (md *MajorDirection) String() string {
	var builder strings.Builder
	builder.WriteString("MajorDirection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", md.ID))
	builder.WriteString("name=")
	builder.WriteString(md.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MajorDirections is a parsable slice of MajorDirection.
type MajorDirections []*MajorDirection
