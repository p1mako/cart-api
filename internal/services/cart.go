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
	db database.CartStorage
}

func (s *CartService) Create() (cart models.Cart, err error) {
	cart, err = s.db.Create()
	return
}

func (s *CartService) AddItem(cart models.Cart, item models.CartItem) (result models.Cart, err error) {
	result, err = s.Get(cart.Id)
	if item.Quantity <= 0 {
		return result, ErrBadQuantity
	}
	if item.Product == "" {
		return result, ErrNoProductName
	}
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
	result, err = s.Get(cart.Id)
	if err != nil {
		return
	}
	itemServ := NewCartItemService()
	err = itemServ.Remove(item)
	if err != nil {
		return
	}
	result.Items, err = itemServ.GetCartItems(result.Id)
	if err != nil {
		return
	}
	return
}

func (s *CartService) Get(id int) (models.Cart, error) {
	oldCart, err := s.db.Load(id)
	if errors.Is(err, sql.ErrNoRows) {
		return oldCart, ErrNoSuchCart{Id: id}
	}
	itemServ := NewCartItemService()
	oldCart.Items, err = itemServ.GetCartItems(id)
	return oldCart, err
}

type ErrNoSuchCart struct {
	Id int
}

func (e ErrNoSuchCart) Error() string {
	return fmt.Sprintf("Cannot find cart with id %v\n", e.Id)
}

var ErrNoProductName = errors.New("no product name provided")
var ErrBadQuantity = errors.New("quantity of product is non positive")
