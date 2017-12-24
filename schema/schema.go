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
