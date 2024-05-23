package services

import (
	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartItemService() *CartItemService {
	return &CartItemService{database.NewCartItemDB()}
}

type CartItemService struct {
	db *database.CartItemDB
}

func (s CartItemService) Create(items ...models.CartItem) (results []models.CartItem, err error) {
	results, err = s.db.Create(items...)
	if err != nil {
		return
	}
	return
}
