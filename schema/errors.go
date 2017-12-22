package schema

import (
	"fmt"
	"reflect"
)

// This is the type of error to raise if the typeName cannot be resolved.
// typeName is the unresolved typeName.
// msg is an error message to return.
type unknownTypeError struct {
	typeName string
	msg      string
}

// Produce an error message for an unknownTypeError.
func (e unknownTypeError) Error() string {
	return e.msg + " : " + e.typeName
}

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
