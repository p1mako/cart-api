package services

import (
	"errors"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartItemService() *CartItemService {
	return &CartItemService{database.NewCartItemDB()}
}

type CartItemService struct {
	db database.CartItemStorage
}

func (s *CartItemService) Create(item models.CartItem) (models.CartItem, error) {
	results, err := s.db.Create(item)
	if err != nil {
		return models.CartItem{}, errors.Join(ErrNoSuchCart{item.CartId}, err)
	}
	return results[0], err
}

func (s *CartItemService) GetCartItems(id int) (items []models.CartItem, err error) {
	items, err = s.db.LoadCartItems(id)
	if err != nil {
		return
	}
	return
}

func (s *CartItemService) Remove(item models.CartItem) error {
	cnt, err := s.db.Remove(item)
	if err != nil {
		return errors.Join(ErrNoSuchCart{item.CartId}, err)
	}
	if cnt != 1 {
		return ErrNoSuchItem{Id: item.Id}
	}
	return err
}
