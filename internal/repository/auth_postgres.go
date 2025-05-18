package repository

import (
	"fmt"
	"socket_chat_backend/types"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user types.NewUserData) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (email, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Username, user.Password)
	// Simple ros.Scan to write an id into the declared variable
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (types.User, error) {
	var user types.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	// db.Get to write recieved values straight into the 'user' structure
	err := r.db.Get(&user, query, username, password)

	return user, err
}
