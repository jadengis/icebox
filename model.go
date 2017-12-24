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
