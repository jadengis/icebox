package schema

import (
	"github.com/jadengis/icebox/logger"
	"github.com/jadengis/icebox/tags"
	"github.com/jadengis/icebox/types"
	"reflect"
)

// generateSchema will construct a database schema given a name for the database
// and a list of objects.
func generateSchema(name string, objects ...interface{}) (*Schema, error) {
	// Iterate through the all the objects, and build tables for each.
	schema := NewSchema(name)
	for i := 0; i < len(objects); i++ {
		table, err := generateTable(objects[i])
		if err != nil {
			return nil, &schemaGenError{
				cause: err,
				msg:   "error generating table"}
		}
		schema.Tables[table.Type] = table
	}
	return schema, nil
}

func generateTable(object interface{}) (*Table, error) {
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
	table := NewTable(objectType, name)
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
				table.Columns[column.Name] = column
			}
		}
	}
	return table, nil
}

func getConcreteObjectType(objectType reflect.Type) reflect.Type {
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
	}
	return objectType
}

func handleColumnTag(field reflect.StructField, parsedTag tags.ParsedTag) *Column {
	if info, found := parsedTag.GetInfo(tags.Column); found {
		delete(parsedTag, tags.Column)
		if len(info) == 0 {
			info = sqlNameFromCamelCase(field.Name)
		}
		sqlType, err := mapSQLTypeFromField(field)
		if err != nil {
			return nil
		}
		return NewColumn(info, sqlType)
	}
	return nil
}

func mapSQLTypeFromField(field reflect.StructField) (*types.SQLType, error) {
	switch kind := getConcreteObjectType(field.Type).Kind(); kind {
	case reflect.Int8:
		return types.NewSQLType(types.TinyInt, nil), nil
	default:
		return nil, &unknownTypeError{
			typeName: kind.String(),
			msg:      "unsupported go type",
		}
	}
}

// Construct a slice of Constraints from the given parsed tag.
func handleConstraintTags(parsedTag tags.ParsedTag) []*Constraint {
	var constraints []*Constraint
	for tag, info := range parsedTag {
		constraintType, err := getConstraintType(tag)
		if err != nil {
			// this is not a contraint tag so skip it.
			continue
		}
		constraint := NewConstraint(constraintType, info)
		delete(parsedTag, tag)
		constraints = append(constraints, constraint)
	}
	return constraints
}
