// Code generated by ent, DO NOT EDIT.

package resource

const (
	// Label holds the string label denoting the resource type in the database.
	Label = "resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldResourceName holds the string denoting the resource_name field in the database.
	FieldResourceName = "resource_name"
	// FieldResourceValue holds the string denoting the resource_value field in the database.
	FieldResourceValue = "resource_value"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeAuthorization holds the string denoting the authorization edge name in mutations.
	EdgeAuthorization = "authorization"
	// Table holds the table name of the resource in the database.
	Table = "resources"
	// RoleTable is the table that holds the role relation/edge. The primary key declared below.
	RoleTable = "role_resource"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_resource"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// AuthorizationTable is the table that holds the authorization relation/edge. The primary key declared below.
	AuthorizationTable = "authorization_resource"
	// AuthorizationInverseTable is the table name for the Authorization entity.
	// It exists in this package in order to avoid circular dependency with the "authorization" package.
	AuthorizationInverseTable = "authorizations"
)

// Columns holds all SQL columns for resource fields.
var Columns = []string{
	FieldID,
	FieldResourceName,
	FieldResourceValue,
}

var (
	// RolePrimaryKey and RoleColumn2 are the table columns denoting the
	// primary key for the role relation (M2M).
	RolePrimaryKey = []string{"role_id", "resource_id"}
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "resource_id"}
	// AuthorizationPrimaryKey and AuthorizationColumn2 are the table columns denoting the
	// primary key for the authorization relation (M2M).
	AuthorizationPrimaryKey = []string{"authorization_id", "resource_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ResourceNameValidator is a validator for the "resource_name" field. It is called by the builders before save.
	ResourceNameValidator func(string) error
	// ResourceValueValidator is a validator for the "resource_value" field. It is called by the builders before save.
	ResourceValueValidator func(string) error
)
