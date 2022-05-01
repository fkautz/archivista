// Code generated by entc, DO NOT EDIT.

package statement

const (
	// Label holds the string label denoting the statement type in the database.
	Label = "statement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatement holds the string denoting the statement field in the database.
	FieldStatement = "statement"
	// EdgeSubjects holds the string denoting the subjects edge name in mutations.
	EdgeSubjects = "subjects"
	// Table holds the table name of the statement in the database.
	Table = "statements"
	// SubjectsTable is the table that holds the subjects relation/edge. The primary key declared below.
	SubjectsTable = "statement_subjects"
	// SubjectsInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectsInverseTable = "subjects"
)

// Columns holds all SQL columns for statement fields.
var Columns = []string{
	FieldID,
	FieldStatement,
}

var (
	// SubjectsPrimaryKey and SubjectsColumn2 are the table columns denoting the
	// primary key for the subjects relation (M2M).
	SubjectsPrimaryKey = []string{"statement_id", "subject_id"}
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
