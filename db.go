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
