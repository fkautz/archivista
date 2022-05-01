// Code generated by entc, DO NOT EDIT.

package subject

const (
	// Label holds the string label denoting the subject type in the database.
	Label = "subject"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeDigests holds the string denoting the digests edge name in mutations.
	EdgeDigests = "digests"
	// EdgeStatement holds the string denoting the statement edge name in mutations.
	EdgeStatement = "statement"
	// Table holds the table name of the subject in the database.
	Table = "subjects"
	// DigestsTable is the table that holds the digests relation/edge.
	DigestsTable = "digests"
	// DigestsInverseTable is the table name for the Digest entity.
	// It exists in this package in order to avoid circular dependency with the "digest" package.
	DigestsInverseTable = "digests"
	// DigestsColumn is the table column denoting the digests relation/edge.
	DigestsColumn = "subject_digests"
	// StatementTable is the table that holds the statement relation/edge. The primary key declared below.
	StatementTable = "statement_subjects"
	// StatementInverseTable is the table name for the Statement entity.
	// It exists in this package in order to avoid circular dependency with the "statement" package.
	StatementInverseTable = "statements"
)

// Columns holds all SQL columns for subject fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// StatementPrimaryKey and StatementColumn2 are the table columns denoting the
	// primary key for the statement relation (M2M).
	StatementPrimaryKey = []string{"statement_id", "subject_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
