package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartManipulator() *CartManipulator {
	return &CartManipulator{db: database.NewCartDB(), itemServ: NewCartItemManipulator()}
}

type CartManipulator struct {
	db       database.CartStorage
	itemServ CartItemService
}

func (s *CartManipulator) Create(ctx context.Context) (models.Cart, error) {
	return s.db.Create(ctx)
}

func (s *CartManipulator) AddItem(ctx context.Context, item models.CartItem) (models.CartItem, error) {
	if item.Quantity <= 0 {
		return models.CartItem{}, ErrBadQuantity
	}
	if item.Product == "" {
		return models.CartItem{}, ErrNoProductName
	}
	return s.itemServ.Create(ctx, item)
}

func (s *CartManipulator) RemoveItem(ctx context.Context, item models.CartItem) error {
	_, err := s.Get(ctx, item.CartId)
	if err != nil {
		return ErrNoSuchCart{Id: item.CartId}
	}
	return s.itemServ.Remove(ctx, item)
}

func (s *CartManipulator) Get(ctx context.Context, id int) (models.Cart, error) {
	oldCart, err := s.db.Load(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return oldCart, ErrNoSuchCart{Id: id}
	}
	oldCart.Items, err = s.itemServ.GetCartItems(ctx, id)
	return oldCart, err
}
