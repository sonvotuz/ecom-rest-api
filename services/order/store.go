package order

import (
	"database/sql"

	"github.com/vnsonvo/ecom-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	lastInsertId := 0
	err := s.db.QueryRow("INSERT INTO orders (user_id, total, status, address) VALUES ($1, $2, $3, $4) RETURNING id", order.UserID, order.Total, order.Status, order.Address).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, price, quantity) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Price, orderItem.Quantity)

	return err
}
