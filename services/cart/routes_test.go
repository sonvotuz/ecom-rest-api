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
