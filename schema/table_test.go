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
	"testing"
)

const (
	tableName string = "fun_table_name"
)

// A simple struct.
type mockData struct {
	num  int
	word string
}

// Struct implementing the TableEntity interface.
type namedMockData struct {
	word string
}

func (n namedMockData) TableName() string {
	return tableName
}

// Test that we generate the right names for camel case structs.
func TestTableNameFromObject(t *testing.T) {
	mock := mockData{num: 5, word: "stuff"}
	name := tableNameFromObject(mock)
	expected := "mock_data"
	if name != "mock_datas" {
		t.Errorf("table name was not %s as expected: name = %s", expected, name)
	}
}

// Check that we generate the right table names.
func TestGetTableName(t *testing.T) {
	testCases := []struct {
		input  interface{}
		result string
	}{
		{namedMockData{"data"}, tableName},
		{mockData{5, "data"}, "mock_datas"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test table name object %d", i),
			func(t *testing.T) {
				name := getTableName(tc.input)
				if name != tc.result {
					t.Errorf("table name incorrect: name = %s, expected = %s",
						name, tc.result)
				}
			},
		)
	}
}
