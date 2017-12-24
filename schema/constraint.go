// Copyright 2017 John Dengis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"github.com/jadengis/icebox/tags"
	"strconv"
)

// Constraint is a representation of a constraint on a SQL column.
// These can be things such as PRIMARY KEY, NOT NULL, or UNIQUE.
// A constraint can be fully described by its constraint type, along with
// some optional details to describe that type.
//
// Type returns the constraint type of this constraint.
//
// Details returns the specific details of this contraint. For example the
// DEFAULT should have some default value specified in the details.
type Constraint interface {
	Type() ConstraintType
	Details() string
}

// ConstraintType is the type of a constraint on a column.
//
// There are 7 constraint types to be supported by icebox.
//
// NotNull:    Ensures that a column cannot have a NULL value
//
// Unique:     Ensures that all values in a column are different
//
// PrimaryKey: A combination of a NOT NULL and UNIQUE. Uniquely identifies each
// row in a table
//
// ForeignKey: Uniquely identifies a row/record in another table
//
// Check:      Ensures that all values in a column satisfies a specific condition
//
// Default:    Sets a default value for a column when no value is specified
//
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

// The default implementation of the Constraint interface.
//
// ConstraintType is the ConstraintType of this Constraint.
//
// Details is additional details about this constraint in the form of a string.
type constraintImpl struct {
	constraintType ConstraintType
	details        string
}

// Returns the internal constraint type of this constraint.
func (c *constraintImpl) Type() ConstraintType {
	return c.constraintType
}

// Returns the internal details of the constraint.
func (c *constraintImpl) Details() string {
	return c.details
}

// Constructs a new constraint of the default implementation,
// with the given type and details, and returns a pointer to this object.
func newConstraint(constraintType ConstraintType, details string) *constraintImpl {
	return &constraintImpl{
		constraintType: constraintType,
		details:        details,
	}
}
