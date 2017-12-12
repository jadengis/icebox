package icebox

import (
	"github.com/jadengis/icebox/schema"
	"github.com/jadengis/icebox/tags"
	"reflect"
	"strings"
)

type parsedTag struct {
	Name        string
	Constraints []schema.Constraint
	Relations   []schema.Relation
}

func parseTags(tag reflect.StructTag) {
	if tag, ok := tag.Lookup(tags.Icebox); ok {
		tag = strings.Replace(tag, " ", "", -1)
		subtags := strings.Split(tag, tags.Sep)

		// Iterate and parse all subtags.
		seenSubTags := make(map[string]bool)
		for _, subtag := range subtags {

		}
	}
}
