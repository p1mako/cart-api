package services

import (
	"database/sql"
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

func (s *CartItemService) Remove(item models.CartItem) (err error) {
	err = s.db.Remove(item)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.Join(ErrNoSuchItem{Id: item.Id}, err)
	}
	if err != nil {
		return errors.Join(ErrNoSuchCart{item.CartId}, err)
	}
	return err
}
