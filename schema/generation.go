package schema

import (
	"fmt"
	"github.com/jadengis/icebox/tags"
	"reflect"
)

// schemaGenError is a general wrapper for schema generation errors.
type schemaGenError struct {
	cause error
	msg   string
}

// Error message for a schema generation error.
func (e *schemaGenError) Error() string {
	return fmt.Sprintf("%s - %s", e.msg, e.cause.Error())
}

// Schema generation error relating to a bad type.
type typeError struct {
	badType reflect.Type
	msg     string
}

// Error message for this type error.
func (e *typeError) Error() string {
	return fmt.Sprintf("unsupported type %s - %s", e.badType, e.msg)
}

// generateSchema will construct a database schema given a name for the database
// and a list of objects.
func generateSchema(name string, objects ...interface{}) (*Schema, error) {
	// Iterate through the all the objects, and build tables for each.
	schema := NewSchema(name)
	for i := 0; i < len(objects); i++ {
		table, err := generateTable(object[i])
		if err != nil {
			return nil, &schemaGenError{
				cause: err,
				msg:   "error generating table"}
		}
		Tables[table.Name] = table
	}
	return schema, nil
}

func generateTable(object interface{}) (*Table, error) {
	objectType := reflect.TypeOf(object)
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
	}
	if objectType != reflect.Struct {
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
		tags := field.Tag()

		// Check the field for an Icebox tag, and parse subtags if needed.
		if tag, ok := tag.Lookup(tags.Icebox); ok {
			parsedTags, err := tags.Parse(tag)
			if err != nil {

			}

		}
	}
}
