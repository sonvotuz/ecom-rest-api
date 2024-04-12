package api

import (
	"database/sql"
	"log"
	"net/http"

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

	var server = &http.Server{
		Addr:    ":" + s.addr,
		Handler: mux,
	}

	log.Println("Listening on", s.addr)

	return server.ListenAndServe()
}
