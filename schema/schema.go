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

type Schema interface {
	Name() string
	TableFor(interface{}) (Table, error)
}

type schemaImpl struct {
	name   string
	tables map[reflect.Type]*tableImpl
}

func (s *schemaImpl) Name() string {
	return s.name
}

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

// NewSchema constructs a new Schema with the given name and empty
// Table map.
func newSchema(name string) *schemaImpl {
	return &schemaImpl{
		name:   name,
		tables: make(map[reflect.Type]*tableImpl)}
}
