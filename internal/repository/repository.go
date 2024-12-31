package repository

import (
	"socket_chat_backend/types"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
	GetUser(username, password string) (types.User, error)
}

type Repository struct {
	Authorization
}

// Dependency injection. Needed for untying interface realisation from it's call in main.go
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
