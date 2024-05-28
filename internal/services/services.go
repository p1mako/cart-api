package services

import "github.com/p1mako/cart-api/internal/models"

type ICartItemService interface {
	Create(items models.CartItem) (results models.CartItem, err error)
	GetCartItems(cart int) (items []models.CartItem, err error)
	Remove(item models.CartItem) (err error)
}

type ICartService interface {
	Create() (cart models.Cart, err error)
	AddItem(item models.CartItem) (result models.CartItem, err error)
	RemoveItem(cart models.Cart, item models.CartItem) (result models.Cart, err error)
	Get(id int) (models.Cart, error)
}
