package app

import (
	"net/http"

	"github.com/p1mako/cart-api/internal/transport"
)

func Run() {
	cartService := transport.NewCartHandler()
	http.HandleFunc("POST /cart", cartService.Create)
	err := http.ListenAndServe("localhost:3000", http.DefaultServeMux)
	if err != nil {
		return
	}
}
