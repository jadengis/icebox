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
	"database/sql"
)

// DB is a wrapper structure for the embedded sql.DB.
type DB struct {
	*sql.DB
}

// Tx is a wrapper structure for the embedded sql.Tx.
type Tx struct {
	*sql.Tx
}

func Open(driver, dataSourceName string) (*DB, error) {
	var db, err = sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
