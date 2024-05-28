package services

import (
	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartItemService() *CartItemService {
	return &CartItemService{database.NewCartItemDB()}
}

type CartItemService struct {
	db database.CartItemStorage
}

func (s *CartItemService) Create(items ...models.CartItem) (results []models.CartItem, err error) {
	results, err = s.db.Create(items...)
	if err != nil {
		return
	}
	return
}

func (s *CartItemService) GetCartItems(id int) (items []models.CartItem, err error) {
	items, err = s.db.LoadCartItems(id)
	if err != nil {
		return
	}
	return
}

func (s *CartItemService) Remove(item models.CartItem) (err error) {
	err = s.db.Remove(item)
	return err
}
