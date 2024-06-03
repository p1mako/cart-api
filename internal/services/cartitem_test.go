package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/p1mako/cart-api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CartItemDbMock struct {
	mock.Mock
}

func (db *CartItemDbMock) Create(context context.Context, items ...models.CartItem) ([]models.CartItem, error) {
	args := db.Called(context, items)
	return args.Get(0).([]models.CartItem), args.Error(1)
}
func (db *CartItemDbMock) LoadCartItems(context context.Context, cart int) ([]models.CartItem, error) {
	args := db.Called(context, cart)
	return args.Get(0).([]models.CartItem), args.Error(1)
}
func (db *CartItemDbMock) Remove(context context.Context, item models.CartItem) error {
	args := db.Called(context, item)
	return args.Error(0)
}

func getMockedCreateManipulator() *CartItemManipulator {
	db := new(CartItemDbMock)
	db.On("Create", nil, []models.CartItem{{
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
		On("Create", nil, []models.CartItem{{
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
	sMocked := getMockedCreateManipulator()
	for _, test := range createTests {
		t.Run(test.name, func(t *testing.T) {
			out1, out2 := sMocked.Create(nil, test.input)
			if !assert.ErrorIs(t, out2, test.err) {
				assert.Equal(t, out1, test.expected)
			}
		})
	}
}

func getMockedRemoveManipulator() *CartItemManipulator {
	db := new(CartItemDbMock)
	db.On("Remove", nil, models.CartItem{Id: 1, CartId: 1}).
		Return(nil)
	db.
		On("Remove", nil, models.CartItem{Id: 2, CartId: 1}).
		Return(ErrNoSuchItem{Id: 2})
	db.
		On("Remove", nil, models.CartItem{Id: 1, CartId: 2}).
		Return(ErrNoSuchItem{Id: 1})
	return &CartItemManipulator{db: db}
}

var removeTests = []struct {
	name     string
	input    models.CartItem
	expected error
}{
	{
		name: "Valid",
		input: models.CartItem{
			Id:     1,
			CartId: 1,
		},
		expected: nil,
	},
	{
		name: "Invalid id",
		input: models.CartItem{
			Id:     2,
			CartId: 1,
		},
		expected: ErrNoSuchItem{Id: 2},
	},
	{
		name: "Invalid cartId",
		input: models.CartItem{
			Id:     1,
			CartId: 2,
		},
		expected: ErrNoSuchItem{1},
	},
}

func TestRemove(t *testing.T) {
	sMocked := getMockedRemoveManipulator()
	for _, test := range removeTests {
		t.Run(test.name, func(t *testing.T) {
			out1 := sMocked.Remove(nil, test.input)
			assert.ErrorIs(t, out1, test.expected)
		})
	}
}
