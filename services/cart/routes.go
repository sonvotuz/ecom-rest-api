package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vnsonvo/ecom-rest-api/services/auth"
	"github.com/vnsonvo/ecom-rest-api/types"
	"github.com/vnsonvo/ecom-rest-api/utils"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, productStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, prefixPath string) {
	mux.HandleFunc(fmt.Sprintf("POST %s/cart/checkout", prefixPath), auth.JWTAuthMiddleware(h.handleCheckout, h.userStore))
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid payload: %v", errors))
		return
	}

	// get products
	productIds, err := getCartItemIds(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	products, err := h.productStore.GetProductsById(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	userId := auth.GetUserIDFromContext(r.Context())
	orderId, totalPrice, err := h.CreateOrderFromCart(products, cart.Items, userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"order_id":    orderId,
		"total_price": totalPrice,
	})
}
