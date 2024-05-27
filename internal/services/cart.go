package services

import (
	"database/sql"
	"errors"
	"fmt"

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

func (s *CartService) AddItem(cart models.Cart, item models.CartItem) (result models.Cart, err error) {
	result, err = s.GetCart(cart.Id)
	if err != nil {
		return
	}
	itemServ := NewCartItemService()
	created, err := itemServ.Create(item)
	if err != nil {
		return
	}
	result.Items = append(result.Items, created...)
	return
}

func (s *CartService) RemoveItem(cart models.Cart, item models.CartItem) (result models.Cart, err error) {
	result, err = s.GetCart(cart.Id)
	if err != nil {
		return
	}
	itemServ := NewCartItemService()
	err = itemServ.Remove(item)
	if err != nil {
		return
	}
	items, err := itemServ.GetCartItems(result)
	if err != nil {
		return
	}
	result.Items = append(result.Items, items...)
	return
}

func (s *CartService) View(cart models.Cart) (result models.Cart, err error) {
	result, err = s.GetCart(cart.Id)
	if err != nil {
		return
	}
	itemServ := NewCartItemService()
	items, err := itemServ.GetCartItems(cart)
	if err != nil {
		return
	}
	result.Items = append(result.Items, items...)
	return
}

func (s *CartService) GetCart(id int) (models.Cart, error) {
	oldCart, err := s.db.Get(id)
	if errors.Is(err, sql.ErrNoRows) {
		return oldCart, errors.Join(errors.New(fmt.Sprintf("no cart was found with id %v: ", id)), err)
	}
	return oldCart, err
}
