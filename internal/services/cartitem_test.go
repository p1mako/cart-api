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

func getMockedCreateManipulator() *CartItemManipulator {
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
	sMocked := getMockedCreateManipulator()
	for _, test := range createTests {
		t.Run(test.name, func(t *testing.T) {
			out1, out2 := sMocked.Create(test.input)
			if !assert.ErrorIs(t, out2, test.err) {
				assert.Equal(t, out1, test.expected)
			}
		})
	}
}

func getMockedRemoveManipulator() *CartItemManipulator {
	db := new(CartItemDbMock)
	db.On("Remove", models.CartItem{
		Id:     1,
		CartId: 1,
	}).
		Return(1, nil)
	db.
		On("Remove", models.CartItem{
			Id:     2,
			CartId: 1,
		}).
		Return(0, nil)
	db.
		On("Remove", models.CartItem{
			Id:     1,
			CartId: 2,
		}).
		Return(0, nil)
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
		name: "Invalid cart id",
		input: models.CartItem{
			Id:     1,
			CartId: 2,
		},
		expected: ErrNoSuchCart{Id: 2},
	},
	{
		name: "Invalid id",
		input: models.CartItem{
			Id:     2,
			CartId: 1,
		},
		expected: ErrNoSuchItem{Id: 2},
	},
}

func TestRemove(t *testing.T) {
	sMocked := getMockedRemoveManipulator()
	for _, test := range removeTests {
		t.Run(test.name, func(t *testing.T) {
			out1 := sMocked.Remove(test.input)
			assert.ErrorIs(t, out1, test.expected)
		})
	}
}
