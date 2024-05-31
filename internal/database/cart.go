package database

import (
	"context"
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

func (d *CartDB) Create(ctx context.Context) (models.Cart, error) {
	query, err := d.db.QueryxContext(ctx, "INSERT INTO carts DEFAULT VALUES RETURNING id")
	if err != nil {
		return models.Cart{}, err
	}
	var cart models.Cart
	if !query.Next() {
		return cart, sql.ErrNoRows
	}
	err = query.Scan(&cart.Id)
	if err != nil {
		return cart, err
	}
	return cart, err
}

func (d *CartDB) Load(ctx context.Context, id int) (models.Cart, error) {
	query := d.db.QueryRowxContext(ctx, "SELECT id FROM carts WHERE id = $1", id)
	if query.Err() != nil {
		return models.Cart{}, query.Err()
	}
	var cart models.Cart
	err := query.Scan(&cart.Id)
	if err != nil {
		return models.Cart{}, err
	}
	return cart, err
}
