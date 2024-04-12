package user

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, prefixPath string) {
	mux.HandleFunc(fmt.Sprintf("POST %s/login", prefixPath), h.handleLogin)
	mux.HandleFunc(fmt.Sprintf("POST %s/register", prefixPath), h.handleRegister)

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
