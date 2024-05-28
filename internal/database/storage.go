package database

import "github.com/p1mako/cart-api/internal/models"

type CartItemStorage interface {
	Create(items ...models.CartItem) (results []models.CartItem, err error)
	LoadCartItems(cart int) (items []models.CartItem, err error)
	Remove(item models.CartItem) (err error)
}

type CartStorage interface {
	Create() (cart models.Cart, err error)
	Load(id int) (cart models.Cart, err error)
}
