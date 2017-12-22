package schema

import (
	"bytes"
	"strings"
	"unicode"
)

const (
	// Separator to use in autogenerated table names.
	nameSeparator string = "_"
)

// Extract a SQL-style from an upper or lower CamelCase string.
func sqlNameFromCamelCase(name string) string {
	var words = splitOnCaps(name)

	// Build a string efficiently from the above slice of words using a buffer.
	var buffer bytes.Buffer
	i := 0
	for ; i < len(words)-1; i++ {
		buffer.WriteString(strings.ToLower(words[i]))
		buffer.WriteString(nameSeparator)
	}
	buffer.WriteString(words[i])
	return buffer.String()
}

// Split the given word into a slice of string all starting with
// capital letters.
func splitOnCaps(word string) []string {
	var words []string
	l := 0
	for s := word; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}
	return words
}
