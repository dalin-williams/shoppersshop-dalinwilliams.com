package pg

import "github.com/jmoiron/sqlx"

type vend struct {
	db *sqlx.DB
}