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
		query, err := d.db.Queryx("INSERT INTO cartitems(cartid, product, quantity) VALUES ($1, $2, $3)", item.CartId, item.Product, item.Quantity)
		if err != nil {
			return results, err
		}
		if !query.Next() {
			return results, err
		}
		err = query.Scan(&item.Id)
		if err != nil {
			return results, err
		}
		results = append(results, item)
	}
	return
}

func (d *CartItemDB) GetCartItems(cart int) ([]models.CartItem, error) {
	panic("unimplemented")
}

func (d *CartItemDB) Remove(item models.CartItem) error {
	panic("unimplemented")
}
