package products

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/vnsonvo/ecom-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (types.Product, error) {
	product := types.Product{}
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	_, err := s.db.Exec("INSERT INTO products (name,description, price, image, quantity) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Price, product.Image, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetProductsById(productIds []int) ([]types.Product, error) {
	productIdStrs := make([]string, len(productIds))
	for i, v := range productIds {
		productIdStrs[i] = fmt.Sprintf("%v", v)
	}

	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", strings.Join(productIdStrs, ", "))

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = $1, description = $2, image = $3, price = $4, quantity = $5 WHERE id = $6",
		product.Name, product.Description, product.Image, product.Price, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}
