package database

import "github.com/jmoiron/sqlx"

func NewCartDB() *CartDB {
	return &CartDB{ConnectDB()}
}

type CartDB struct {
	db *sqlx.DB
}
