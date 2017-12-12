package icebox

// ConstraintType is the type of a constraint on a column.
// There are 7 constraint types to be supported by icebox.
// NotNull - Ensures that a column cannot have a NULL value
// Unique - Ensures that all values in a column are different
// PrimaryKey - A combination of a NOT NULL and UNIQUE. Uniquely identifies each row in a table
// ForeignKey - Uniquely identifies a row/record in another table
// Check - Ensures that all values in a column satisfies a specific condition
// Default - Sets a default value for a column when no value is specified
// Index - Used to create and retrieve data from the database very quickly
type ConstraintType uint

const (
	NotNull ConstraintType = iota
	Unique
	PrimaryKey
	ForeignKey
	Check
	Default
	Index
)

func (c ConstraintType) String() string {
	if uint(c) < len(constraintTypeNames) {
		return constraintTypeNames[c]
	}
}

var constraintTypeNames = []string{
	NotNull:    "notNull",
	Unique:     "unique",
	PrimaryKey: "primaryKey",
	ForeignKey: "foreignKey",
	Check:      "check",
	Default:    "default",
	Index:      "index",
}

type Constraint struct {
	Type ConstraintType
	Info string
}
