package transport

import (
	"encoding/json"
	"net/http"

	"github.com/p1mako/cart-api/internal/services"
)

func NewCartHandler() *CartHandler {
	return &CartHandler{service: services.NewCartService()}
}

type CartHandler struct {
	service *services.CartService
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
