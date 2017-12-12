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
