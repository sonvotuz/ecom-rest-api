package user

import (
	"fmt"
	"net/http"

	"github.com/vnsonvo/ecom-rest-api/types"
	"github.com/vnsonvo/ecom-rest-api/utils"
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
	// get payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
