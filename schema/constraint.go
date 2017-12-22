package schema

import (
	"github.com/jadengis/icebox/tags"
	"strconv"
)

// ConstraintType is the type of a constraint on a column.
// There are 7 constraint types to be supported by icebox.
// NotNull:    Ensures that a column cannot have a NULL value
// Unique:     Ensures that all values in a column are different
// PrimaryKey: A combination of a NOT NULL and UNIQUE. Uniquely identifies each
// row in a table
// ForeignKey: Uniquely identifies a row/record in another table
// Check:      Ensures that all values in a column satisfies a specific condition
// Default:    Sets a default value for a column when no value is specified
// Index:      Used to create and retrieve data from the database very quickly
type ConstraintType int

const (
	InvalidConstraint ConstraintType = -1
	NotNull           ConstraintType = iota
	Unique
	PrimaryKey
	ForeignKey
	Check
	Default
	Index
)

// String converts the given ConstraintType into its string respresentation.
func (c ConstraintType) String() string {
	if int(c) < len(constraintTypeTags) {
		return constraintTypeTags[c].String()
	}
	return "constraintType" + strconv.Itoa(int(c))
}

// Mapping between ConstraintType and tag names.
var constraintTypeTags = []tags.SubTag{
	NotNull:    tags.NotNull,
	Unique:     tags.Unique,
	PrimaryKey: tags.PrimaryKey,
	ForeignKey: tags.ForeignKey,
	Check:      tags.Check,
	Default:    tags.Default,
	Index:      tags.Index,
}

// Map from constraint type names to corresponding ConstraintType.
// Will return an unknownTypeError if the name cannot be resolved.
func getConstraintType(typeName tags.SubTag) (ConstraintType, error) {
	switch typeName {
	case tags.NotNull:
		return NotNull, nil
	case tags.Unique:
		return Unique, nil
	case tags.PrimaryKey:
		return PrimaryKey, nil
	case tags.ForeignKey:
		return ForeignKey, nil
	case tags.Check:
		return Check, nil
	case tags.Default:
		return Default, nil
	case tags.Index:
		return Index, nil
	default:
		return InvalidConstraint, &unknownTypeError{
			typeName: typeName.String(),
			msg:      "typeName could not be resolved"}
	}
}

// Constraint is a description of a SQL constraint on a column.
// Type is the ConstraintType of this Constraint.
// Info is additional info about this constaint in the form of a string.
type Constraint struct {
	Type ConstraintType
	Info string
}

// NewConstraint constructs a new constraint with the given type and info, and
// returns a pointer to this object.
func NewConstraint(contraintType ConstraintType, info string) *Constraint {
	return &Constraint{
		Type: contraintType,
		Info: info,
	}
}
