package schema

import (
	"reflect"
)

type Schema struct {
	Name   string
	Tables map[reflect.Type]*Table
}

// NewSchema constructs a new Schema with the given name and empty
// Table map.
func NewSchema(name string) *Schema {
	return &Schema{
		Name:   name,
		Tables: make(map[reflect.Type]*Table)}
}
