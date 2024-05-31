package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartItemDB() *CartItemDB {
	return &CartItemDB{ConnectDB()}
}

type CartItemDB struct {
	db *sqlx.DB
}

func (d *CartItemDB) Create(ctx context.Context, items ...models.CartItem) ([]models.CartItem, error) {
	var results []models.CartItem
	for _, item := range items {
		query, err := d.db.QueryxContext(ctx, "INSERT INTO cartitems(cartid, product, quantity) VALUES ($1, $2, $3) ON CONFLICT (cartid, product) DO UPDATE SET quantity = cartitems.quantity + excluded.quantity RETURNING cartitems.id, cartitems.quantity", item.CartId, item.Product, item.Quantity)
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

func (d *CartItemDB) LoadCartItems(ctx context.Context, cart int) ([]models.CartItem, error) {
	query, err := d.db.QueryxContext(ctx, "SELECT id, cartid, product, quantity FROM cartitems WHERE cartid = $1", cart)
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

func (d *CartItemDB) Remove(ctx context.Context, item models.CartItem) error {
	_, err := d.db.QueryxContext(ctx, "DELETE FROM cartitems WHERE id = $1 AND cartid = $2 RETURNING id", item.Id, item.CartId)
	return err
}
