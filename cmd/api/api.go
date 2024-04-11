package api

import (
	"database/sql"
	"log"
	"net/http"
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

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	var server = &http.Server{
		Addr:    ":" + s.addr,
		Handler: mux,
	}

	log.Println("Listening on", s.addr)

	return server.ListenAndServe()
}
