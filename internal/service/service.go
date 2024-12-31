package service

import (
	"socket_chat_backend/internal/repository"
	"socket_chat_backend/types"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
