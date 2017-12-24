package schema

import (
	"github.com/jadengis/icebox/types"
	"reflect"
	"testing"
)

type fakeRelation struct {
	id int `icebox:"column,primaryKey"`
}

type fakeStruct struct {
	Id    int            `icebox:"column,primaryKey"`
	Float float32        `icebox:"column:float_number"`
	Slice []fakeRelation `icebox:"oneToMany"`
}

func TestGenerateSchema(t *testing.T) {
	// generate a schema and test it for correctness
	schema, err := NewSchema("test_schema", new(fakeStruct), new(fakeRelation))
	if err != nil {
		t.Fatalf("schema could not be generated: error = %s", err.Error())
	}
	if schema.Name() != "test_schema" {
	}
}

func TestGenerateTable(t *testing.T) {
	// generate a table and test it for correctness
	table, err := generateTable(new(fakeStruct))

	// Validate that table properties were appropriately generated.
	if err != nil {
		t.Errorf("table failed to generate: error = %s", err.Error())
	}
	if table.Name() != "fake_structs" {
		t.Errorf("table name is incorrect: name = %s", table.Name())
	}
	if table.Type() != reflect.TypeOf((*fakeStruct)(nil)).Elem() {
		t.Errorf("table type is incorrect: type = %s, expected = %s",
			table.Type(), reflect.TypeOf((*fakeStruct)(nil)).Elem())
	}

	// Verify the id column is the way it should be.
	idColumn, err := table.ColumnFor("id")
	if err != nil {
		t.Errorf("table is missing id column")
	} else {
		if idColumn.Type().Type() != types.Int {
			t.Errorf("id column type is incorrect: type = %d, expected = %d",
				idColumn.Type().Type(), types.Int)
		}
		if idColumn.Name() != "id" {
			t.Errorf("id column name is incorrect: name = %s, expected = %s",
				idColumn.Name(), "id")
		}
		if _, found := idColumn.ConstraintFor(PrimaryKey); !found {
			t.Errorf("id column missing expected PrimaryKey constraint")
		}
	}

	// Verify the float_number column is the way it should be.
	// This is verifying that custom names work.
	floatColumn, err := table.ColumnFor("float_number")
	if err != nil {
		t.Errorf("table is missing float_number column")
	} else {
		if floatColumn.Type().Type() != types.Float {
			t.Errorf("float column type is incorrect: type = %d, expected = %d",
				floatColumn.Type().Type(), types.Float)
		}
		if floatColumn.Name() != "float_number" {
			t.Errorf("float column name is incorrect: name = %s, expected = %s",
				floatColumn.Name(), "float_number")
		}
	}
}
