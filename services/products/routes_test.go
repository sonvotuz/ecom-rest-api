package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vnsonvo/ecom-rest-api/types"
)

func TestProductServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(productStore, userStore)

	t.Run("should handle get products", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("GET /products", handler.handleGetProducts)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
		}
	})

	t.Run("should fail if the product ID is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/abc", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("GET /products/{productId}", handler.handleGetProduct)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})

	t.Run("should handle get product by ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("GET /products/{productId}", handler.handleGetProduct)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
		}
	})

	t.Run("should fail creating a product if the payload is missing", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /products", handler.handlerCreateProduct)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})

	t.Run("should handle creating a product", func(t *testing.T) {
		payload := types.CreateProductPayload{
			Name:        "test",
			Price:       100,
			Image:       "test.jpg",
			Description: "test description",
			Quantity:    10,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /products", handler.handlerCreateProduct)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, responseRecorder.Code)
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

type mockUserStore struct{}

func (m *mockUserStore) GetUserByID(userID int) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}
