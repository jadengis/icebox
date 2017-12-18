package tags

import (
	"fmt"
	"strings"
	"testing"
)

// Test that the parser errors appropriately
func TestParseErrors(t *testing.T) {
	var unknown string = "column:id,asdf"
	_, err := Parse(unknown)

	// Parse should error for garbage input
	if err == nil {
		t.Errorf("error not raised for garbage input")
	} else {
		// Error should mention its unknown
		if !strings.Contains(err.Error(), "unknown") {
			t.Errorf("raised error doesn't mention it is unknown: error = %s", err.Error())
		}
		// Error should contain some identifiable error info
		if !strings.Contains(err.Error(), "asdf") {
			t.Errorf("raised error doesn't mention the unknown tag: error = %s", err.Error())
		}
	}

	var duplicate string = "column:id,column:name"
	_, err = Parse(duplicate)

	// Parse should error duplicate.
	if err == nil {
		t.Errorf("error not raised for duplicate tag")
	} else {
		// Error should mention its a duplicate
		if !strings.Contains(err.Error(), "duplicate") {
			t.Errorf("raised error doesn't it is duplicate: error = %s", err.Error())
		}
		// Error should contain some identifiable error info
		if !strings.Contains(err.Error(), "column") {
			t.Errorf("raised error doesn't mention the duplicate: error = %s", err.Error())
		}
	}
}

// Test the tag parsing
func TestParse(t *testing.T) {
	testCases := []struct {
		tag     string
		subTags []SubTag
		infos   []string
	}{
		{"column:id", []SubTag{Column}, []string{"id"}},
		{"column:id,primaryKey,default:0",
			[]SubTag{Column, PrimaryKey, Default},
			[]string{"id", "", "0"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("test parse %s", tc.tag),
			func(t *testing.T) {
				tagMap, err := Parse(tc.tag)

				if err != nil {
					t.Errorf("unexpected error parsing %s: error = %s", tc.tag, err.Error())
				} else {
					for i, subTag := range tc.subTags {
						if val, ok := tagMap.GetInfo(subTag); ok {
							if val != tc.infos[i] {
								t.Errorf("parsed tag info incorrect: %s instead of %s", val, tc.infos[i])
							}
						} else {
							t.Errorf("parsed tag missing key %s", subTag)
						}
					}
				}
			})
	}
}
