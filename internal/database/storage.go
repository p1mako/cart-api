package database

import (
	"context"

	"github.com/p1mako/cart-api/internal/models"
)

type CartItemStorage interface {
	Create(ctx context.Context, items ...models.CartItem) ([]models.CartItem, error)
	LoadCartItems(ctx context.Context, cart int) ([]models.CartItem, error)
	Remove(ctx context.Context, item models.CartItem) error
}

type CartStorage interface {
	Create(ctx context.Context) (models.Cart, error)
	Load(ctx context.Context, id int) (models.Cart, error)
}
