package schema

import (
	"reflect"
)

// Table is a description of a SQL table.
//
// Type is the type of the object that generated this table.
//
// Name is the name of the SQL table.
//
// Columns is the collection of columns in the SQL table mapped by name.
type Table struct {
	Type      reflect.Type
	Name      string
	Columns   map[string]*Column
	Relations []Relation
}

// NewTable constructs a new Table with the given name and an
// empty column map.
func NewTable(dataType reflect.Type, name string) *Table {
	return &Table{
		Type:      dataType,
		Name:      name,
		Columns:   make(map[string]*Column),
		Relations: make([]Relation, 0),
	}
}

// TableEntity provides externaly the requirements on a type to be used as a
// table entity.
type TableEntity interface {
	namedTable
}

// The type namedTable provides the machinery for given table entities a custom
// table name.
//
// TableName returns the custom table name to use.
type namedTable interface {
	TableName() string
}

func getTableName(object interface{}) string {
	var name string
	if val, ok := object.(namedTable); ok {
		name = val.TableName()
	} else {
		name = tableNameFromObject(object)
	}
	return name
}

func tableNameFromObject(object interface{}) string {
	return tableNameFromType(reflect.TypeOf(object))
}

// Extract a SQL style table name from the type of an object.
func tableNameFromType(objectType reflect.Type) string {
	var typeName = objectType.String()
	return sqlNameFromCamelCase(typeName) + "s"
}
