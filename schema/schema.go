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
)

// Schema is a representation of a SQL database schema. Such a schema is determined
// by its name, and the collection of table that it holds.
//
// In icebox, tables are generated to correspond with application objects, and so
// Schemas in icebox provide a 1-to-1 interface between applications objects and the
// tables that represent them in the database.
//
// Name returns the name of this schema.
//
// TableFor returns the Table corresponding to the type of the given object. If there is
// no table for the given object, TableFor returns an error.
type Schema interface {
	Name() string
	TableFor(interface{}) (Table, error)
}

// The default implementation of the Schema interface.
//
// Name is the name of this schema.
//
// Tables is a map from reflect.Type, that is the type of a given object, to its
// corresponding table.
type schemaImpl struct {
	name   string
	tables map[reflect.Type]*tableImpl
}

// Returns the internal name of the schema.
func (s *schemaImpl) Name() string {
	return s.name
}

// Returns the Table in the schema, corresponding to the given object.
// This returns an error if the given object can't be found.
func (s *schemaImpl) TableFor(object interface{}) (Table, error) {
	objectType := getConcreteObjectType(reflect.TypeOf(object))
	table, found := s.tables[objectType]
	if !found {
		return nil, &notFoundError{
			key: objectType,
			msg: "no table for this object",
		}
	}
	return table, nil
}

// Constructs a new Schema of the default implementation with the given name
// and empty table map.
func newSchema(name string) *schemaImpl {
	return &schemaImpl{
		name:   name,
		tables: make(map[reflect.Type]*tableImpl)}
}
