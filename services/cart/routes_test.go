package cart

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vnsonvo/ecom-rest-api/types"
)

var mockProducts = []types.Product{
	{ID: 1, Name: "product 1", Price: 11, Quantity: 200},
	{ID: 2, Name: "product 2", Price: 20, Quantity: 100},
	{ID: 4, Name: "out of stock", Price: 50, Quantity: 0},
	{ID: 5, Name: "last stock", Price: 30, Quantity: 1},
}

func TestCartServiceHandler(t *testing.T) {
	productStore := &mockProductStore{}
	orderStore := &mockOrderStore{}
	handler := NewHandler(orderStore, productStore, nil)

	t.Run("should fail to checkout if the cart items do not exist", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: 6, Quantity: 16},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /cart/checkout", handler.handleCheckout)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})

	t.Run("should fail to checkout if the cart has invalid quantity", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: 1, Quantity: 0},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /cart/checkout", handler.handleCheckout)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})

	t.Run("should fail to checkout if there is no stock for an item", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: 4, Quantity: 4},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /cart/checkout", handler.handleCheckout)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})

	t.Run("should fail to checkout if there is not enough stock", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: 5, Quantity: 5},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /cart/checkout", handler.handleCheckout)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProductByID(productID int) (*types.Product, error) {
	return &types.Product{}, nil
}

func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return []types.Product{}, nil
}

func (m *mockProductStore) CreateProduct(product types.CreateProductPayload) error {
	return nil
}

func (m *mockProductStore) UpdateProduct(product types.Product) error {
	return nil
}

func (m *mockProductStore) GetProductsById(ids []int) ([]types.Product, error) {
	return []types.Product{}, nil
}

type mockOrderStore struct{}

func (m *mockOrderStore) CreateOrder(order types.Order) (int, error) {
	return 0, nil
}

func (m *mockOrderStore) CreateOrderItem(orderItem types.OrderItem) error {
	return nil
}
