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

package tags

// Tag represents a struct-level tag.
type Tag string

// String maps a Tag to its string representation.
func (t Tag) String() string {
	return string(t)
}

// SubTag is like a tag but differ in that they should never appear at the
// struct-level, they will always live in the payload of another tag.
type SubTag string

// String maps a SubTag to its string representation.
func (t SubTag) String() string {
	return string(t)
}

// ParsedTag is the result of parsing a Tags SubTag string. It is essentially
// a mapping a subtag to its sub tag info.
type ParsedTag map[SubTag]string

// GetInfo retrieves the sub tag info for the given subtag
// from the given ParsedTag.
func (t ParsedTag) GetInfo(subTag SubTag) (string, bool) {
	val, found := map[SubTag]string(t)[subTag]
	return val, found
}

// Const Declarations of Tag Literals
//
// Icebox: The tag for scoping all icebox subtags. Any StructField with the
// Icebox tag will be considered a persisted field.
const (
	Icebox Tag = "icebox"
)

// Const Declarations for SubTag Literals
//
// Column:     The subtag for marking a field as a table column. Subtag info
// contains the name of the column to use.
//
// NotNull:    The subtag for marking a field as not nullable.
//
// Unique:     The subtag for marking a field as unique in the column.
//
// PrimaryKey: The subtag for marking a field as primary key in the table.
//
// ForeignKey: The subtag for marking a field as a foreign key in another table.
// Subtag info contains information about the target table.
//
// Check:      The subtag for marking a field as
//
// Default:    The subtag for specifying a default value to use for this field.
// Subtag info contains the default to use.
//
// Index:      The subtag for specifying a field should be indexed.
//
// OneToOne:   The subtag for making a one to one relationship between a
// and one of its fields.
//
// OneToMany:  The subtag for making a one to many relationship between a
// column and another table. Subtag info contains field name to select on.
//
// ManyToOne:  The subtag for marking a many to one relationship between a table
// and one of its fields.
//
// ManyToMany: TODO
const (
	Column     SubTag = "column"
	NotNull    SubTag = "notNull"
	Unique     SubTag = "unique"
	PrimaryKey SubTag = "primaryKey"
	ForeignKey SubTag = "foreignKey"
	Check      SubTag = "check"
	Default    SubTag = "default"
	Index      SubTag = "index"
	OneToOne   SubTag = "oneToOne"
	OneToMany  SubTag = "oneToMany"
	ManyToOne  SubTag = "manyToOne"
	ManyToMany SubTag = "manyToMany"
)

// Mapping from subtag string name to subtag.
var subTagMap = map[string]SubTag{
	Column.String():     Column,
	NotNull.String():    NotNull,
	Unique.String():     Unique,
	PrimaryKey.String(): PrimaryKey,
	ForeignKey.String(): ForeignKey,
	Check.String():      Check,
	Default.String():    Default,
	Index.String():      Index,
	OneToOne.String():   OneToOne,
	OneToMany.String():  OneToMany,
	ManyToOne.String():  ManyToOne,
	ManyToMany.String(): ManyToMany,
}
