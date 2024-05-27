package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/transport"
)

func Run() {
	defer closeDB()
	setPaths()
	exit := make(chan error)
	srv := http.Server{Addr: ":3000"}
	go startServer(&srv, exit)
	defer func(srv *http.Server) {
		err := srv.Close()
		if err != nil {
			_ = fmt.Errorf("unable to stop server: %w\n", err)
			return
		}
		fmt.Printf("Sucessfully closed server: %v\n", srv.Addr)
	}(&srv)
	go userInput(exit)
	select {
	case err := <-exit:
		if err != nil {
			_ = fmt.Errorf("unable to start server: %w", err)
		}
		return
	}
}

func userInput(c chan error) {
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			_ = fmt.Errorf("failed to start listening to user input: %w\n", err)
			return
		}
		if strings.ToLower(input) == "exit" {
			c <- nil
			return
		}
	}
}

func startServer(srv *http.Server, c chan error) {
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		c <- err
	}
}

func setPaths() {
	cartService := transport.NewCartHandler()
	http.HandleFunc("POST /carts", cartService.Create)
	http.HandleFunc("POST /carts/{id}/items", cartService.AddItem)
	http.HandleFunc("DELETE /carts/{id}/items/{item_id}", cartService.RemoveItem)
	http.HandleFunc("GET /carts/{id}", cartService.View)
}

func closeDB() {
	err := database.Close()
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "cannot close database connection, beacause: %v", err)
		if err != nil {
			panic("unable to log errors")
		}
	}
	fmt.Print("Closed bd")
}
