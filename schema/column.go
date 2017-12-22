package schema

import (
	"github.com/jadengis/icebox/types"
)

// Column is a description of a column in a SQL table.
// A column in a SQL table can be specified by its name, a datatype, and
// a list of constraints on elements of the column.
//
// Name is the name of the column in the schema.
//
// Type is the SQLType of the column in the schema.
//
// Constraints is the list of contraints on the column, such as
// PRIMARY KEY or NOT NULL.
type Column struct {
	Name        string
	Type        *types.SQLType
	Constraints map[ConstraintType]*Constraint
}

// NewColumn constructs a new column object with the given name and type.
func NewColumn(name string, sqlType *types.SQLType) *Column {
	return &Column{
		Name:        name,
		Type:        sqlType,
		Constraints: make(map[ConstraintType]*Constraint),
	}
}

// Add a slice of Constraints to the Column.
func (c *Column) bulkAddConstraints(constraints []*Constraint) {
	for _, constraint := range constraints {
		c.Constraints[constraint.Type] = constraint
	}
}
