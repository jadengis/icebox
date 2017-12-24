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
	"github.com/jadengis/icebox/types"
)

// Column is a description of a column in a SQL table.
// A column in a SQL table can be specified by its name, a datatype, and
// a list of constraints on elements of the column. This interface is an
// abstraction of this concept into its key components
//
// Name retruns the name of the column in the schema.
//
// Type returns the SQLType of the column in the schema.
//
// Constraints returns the list of contraints on the column, such as
// PRIMARY KEY or NOT NULL.
//
// ConstraintFor returns the constraint on the column for the given
// constraint type, and whether or not it exists.
type Column interface {
	Name() string
	Type() types.SQLType
	Constraints() []Constraint
	ConstraintFor(ConstraintType) (Constraint, bool)
}

// The default implementation of the Column interface.
//
// Name is the name of the column in the schema.
//
// Type is the SQLType of the column in the schema.
//
// Constraints is a map of all contraints on the column, such as
// PRIMARY KEY or NOT NULL, key off by type.
type columnImpl struct {
	name        string
	sqlType     types.SQLType
	constraints map[ConstraintType]*constraintImpl
}

// Return the internal name of the column.
func (c *columnImpl) Name() string {
	return c.name
}

func (c *columnImpl) Type() types.SQLType {
	return c.sqlType
}

// Look up the given constraint type in the column, and return it if it exists, along
// with a bool indicating its existence.
func (c *columnImpl) ConstraintFor(constraintType ConstraintType) (Constraint, bool) {
	constraint, found := c.constraints[constraintType]
	return constraint, found
}

// Return a list of all constraints in the constraints map for this column.
func (c *columnImpl) Constraints() []Constraint {
	constraints := make([]Constraint, len(c.constraints))
	for _, constraint := range c.constraints {
		constraints = append(constraints, constraint)
	}
	return constraints
}

// Construct a new columnImpl object with the given name and SQLType,
// and return a pointer to it.
func newColumn(name string, sqlType types.SQLType) *columnImpl {
	return &columnImpl{
		name:        name,
		sqlType:     sqlType,
		constraints: make(map[ConstraintType]*constraintImpl),
	}
}

// Add a slice of constraints to the column.
func (c *columnImpl) bulkAddConstraints(constraints []*constraintImpl) {
	for _, constraint := range constraints {
		c.constraints[constraint.constraintType] = constraint
	}
}
