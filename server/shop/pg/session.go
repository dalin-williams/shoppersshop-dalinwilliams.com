package pg

import "github.com/jmoiron/sqlx"

type session struct {
	db *sqlx.DB
}
