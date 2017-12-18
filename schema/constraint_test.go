package schema

import (
	"github.com/jadengis/icebox/tags"
	"strings"
	"testing"
)

// Test to ensure that the string to contraint type mapping has the correct
// output for all possible inputs.
func TestStringToConstraintMapping(t *testing.T) {
	constraintType, err := getConstraintType(tags.NotNull)
	if err != nil || constraintType != NotNull {
		t.Errorf("constraint type was %d instead of %d", constraintType, NotNull)
	}

	constraintType, err = getConstraintType(tags.Unique)
	if err != nil || constraintType != Unique {
		t.Errorf("constraint type was %d instead of %d", constraintType, Unique)
	}

	constraintType, err = getConstraintType(tags.PrimaryKey)
	if err != nil || constraintType != PrimaryKey {
		t.Errorf("constraint type was %d instead of %d", constraintType, PrimaryKey)
	}

	constraintType, err = getConstraintType(tags.ForeignKey)
	if err != nil || constraintType != ForeignKey {
		t.Errorf("constraint type was %d instead of %d", constraintType, ForeignKey)
	}

	constraintType, err = getConstraintType(tags.Check)
	if err != nil || constraintType != Check {
		t.Errorf("constraint type was %d instead of %d", constraintType, Check)
	}

	constraintType, err = getConstraintType(tags.Default)
	if err != nil || constraintType != Default {
		t.Errorf("constraint type was %d instead of %d", constraintType, Default)
	}

	constraintType, err = getConstraintType(tags.Index)
	if err != nil || constraintType != Index {
		t.Errorf("constraint type was %d instead of %d", constraintType, Index)
	}

	var garbageInput = tags.SubTag("asdf")
	// Test that we error out on an unsupported tag
	constraintType, err = getConstraintType(garbageInput)
	if err == nil {
		t.Errorf("error was not returned for garbage input")
	} else {
		if !(strings.Contains(err.Error(), garbageInput.String())) {
			t.Errorf("error returned by function contains no information" +
				"about the input that caused it")
		}
	}
}

func TestConstraintTypeToString(t *testing.T) {
	typeName := NotNull.String()
	if typeName != "notNull" {
		t.Errorf("NotNull string is incorrect: %s instead of %s", typeName, "notNull")
	}

	typeName = Unique.String()
	if typeName != "unique" {
		t.Errorf("Unique string is incorrect: %s instead of %s", typeName, "unique")
	}

	typeName = PrimaryKey.String()
	if typeName != "primaryKey" {
		t.Errorf("PrimaryKey string is incorrect: %s instead of %s", typeName, "primaryKey")
	}

	typeName = ForeignKey.String()
	if typeName != "foreignKey" {
		t.Errorf("ForeignKey string is incorrect: %s instead of %s", typeName, "foreignKey")
	}

	typeName = Check.String()
	if typeName != "check" {
		t.Errorf("Check string is incorrect: %s instead of %s", typeName, "check")
	}

	typeName = Default.String()
	if typeName != "default" {
		t.Errorf("Default string is incorrect: %s instead of %s", typeName, "default")
	}

	typeName = Index.String()
	if typeName != "index" {
		t.Errorf("Index string is incorrect: %s instead of %s", typeName, "index")
	}
}
