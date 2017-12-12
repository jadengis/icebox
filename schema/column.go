package schema

// Column is a description of a column in a SQL table.
// A column in a SQL table can be specified by its name, a datatype, and
// a list of constraints on elements of the column.
type Column struct {
	// Name is the name of the column in the schema.
	Name string
	// Type is the type of the column in the schema.
	Type string
	// Constraints is the list of contraints on the column, such as
	// PRIMARY KEY or NOT NULL.
	Constraints []Constraint
}

// Func NewColumn constructs a new column object with the given name and type.
func NewColumn(name string, sqlType string) *Column {
	return &Column{
		Name: name,
		Type: sqlType}
}
