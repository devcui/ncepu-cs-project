// Code generated by ent, DO NOT EDIT.

package domain

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/devcui/ncepu-cs-project/domain/department"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 系部名称
	Name string `json:"name,omitempty"`
	// 系部代码
	Code string `json:"code,omitempty"`
	// 系部描述
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges DepartmentEdges `json:"edges"`
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Major holds the value of the major edge.
	Major []*Major `json:"major,omitempty"`
	// Class holds the value of the class edge.
	Class []*Class `json:"class,omitempty"`
	// Student holds the value of the student edge.
	Student []*Student `json:"student,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MajorOrErr returns the Major value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) MajorOrErr() ([]*Major, error) {
	if e.loadedTypes[0] {
		return e.Major, nil
	}
	return nil, &NotLoadedError{edge: "major"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) ClassOrErr() ([]*Class, error) {
	if e.loadedTypes[1] {
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) StudentOrErr() ([]*Student, error) {
	if e.loadedTypes[2] {
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			values[i] = new(sql.NullInt64)
		case department.FieldName, department.FieldCode, department.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Department", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case department.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case department.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				d.Code = value.String
			}
		case department.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		}
	}
	return nil
}

// QueryMajor queries the "major" edge of the Department entity.
func (d *Department) QueryMajor() *MajorQuery {
	return NewDepartmentClient(d.config).QueryMajor(d)
}

// QueryClass queries the "class" edge of the Department entity.
func (d *Department) QueryClass() *ClassQuery {
	return NewDepartmentClient(d.config).QueryClass(d)
}

// QueryStudent queries the "student" edge of the Department entity.
func (d *Department) QueryStudent() *StudentQuery {
	return NewDepartmentClient(d.config).QueryStudent(d)
}

// Update returns a builder for updating this Department.
// Note that you need to call Department.Unwrap() before calling this method if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return NewDepartmentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Department entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("domain: Department is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(d.Code)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department