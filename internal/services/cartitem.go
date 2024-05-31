package services

import (
	"context"
	"errors"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartItemManipulator() *CartItemManipulator {
	return &CartItemManipulator{database.NewCartItemDB()}
}

type CartItemManipulator struct {
	db database.CartItemStorage
}

func (s *CartItemManipulator) Create(ctx context.Context, item models.CartItem) (models.CartItem, error) {
	results, err := s.db.Create(ctx, item)
	if err != nil {
		return models.CartItem{}, errors.Join(ErrNoSuchCart{item.CartId}, err)
	}
	return results[0], err
}

func (s *CartItemManipulator) GetCartItems(ctx context.Context, id int) ([]models.CartItem, error) {
	return s.db.LoadCartItems(ctx, id)
}

func (s *CartItemManipulator) Remove(ctx context.Context, item models.CartItem) error {
	err := s.db.Remove(ctx, item)
	if err != nil {
		return errors.Join(ErrNoSuchItem{item.Id}, err)
	}
	return err
}
