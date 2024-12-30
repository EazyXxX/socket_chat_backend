package repository

import (
	"fmt"

	"socket_chat_backend/types"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

func NewPostgresDB(cfg types.DBConfig) (*sqlx.DB, error) {
	// Creating a database connection
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)

	// Connection opening using sqlx
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logrus.WithError(err).Error("Could not connect to the database")
	}

	logrus.Info("Successfully connected to the database")
	return db, nil
}
