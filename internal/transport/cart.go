package transport

import (
	"net/http"

	"github.com/p1mako/cart-api/internal/services"
)

func NewCartHandler() *CartHandler {
	return &CartHandler{service: services.NewCartService()}
}

type CartHandler struct {
	service *services.CartService
}

func (c *CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	return
}
