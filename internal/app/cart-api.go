package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/p1mako/cart-api/internal/database"
	"github.com/p1mako/cart-api/internal/transport"
)

func Run() {
	database.InitDb()
	defer closeDB()
	setPaths()
	exitCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	srv := http.Server{Addr: ":3000"}
	go startServer(&srv, stop)
	defer stopServer(&srv)
	<-exitCtx.Done()
}

func stopServer(srv *http.Server) {
	err := srv.Close()
	if err != nil {
		log.Fatalf("unable to stop server: %v\n", err.Error())
	}
	log.Printf("Sucessfully closed server: %v\n", srv.Addr)

}

func startServer(srv *http.Server, c context.CancelFunc) {
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Cannot start server :%v\n", err.Error())
		c()
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
