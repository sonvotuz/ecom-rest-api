package products

import (
	"fmt"
	"net/http"

	"github.com/vnsonvo/ecom-rest-api/types"
	"github.com/vnsonvo/ecom-rest-api/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, prefixPath string) {
	mux.HandleFunc(fmt.Sprintf("GET %s/products", prefixPath), h.handleGetProducts)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}
