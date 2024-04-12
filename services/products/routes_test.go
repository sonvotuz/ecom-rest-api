package products

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
