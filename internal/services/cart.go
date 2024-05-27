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

func (s *CartService) Create() (cart models.Cart, err error) {
	cart, err = s.db.Create()
	return
}

func (s *CartService) AddItem(cart models.Cart, item models.CartItem) (models.Cart, error) {
	panic("unimplemented")
}

func (s *CartService) RemoveItem(cart models.Cart, item models.CartItem) (models.Cart, error) {
	panic("unimplemented")
}

func (s *CartService) View(cart models.Cart) (models.Cart, error) {
	panic("unimplemented")
}
