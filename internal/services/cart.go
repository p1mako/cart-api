package services

import (
	"github.com/p1mako/cart-api/internal/database"
)

func NewCartService() *CartService {
	return &CartService{db: database.NewCartDB()}
}

type CartService struct {
	db *database.CartDB
}
