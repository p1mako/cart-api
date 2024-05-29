package services

import (
	"database/sql"
	"errors"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/models"
)

func NewCartService() *CartService {
	return &CartService{db: database.NewCartDB(), itemServ: NewCartItemService()}
}

type CartService struct {
	db       database.CartStorage
	itemServ ICartItemService
}

func (s *CartService) Create() (models.Cart, error) {
	return s.db.Create()
}

func (s *CartService) AddItem(item models.CartItem) (models.CartItem, error) {
	if item.Quantity <= 0 {
		return models.CartItem{}, ErrBadQuantity
	}
	if item.Product == "" {
		return models.CartItem{}, ErrNoProductName
	}
	return s.itemServ.Create(item)
}

func (s *CartService) RemoveItem(item models.CartItem) error {
	return s.itemServ.Remove(item)
}

func (s *CartService) Get(id int) (models.Cart, error) {
	oldCart, err := s.db.Load(id)
	if errors.Is(err, sql.ErrNoRows) {
		return oldCart, ErrNoSuchCart{Id: id}
	}
	oldCart.Items, err = s.itemServ.GetCartItems(id)
	return oldCart, err
}
