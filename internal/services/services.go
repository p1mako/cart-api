package services

import (
	"errors"
	"fmt"

	"github.com/p1mako/cart-api/internal/models"
)

type ICartItemService interface {
	Create(items models.CartItem) (results models.CartItem, err error)
	GetCartItems(cart int) (items []models.CartItem, err error)
	Remove(item models.CartItem) (err error)
}

type ICartService interface {
	Create() (cart models.Cart, err error)
	AddItem(item models.CartItem) (result models.CartItem, err error)
	RemoveItem(item models.CartItem) (err error)
	Get(id int) (models.Cart, error)
}

type ErrNoSuchCart struct {
	Id int
}

func (e ErrNoSuchCart) Error() string {
	return fmt.Sprintf("Cannot find cart with id %v\n", e.Id)
}

func (e ErrNoSuchCart) Is(target error) bool {
	var converted ErrNoSuchCart
	ok := errors.As(target, &converted)
	if !ok {
		return false
	}
	return e.Id == converted.Id
}

type ErrNoSuchItem struct {
	Id int
}

func (e ErrNoSuchItem) Error() string {
	return fmt.Sprintf("Cannot find cart with id %v\n", e.Id)
}

func (e ErrNoSuchItem) Is(target error) bool {
	var converted ErrNoSuchItem
	ok := errors.As(target, &converted)
	if !ok {
		return false
	}
	return e.Id == converted.Id
}

var ErrNoProductName = errors.New("no product name provided")
var ErrBadQuantity = errors.New("quantity of product is non positive")
