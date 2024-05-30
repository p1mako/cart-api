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

func (d *CartItemDB) Create(items ...models.CartItem) ([]models.CartItem, error) {
	var results []models.CartItem
	for _, item := range items {
		query, err := d.db.Queryx("INSERT INTO cartitems(cartid, product, quantity) VALUES ($1, $2, $3) ON CONFLICT (cartid, product) DO UPDATE SET quantity = cartitems.quantity + excluded.quantity RETURNING cartitems.id, cartitems.quantity", item.CartId, item.Product, item.Quantity)
		if err != nil || !query.Next() {
			return nil, err
		}
		err = query.Scan(&item.Id, &item.Quantity)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}
	return results, nil
}

func (d *CartItemDB) LoadCartItems(cart int) ([]models.CartItem, error) {
	query, err := d.db.Queryx("SELECT id, cartid, product, quantity FROM cartitems WHERE cartid = $1", cart)
	if err != nil {
		return nil, err
	}
	var items []models.CartItem
	for query.Next() {
		var item models.CartItem
		err := query.Scan(&item.Id, &item.CartId, &item.Product, &item.Quantity)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (d *CartItemDB) Remove(item int) (int, error) {
	res, err := d.db.Exec("DELETE FROM cartitems WHERE id = $1", item)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return int(affected), err
	}
	return int(affected), err
}
