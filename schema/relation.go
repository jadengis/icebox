package schema

import (
	"github.com/jadengis/icebox/tags"
)

// Relation represents a relation between two tables in a schema, for example
// if a column in Table A holds a primary key to a row in Table B, these two
// tables are related.
//
// A relation can be described by its type, e.g OneToOne, OneToMany etc., what
// table it points to, and some type specific details.
//
// Type returns the relation type of this relation.
//
// PointsTo returns the table that this relation points to.
//
// Details returns the type specific details of the relation.
type Relation interface {
	Type() RelationType
	PointsTo() Table
	Details() string
}

// RelationType is the type of a relation between two tables.
//
// There are 4 relation types to be supported by icebox.
//
// OneToOne:   One row of the left table corresponds to one in the right table.
//
// OneToMany:  One row of the left table corresponds to many in the right table.
//
// ManyToOne:  Many rows in the left table correspond to one in the right table.
//
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

type relationImpl struct {
	relationType RelationType
	table        *tableImpl
	details      string
}

func (r *relationImpl) Type() RelationType {
	return r.relationType
}

func (r *relationImpl) PointsTo() Table {
	return r.table
}

func (r *relationImpl) Details() string {
	return r.details
}
