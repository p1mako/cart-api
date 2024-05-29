package services

import (
	"github.com/p1mako/cart-api/internal/models"
	"github.com/stretchr/testify/mock"
)

type CartItemDbMock struct {
	mock.Mock
}

func (db *CartItemDbMock) Create(items ...models.CartItem) ([]models.CartItem, error) {
	args := db.Called(items)
	return args.Get(0).([]models.CartItem), args.Error(1)
}
func (db *CartItemDbMock) LoadCartItems(cart int) ([]models.CartItem, error) {
	args := db.Called(cart)
	return args.Get(0).([]models.CartItem), args.Error(1)
}
func (db *CartItemDbMock) Remove(item models.CartItem) (int, error) {
	args := db.Called(item)
	return args.Int(0), args.Error(1)
}
