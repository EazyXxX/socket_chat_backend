package repository

import (
	"fmt"
	"socket_chat_backend/types"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg types.DBConfig) (*sqlx.DB, error) {
	// Creating a database connection
	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	// Checking database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Returning a database connection instance for future use
	return db, nil
}
