package services

import (
	"database/sql"
	"testing"

	"github.com/p1mako/cart-api/internal/models"
	"github.com/stretchr/testify/assert"
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

func getCartManipulatorWithMock() *CartItemManipulator {
	db := new(CartItemDbMock)
	db.On("Create", []models.CartItem{{
		CartId:   1,
		Product:  "qqqqq",
		Quantity: 10,
	}}).
		Return(
			[]models.CartItem{{
				Id:       1,
				CartId:   1,
				Product:  "qqqqq",
				Quantity: 10,
			}},
			nil,
		)
	db.
		On("Create", []models.CartItem{{
			CartId:   2,
			Product:  "qqqqq",
			Quantity: 10,
		}}).
		Return(
			[]models.CartItem{},
			sql.ErrNoRows,
		)
	return &CartItemManipulator{db: db}
}

var createTests = []struct {
	name     string
	input    models.CartItem
	expected models.CartItem
	err      error
}{
	{
		name: "Valid",
		input: models.CartItem{
			CartId:   1,
			Product:  "qqqqq",
			Quantity: 10,
		},
		expected: models.CartItem{
			Id:       1,
			CartId:   1,
			Product:  "qqqqq",
			Quantity: 10,
		},
	},
	{
		name: "Invalid cart",
		input: models.CartItem{
			CartId:   2,
			Product:  "qqqqq",
			Quantity: 10,
		},
		err: ErrNoSuchCart{Id: 2},
	},
}

func TestCreate(t *testing.T) {
	sMocked := getCartManipulatorWithMock()
	for _, test := range createTests {
		t.Run(test.name, func(t *testing.T) {
			out1, out2 := sMocked.Create(test.input)
			if !assert.ErrorIs(t, out2, test.err) {
				assert.Equal(t, out1, test.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {

}
