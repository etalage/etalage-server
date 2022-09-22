package infra

import "github.com/jmoiron/sqlx"

type DB struct {
	Instance *sqlx.DB
}

func (db *DB) Init() {
}
