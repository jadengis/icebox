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
	"reflect"
	"strings"
)

// Table is a description of a SQL table as an interface.
//
// Type returns the type of the object that generated this table.
//
// Name returns the name of the SQL table.
//
// Columns returns the slice of all columns in this table.
//
// ColumnFor returns the Table column for the given column name.
//
// Relations returns the slice of Relations on this table.
//
// RelationFor returns the relation on the table for the given relation type.
type Table interface {
	Type() reflect.Type
	Name() string
	Columns() []Column
	ColumnFor(string) (Column, error)
	Relations() []Relation
	RelationFor(RelationType) (Relation, bool)
}

// The default implementation of the Table interface.
//
// Type is the type of the object that generated this table.
//
// Name is the name of the SQL table.
//
// Columns is the collection of columns in the SQL table mapped by name.
//
// Relations is the collection of relations on the tabel mapped by type.
type tableImpl struct {
	dataType  reflect.Type
	name      string
	columns   map[string]*columnImpl
	relations map[RelationType]*relationImpl
}

// Returns the reflect.Type this table corresponds to.
func (t *tableImpl) Type() reflect.Type {
	return t.dataType
}

// Returns the name of this table.
func (t *tableImpl) Name() string {
	return t.name
}

// Return the slice of columns in this table by pulling them from the column map.
func (t *tableImpl) Columns() []Column {
	columns := make([]Column, len(t.columns))
	for _, column := range t.columns {
		columns = append(columns, column)
	}
	return columns
}

// Looks-up the given column name in the table, and returns the corresponding
// Column, if it exists.
func (t *tableImpl) ColumnFor(name string) (Column, error) {
	column, found := t.columns[name]
	if !found {
		return nil, &notFoundError{
			key: name,
			msg: "no column with the given name",
		}
	}
	return column, nil
}

// Returns the slice of relations on this table by pulling them from the relation map.
func (t *tableImpl) Relations() []Relation {
	relations := make([]Relation, len(t.relations))
	for _, relation := range t.relations {
		relations = append(relations, relation)
	}
	return relations
}

// Looks-up the given column name in the table, and returns the corresponding
// Column, if it exists.
func (t *tableImpl) RelationFor(relationType RelationType) (Relation, bool) {
	relation, found := t.relations[relationType]
	return relation, found
}

// Constructs a new table of the default implementation with the given name and an
// empty column map and empty relation map.
func newTable(dataType reflect.Type, name string) *tableImpl {
	return &tableImpl{
		dataType:  dataType,
		name:      name,
		columns:   make(map[string]*columnImpl),
		relations: make(map[RelationType]*relationImpl),
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

// Get the table name for the given object.
// This table name can either be specified (via the TableEntity interface)
// or automatically generated from the types name.
func getTableName(object interface{}) string {
	if val, ok := object.(namedTable); ok {
		return val.TableName()
	}
	return tableNameFromObject(object)
}

// Extract a SQL style table name for an object.
func tableNameFromObject(object interface{}) string {
	return tableNameFromType(reflect.TypeOf(object))
}

// Extract a SQL style table name for an object given its type.
func tableNameFromType(objectType reflect.Type) string {
	var typeName = objectType.String()
	if idx := strings.Index(typeName, "."); idx != -1 {
		typeName = typeName[idx+1:]
	}
	return sqlNameFromCamelCase(typeName) + "s"
}
