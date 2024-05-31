package database

import "github.com/p1mako/cart-api/internal/models"

type CartItemStorage interface {
	Create(items ...models.CartItem) ([]models.CartItem, error)
	LoadCartItems(cart int) ([]models.CartItem, error)
	Remove(item models.CartItem) error
}

type CartStorage interface {
	Create() (models.Cart, error)
	Load(id int) (models.Cart, error)
}
