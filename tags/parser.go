package tags

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const (
	// Sep is the icebox subtag. Subtags of the icebox tag will be separated by
	// this value.
	subTagSeparator  string = ","
	payloadSeparator string = ":"
)

// Type of error to return when the tag is invalid.
type invalidTagError struct {
	tag string
	msg string
}

// Error producing logic for invalid tag error.
func (e *invalidTagError) Error() string {
	return fmt.Sprintf("the given tag %s is invalid : %s", e.tag, e.msg)
}

// Parse will parse the given subtags and produce a mapping between existing
// subtags and there subtag info (if available).
func Parse(subtags string) (map[SubTag]string, error) {
	result := make(map[SubTag]string)
	// Make a map to store already seen subtags
	seenSubTags := make(map[string]bool)

	// Scan the subtag string for subtag separator delimited chunks.
	tagScanner := bufio.NewScanner(strings.NewReader(subtags))
	tagScanner.Split(scanSubTagsSeparators)
	for tagScanner.Scan() {
		name, info := parseNameAndInfo(tagScanner.Text())

		// Validate the name
		if _, found := seenSubTags[name]; found {
			// Tag is duplicate so error
			return nil, &invalidTagError{
				tag: name,
				msg: "tag is duplicate"}
		}

		subtag, found := subTagMap[name]
		if !found {
			// Tag is invalid so error
			return nil, &invalidTagError{
				tag: name,
				msg: "tag is unrecognized"}
		}

		// Tag is valid so add to return value
		result[subtag] = info
	}
	return result, nil
}

// Parse the subtag name and info field from the given subtag string.
func parseNameAndInfo(subtag string) (name, info string) {
	if i := strings.Index(subtag, payloadSeparator); i >= 0 {
		name = subtag[0 : i-1]
		if i == (len(subtag) - 1) {
			return name, ""
		}
		info = subtag[i+1 : len(subtag)-1]
		return name, info

	}
	return subtag, ""
}

// SplitFunc for separating on subtag separators.
func scanSubTagsSeparators(
	data []byte, atEOF bool) (
	advance int, token []byte, err error) {
	return scanSeparators(data, atEOF, subTagSeparator)
}

// Common method for implementing the parsing split funcs.
func scanSeparators(
	data []byte, atEOF bool, separator string) (
	advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// If we find a separator, return all the bytes from the start to
	// the separator.
	if i := bytes.IndexAny(data, separator); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// At EOF, and no more separators, so return everything.
	if atEOF {
		return len(data), data, nil
	}

	// Request more bytes
	return 0, nil, nil
}
