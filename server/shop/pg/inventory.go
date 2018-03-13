package pg

import "github.com/jmoiron/sqlx"

type inventory struct {
	db *sqlx.DB
}
