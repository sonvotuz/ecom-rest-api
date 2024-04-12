package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vnsonvo/ecom-rest-api/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "test",
			Email:     "invalid",
			Password:  "123456",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /register", handler.handleRegister)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("Unexpected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "test",
			Email:     "user@gmail.com",
			Password:  "123456",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		mux := http.NewServeMux()

		mux.HandleFunc("POST /register", handler.handleRegister)
		mux.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusCreated {
			t.Errorf("Unexpected status code %d, got %d", http.StatusCreated, responseRecorder.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
