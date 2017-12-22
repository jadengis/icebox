package icebox

import (
	"time"
)

// Id is a the unique Identity type for icebox Model objects.
// An Id should unique identify a given persisted object.
type Id uint

// IdEntity provides a method set for getting and setting an objects
// internal Id. In short, IdEntity is a contract insuring that the implementing
// entity indeed has an Id.
//
// Id provides access to the underlying Id.
// SetId allows setting of the underlying Id.
type IdEntity interface {
	Id() Id
	SetId(Id)
}

// Model is the base structure underlying all icebox relational entities.
// For a new type to be recognized as an icebox entity, it must contain this
// Model as an embedded type.
type Model struct {
	Id        Id         `icebox:"column:id,primaryKey"`
	CreatedAt *time.Time `icebox:"column:created_at"`
	UpdatedAt *time.Time `icebox:"column:updated_at"`
}

// Id retrieves this Models underlying Id.
func (m *Model) Id() Id {
	return m.id
}

// SetId sets this Models underlying Id.
func (m *Model) SetId(id Id) {
	m.id = id
}
