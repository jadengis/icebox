package schema

type Relation struct {
	relatedTable *Table
}

type Schema struct {
	Name   string
	Tables map[string]Table
}

// Func NewSchema constructs a new Schema with the given name and empty
// Table map.
func NewSchema(name string) *Schema {
	return &Schema{
		Name:   name,
		Tables: make(map[string]Table)}
}
