package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/p1mako/cart-api/internal/models"
	"github.com/p1mako/cart-api/internal/services"
)

func NewCartHandler() *CartHandler {
	return &CartHandler{service: services.NewCartManipulator()}
}

type CartHandler struct {
	service services.CartService
}

func (c *CartHandler) Create(w http.ResponseWriter, _ *http.Request) {
	cart, err := c.service.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cartMarshalled, err := json.Marshal(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(cartMarshalled)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item models.CartItem
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&item)
	item.CartId = id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	created, err := c.service.AddItem(item)
	if errors.Is(err, services.ErrNoProductName) || errors.Is(err, services.ErrBadQuantity) || errors.Is(err, services.ErrNoSuchCart{Id: item.CartId}) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshaledCart, err := json.Marshal(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(marshaledCart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	itemId, err := strconv.Atoi(r.PathValue("item_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item := models.CartItem{Id: itemId, CartId: id}
	err = c.service.RemoveItem(item)
	if errors.Is(err, services.ErrNoSuchCart{Id: id}) || errors.Is(err, services.ErrNoSuchItem{Id: itemId}) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *CartHandler) View(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cart := models.Cart{Id: id}
	result, err := c.service.Get(cart.Id)
	if errors.Is(err, new(services.ErrNoSuchCart)) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshaled, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(marshaled)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
