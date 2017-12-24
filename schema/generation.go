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
	"github.com/jadengis/icebox/logger"
	"github.com/jadengis/icebox/tags"
	"github.com/jadengis/icebox/types"
	"reflect"
)

// NewSchema will construct a database schema given a name for the database
// and a list of objects that will comprise this schema.
//
// If there are an errors during schema generation, this function will return
// an error.
func NewSchema(name string, objects ...interface{}) (Schema, error) {
	// Iterate through the all the objects, and build tables for each.
	schema := newSchema(name)
	for i := 0; i < len(objects); i++ {
		table, err := generateTable(objects[i])
		if err != nil {
			return nil, &schemaGenError{
				cause: err,
				msg:   "error generating table"}
		}
		schema.tables[table.dataType] = table
	}
	return schema, nil
}

// Construct and populate the database table corresponding to the given object.
// This function use the default implementation of the Table interface.
//
// If table generation fails, this method will return an error.
func generateTable(object interface{}) (*tableImpl, error) {
	objectType := reflect.TypeOf(object)
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
	}
	if objectType.Kind() != reflect.Struct {
		return nil, &typeError{
			badType: objectType,
			msg:     "only structs and ptr to struct are supported types"}
	}

	// Build a Table for objectType by iterating through its struct fields,
	// parsing out tags, and building the appropraite columns with the appropriate
	// constraints.
	name := getTableName(object)
	table := newTable(objectType, name)
	for i := 0; i < objectType.NumField(); i++ {
		field := objectType.Field(i)

		// Check the field for an Icebox tag, and parse subtags if needed.
		if tag, ok := field.Tag.Lookup(tags.Icebox.String()); ok {
			parsedTag, err := tags.Parse(tag)
			if err != nil {
				logger.Error("could not parse tags", err)
			}
			column := handleColumnTag(field, parsedTag)
			if column != nil {
				constraints := handleConstraintTags(parsedTag)
				column.bulkAddConstraints(constraints)
				table.columns[column.name] = column
			}
		}
	}
	return table, nil
}

// Get the concrete type of a reflect.Type, that is, resolve what the given type
// points to.
func getConcreteObjectType(objectType reflect.Type) reflect.Type {
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
	}
	return objectType
}

// Process a column tag on struct, and return a corresponding column.
// This function uses the default Column implementation.
func handleColumnTag(field reflect.StructField, parsedTag tags.ParsedTag) *columnImpl {
	if info, found := parsedTag.GetInfo(tags.Column); found {
		delete(parsedTag, tags.Column)
		if len(info) == 0 {
			info = sqlNameFromCamelCase(field.Name)
		}
		sqlType, err := mapSQLTypeFromField(field)
		if err != nil {
			return nil
		}
		return newColumn(info, sqlType)
	}
	return nil
}

// Map the type of the given struct field to its corresponding SQLType.
// This returns an error if the struct field type is not supported.
func mapSQLTypeFromField(field reflect.StructField) (types.SQLType, error) {
	switch kind := getConcreteObjectType(field.Type).Kind(); kind {
	case reflect.Bool:
		return types.NewSQLType(types.Bit), nil
	case reflect.Int8:
		return types.NewSQLType(types.TinyInt), nil
	case reflect.Int16:
		return types.NewSQLType(types.SmallInt), nil
	case reflect.Int32:
		return types.NewSQLType(types.MediumInt), nil
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		return types.NewSQLType(types.Int), nil
	case reflect.Uint8:
		return types.NewSQLType(types.TinyUint), nil
	case reflect.Uint16:
		return types.NewSQLType(types.SmallUint), nil
	case reflect.Uint32:
		return types.NewSQLType(types.MediumUint), nil
	case reflect.Uint64:
		fallthrough
	case reflect.Uint:
		return types.NewSQLType(types.Uint), nil
	case reflect.Float32:
		return types.NewSQLType(types.Float), nil
	case reflect.Float64:
		return types.NewSQLType(types.Double), nil
	case reflect.String:
		return types.NewSQLTypeWithSize(types.VarChar, "255"), nil
	default:
		return nil, &unknownTypeError{
			typeName: kind.String(),
			msg:      "unsupported go type",
		}
	}
}

// Construct a slice of constraints from the given parsed tag.
func handleConstraintTags(parsedTag tags.ParsedTag) []*constraintImpl {
	var constraints []*constraintImpl
	for tag, info := range parsedTag {
		constraintType, err := getConstraintType(tag)
		if err != nil {
			// this is not a contraint tag so skip it.
			continue
		}
		constraint := newConstraint(constraintType, info)
		delete(parsedTag, tag)
		constraints = append(constraints, constraint)
	}
	return constraints
}
