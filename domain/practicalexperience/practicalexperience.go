// Code generated by ent, DO NOT EDIT.

package practicalexperience

const (
	// Label holds the string label denoting the practicalexperience type in the database.
	Label = "practical_experience"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// Table holds the table name of the practicalexperience in the database.
	Table = "practical_experiences"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "practical_experiences"
	// StudentInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentInverseTable = "students"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "practical_experience_student"
)

// Columns holds all SQL columns for practicalexperience fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "practical_experiences"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"practical_experience_student",
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