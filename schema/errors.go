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
	"fmt"
	"reflect"
)

// Error type for missing object during a lookup.
type notFoundError struct {
	key interface{}
	msg string
}

// Produce and error message for a notFoundError.
func (e *notFoundError) Error() string {
	return fmt.Sprintf("key not found: key = %s, msg = %s", e.key, e.msg)
}

// This is the type of error to raise if the typeName cannot be resolved.
//
// typeName is the unresolved typeName.
//
// msg is an error message to return.
type unknownTypeError struct {
	typeName string
	msg      string
}

// Produce an error message for an unknownTypeError.
func (e *unknownTypeError) Error() string {
	return e.msg + " : " + e.typeName
}

// schemaGenError is a general wrapper for schema generation errors.
type schemaGenError struct {
	cause error
	msg   string
}

// Error message for a schema generation error.
func (e *schemaGenError) Error() string {
	return fmt.Sprintf("%s : %s", e.msg, e.cause.Error())
}

// Schema generation error relating to a bad type.
type typeError struct {
	badType reflect.Type
	msg     string
}

// Error message for this type error.
func (e *typeError) Error() string {
	return fmt.Sprintf("unsupported type %s : %s", e.badType, e.msg)
}
