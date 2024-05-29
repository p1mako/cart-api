package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartDB() *CartDB {
	return &CartDB{ConnectDB()}
}

type CartDB struct {
	db *sqlx.DB
}

func (d *CartDB) Create() (cart models.Cart, err error) {
	var query *sqlx.Rows
	query, err = d.db.Queryx("INSERT INTO carts DEFAULT VALUES RETURNING id")
	if err != nil {
		return
	}
	if !query.Next() {
		return cart, sql.ErrNoRows
	}
	err = query.Scan(&cart.Id)
	if err != nil {
		return
	}
	return
}

func (d *CartDB) Load(id int) (cart models.Cart, err error) {
	query := d.db.QueryRowx("SELECT id FROM carts WHERE id = $1", id)
	if query.Err() != nil {
		return cart, query.Err()
	}
	err = query.Scan(&cart.Id)
	if err != nil {
		return
	}
	return
}
