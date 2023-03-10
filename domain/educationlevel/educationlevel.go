// Code generated by ent, DO NOT EDIT.

package educationlevel

const (
	// Label holds the string label denoting the educationlevel type in the database.
	Label = "education_level"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// Table holds the table name of the educationlevel in the database.
	Table = "education_levels"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "education_levels"
	// StudentInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentInverseTable = "students"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "education_level_student"
)

// Columns holds all SQL columns for educationlevel fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "education_levels"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"education_level_student",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
