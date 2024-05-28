package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartItemDB() *CartItemDB {
	return &CartItemDB{ConnectDB()}
}

type CartItemDB struct {
	db *sqlx.DB
}

func (d *CartItemDB) Create(items ...models.CartItem) (results []models.CartItem, err error) {
	for _, item := range items {
		var query *sqlx.Rows
		query, err = d.db.Queryx("INSERT INTO cartitems(cartid, product, quantity) VALUES ($1, $2, $3) RETURNING id", item.CartId, item.Product, item.Quantity)
		if err != nil || !query.Next() {
			return
		}
		err = query.Scan(&item.Id)
		if err != nil {
			return
		}
		results = append(results, item)
	}
	return
}

func (d *CartItemDB) LoadCartItems(cart int) (items []models.CartItem, err error) {
	query, err := d.db.Queryx("SELECT id, cartid, product, quantity FROM cartitems WHERE cartid = $1", cart)
	if err != nil {
		return nil, err
	}
	for query.Next() {
		var item models.CartItem
		err := query.Scan(&item.Id, &item.CartId, &item.Product, &item.Quantity)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return
}

func (d *CartItemDB) Remove(item models.CartItem) (err error) {
	_, err = d.db.Exec("DELETE FROM cartitems WHERE id = $1", item.Id)
	return
}
