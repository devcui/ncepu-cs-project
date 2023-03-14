// Code generated by ent, DO NOT EDIT.

package department

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeMajor holds the string denoting the major edge name in mutations.
	EdgeMajor = "major"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// Table holds the table name of the department in the database.
	Table = "departments"
	// MajorTable is the table that holds the major relation/edge.
	MajorTable = "majors"
	// MajorInverseTable is the table name for the Major entity.
	// It exists in this package in order to avoid circular dependency with the "major" package.
	MajorInverseTable = "majors"
	// MajorColumn is the table column denoting the major relation/edge.
	MajorColumn = "major_department"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "classes"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_department"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "students"
	// StudentInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentInverseTable = "students"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "student_department"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCode,
	FieldDescription,
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