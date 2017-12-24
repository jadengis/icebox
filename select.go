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
//"database/sql"
//"reflect"
)

// A Selecter can populate itself via a select query against a given DB.
type Selecter interface {
	// Select runs this objects Select query against the given DB.
	Select(*DB) error
	// SelectTx runs this objects Select query against the given Tx.
	SelectTx(*Tx) error
}

func (m *Model) Select(db *DB) error {
	var tx, err = db.Begin()
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) SelectTx(db *DB) error {

}
