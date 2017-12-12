package schema

// Table is a description of a SQL table.
type Table struct {
	// Type is the type of the object that generated this table.
	Type reflect.Type
	// Name is the name of the SQL table.
	Name string
	// Columns is the collection of columns in the SQL table ampped by name.
	Columns map[string]Column
}

// Func NewTable constructs a new Table with the given name and an
// empty column map.
func NewTable(dataType reflect.Type, name string) *Table {
	return &Table{
		Type:    dataType,
		Name:    name,
		Columns: make(map[string]Column)}
}
