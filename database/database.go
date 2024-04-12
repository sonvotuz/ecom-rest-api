package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresSQLStorage(cfg string) (*sql.DB, error) {
	connection, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, fmt.Errorf("Can't connect to database: %s", err.Error())
	}

	return connection, nil
}
