package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/vnsonvo/ecom-rest-api/services/cart"
	"github.com/vnsonvo/ecom-rest-api/services/order"
	"github.com/vnsonvo/ecom-rest-api/services/products"
	"github.com/vnsonvo/ecom-rest-api/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

const prefixPath = "/api/v1"

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	userStore := user.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(mux, prefixPath)

	productStore := products.NewStore(s.db)
	productHandler := products.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(mux, prefixPath)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(mux, prefixPath)

	var server = &http.Server{
		Addr:    ":" + s.addr,
		Handler: mux,
	}

	log.Println("Listening on", s.addr)

	return server.ListenAndServe()
}
