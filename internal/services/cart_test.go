package services

import (
	"testing"

	"github.com/p1mako/cart-api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CartDbMocked struct {
	mock.Mock
}

func (o *CartDbMocked) Create() (models.Cart, error) {
	args := o.Called()
	return args.Get(0).(models.Cart), args.Error(1)
}

func (o *CartDbMocked) Load(id int) (models.Cart, error) {
	args := o.Called(id)
	return args.Get(0).(models.Cart), args.Error(1)
}

type CartItemServiceMocked struct {
	mock.Mock
}

func (o *CartItemServiceMocked) Create(items models.CartItem) (models.CartItem, error) {
	args := o.Called(items)
	return args.Get(0).(models.CartItem), args.Error(1)
}

func (o *CartItemServiceMocked) GetCartItems(cart int) ([]models.CartItem, error) {
	args := o.Called(cart)
	return args.Get(0).([]models.CartItem), args.Error(1)
}

func (o *CartItemServiceMocked) Remove(item models.CartItem) error {
	args := o.Called(item)
	return args.Error(0)
}

func getDependencies() (*CartDbMocked, *CartItemServiceMocked) {
	cartDb := new(CartDbMocked)
	cartItemServ := new(CartItemServiceMocked)
	cartDb.
		On("Create").
		Return(models.Cart{
			Id:    1,
			Items: nil,
		}, nil)
	cartDb.
		On("Load", 1).
		Return(models.Cart{
			Id:    1,
			Items: nil,
		}, nil)
	cartItemServ.
		On("Create", models.CartItem{
			CartId:   1,
			Product:  "qwerty",
			Quantity: 1,
		}).
		Return(models.CartItem{
			Id:       1,
			CartId:   1,
			Product:  "qwerty",
			Quantity: 1,
		}, nil)
	cartItemServ.
		On("GetCartItems", 1).
		Return([]models.CartItem{{
			Id:       1,
			CartId:   1,
			Product:  "qwerty",
			Quantity: 1,
		}}, nil)
	cartItemServ.
		On("Remove", models.CartItem{Id: 1, CartId: 1}).
		Return(nil)
	cartItemServ.
		On("Remove", models.CartItem{Id: 2, CartId: 1}).
		Return(ErrNoSuchItem{})
	return cartDb, cartItemServ
}

var addItemTests = []struct {
	input    models.CartItem
	expected models.CartItem
	name     string
	err      error
}{
	{
		name: "Valid",
		input: models.CartItem{
			CartId:   1,
			Product:  "qwerty",
			Quantity: 1,
		},
		expected: models.CartItem{
			Id:       1,
			CartId:   1,
			Product:  "qwerty",
			Quantity: 1,
		},
		err: nil,
	},
	{
		name: "No name",
		input: models.CartItem{
			CartId:   1,
			Product:  "",
			Quantity: 1,
		},
		err: ErrNoProductName,
	},
	{
		name: "Invalid quantity",
		input: models.CartItem{
			CartId:   1,
			Product:  "qwerty",
			Quantity: 0,
		},
		err: ErrBadQuantity,
	},
}

func TestCartService_AddItem(t *testing.T) {
	cartDb, cartItemServ := getDependencies()

	cartServ := CartManipulator{
		db:       cartDb,
		itemServ: cartItemServ,
	}
	for _, test := range addItemTests {
		t.Run(test.name, func(t *testing.T) {
			out1, out2 := cartServ.AddItem(test.input)
			assert.Equal(t, out1, test.expected)
			assert.Equal(t, out2, test.err)
		})
	}
}
