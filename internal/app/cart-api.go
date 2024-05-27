package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/transport"
)

func Run() {
	defer func() {
		err := database.Close()
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "cannot close database connection, beacause: %v", err)
			if err != nil {
				panic("Unable to log errors")
			}
		}
	}()
	cartService := transport.NewCartHandler()
	http.HandleFunc("POST /cart", cartService.Create)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		return
	}
}
