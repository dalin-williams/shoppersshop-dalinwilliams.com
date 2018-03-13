package pg

import "github.com/jmoiron/sqlx"

type pay struct {
	db *sqlx.DB
}