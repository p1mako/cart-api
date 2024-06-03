package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
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
			log.Fatalf("unable to stop server: %v\n", err.Error())
		}
		log.Printf("Sucessfully closed server: %v\n", srv.Addr)
	}(&srv)
	go userInput(exit)
	select {
	case err := <-exit:
		if err != nil {
			log.Fatalf("unable to start server: %v\n", err.Error())
		}
		return
	}
}

func userInput(c chan error) {
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatalf("failed to start listening to user input: %v\n", err.Error())
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
		log.Printf("cannot close database connection, beacause: %v\n", err)
	}
	log.Println("Closed bd")
}
