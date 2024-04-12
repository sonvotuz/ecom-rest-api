package user

import (
	"database/sql"
	"fmt"

	"github.com/vnsonvo/ecom-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	user := &types.User{}
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := types.User{}
	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", id)
	if err != nil {
		return nil, err
	}
	user := &types.User{}
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
