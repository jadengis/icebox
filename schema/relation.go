package schema

import (
	"github.com/jadengis/icebox/tags"
)

// RelationType is the type of a relation between two tables.
// There are 3 relation types to be supported by icebox.
// OneToMany:  One row of the left table corresponds to many in the right table.
// ManyToOne:  Many rows in the left table correspond to one in the right table.
// ManyToMany: Many rows in the left table correspond to many in the right table.
type RelationType int

const (
	InvalidRelation RelationType = -1
	OneToOne        RelationType = iota
	OneToMany
	ManyToOne
	ManyToMany
)

func getRelationType(typeName tags.SubTag) (RelationType, error) {
	switch typeName {
	case tags.OneToOne:
		return OneToOne, nil
	case tags.OneToMany:
		return OneToMany, nil
	case tags.ManyToOne:
		return ManyToOne, nil
	case tags.ManyToMany:
		return ManyToMany, nil
	default:
		return InvalidRelation, &unknownTypeError{
			typeName: typeName.String(),
			msg:      "typeName could not be resolved"}
	}

}

type Relation struct {
	Type  RelationType
	Table *Table
	Info  string
}
