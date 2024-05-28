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

func (s *CartService) Create() (cart models.Cart, err error) {
	cart, err = s.db.Create()
	return
}

func (s *CartService) AddItem(item models.CartItem) (result models.CartItem, err error) {
	if item.Quantity <= 0 {
		return result, ErrBadQuantity
	}
	if item.Product == "" {
		return result, ErrNoProductName
	}
	return s.itemServ.Create(item)
}

func (s *CartService) RemoveItem(item models.CartItem) (err error) {
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
