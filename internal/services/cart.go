package services

import "github.com/p1mako/cart-api/internal/database"

type CartService struct {
	db *database.CartDB
}
