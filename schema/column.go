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
	ConstraintFor(ConstraintType) (Constraint, bool)
	Constraints() []Constraint
}

// The default implementation of the Column interface.
//
// Name is the name of the column in the schema.
//
// Type is the SQLType of the column in the schema.
//
// Constraints is the list of contraints on the column, such as
// PRIMARY KEY or NOT NULL.
type columnImpl struct {
	name        string
	sqlType     types.SQLType
	constraints map[ConstraintType]*constraintImpl
}

func (c *columnImpl) Name() string {
	return c.name
}

func (c *columnImpl) Type() types.SQLType {
	return c.sqlType
}

func (c *columnImpl) ConstraintFor(constraintType ConstraintType) (Constraint, bool) {
	constraint, found := c.constraints[constraintType]
	return constraint, found
}

func (c *columnImpl) Constraints() []Constraint {
	constraints := make([]Constraint, len(c.constraints))
	for _, constraint := range c.constraints {
		constraints = append(constraints, constraint)
	}
	return constraints
}

// NewColumn constructs a new column object with the given name and type.
func newColumn(name string, sqlType types.SQLType) *columnImpl {
	return &columnImpl{
		name:        name,
		sqlType:     sqlType,
		constraints: make(map[ConstraintType]*constraintImpl),
	}
}

// Add a slice of Constraints to the Column.
func (c *columnImpl) bulkAddConstraints(constraints []*constraintImpl) {
	for _, constraint := range constraints {
		c.constraints[constraint.constraintType] = constraint
	}
}
