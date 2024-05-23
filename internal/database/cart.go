package database

import "github.com/jmoiron/sqlx"

type CartDB struct {
	db *sqlx.DB
}
