package pg

import "github.com/jmoiron/sqlx"

type category struct {
	db *sqlx.DB
}