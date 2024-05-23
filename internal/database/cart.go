package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartDB() *CartDB {
	return &CartDB{ConnectDB()}
}

type CartDB struct {
	db *sqlx.DB
}

func (d CartDB) Create() (cart models.Cart, err error) {
	query, err := d.db.Queryx("INSERT INTO carts DEFAULT VALUES")
	if err != nil {
		return
	}
	if !query.Next() {
		return
	}
	err = query.Scan(&cart.Id)
	if err != nil {
		return
	}
	return
}
