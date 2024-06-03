package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/p1mako/cart-api/internal/models"
	"github.com/p1mako/cart-api/internal/services"
)

func NewCartHandler() *CartHandler {
	return &CartHandler{service: services.NewCartManipulator()}
}

type CartHandler struct {
	service services.CartService
}

func (c *CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	cart, err := c.service.Create(ctx)
	cancel()
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
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	created, err := c.service.AddItem(ctx, item)
	cancel()
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
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	err = c.service.RemoveItem(ctx, item)
	cancel()
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
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	result, err := c.service.Get(ctx, cart.Id)
	cancel()
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
