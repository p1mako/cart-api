package services

import (
	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartService() *CartService {
	return &CartService{db: database.NewCartDB()}
}

type CartService struct {
	db *database.CartDB
}

func (s CartService) Create() (cart models.Cart, err error) {
	cart, err = s.db.Create()
	if err != nil {
		return
	}
	return
}
