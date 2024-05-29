package services

import (
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

func (s *CartManipulator) Create() (models.Cart, error) {
	return s.db.Create()
}

func (s *CartManipulator) AddItem(item models.CartItem) (models.CartItem, error) {
	if item.Quantity <= 0 {
		return models.CartItem{}, ErrBadQuantity
	}
	if item.Product == "" {
		return models.CartItem{}, ErrNoProductName
	}
	return s.itemServ.Create(item)
}

func (s *CartManipulator) RemoveItem(item models.CartItem) error {
	return s.itemServ.Remove(item)
}

func (s *CartManipulator) Get(id int) (models.Cart, error) {
	oldCart, err := s.db.Load(id)
	if errors.Is(err, sql.ErrNoRows) || oldCart.Id == 0 {
		return oldCart, ErrNoSuchCart{Id: id}
	}
	oldCart.Items, err = s.itemServ.GetCartItems(id)
	return oldCart, err
}
