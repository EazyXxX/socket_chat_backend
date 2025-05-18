package service

import (
	"socket_chat_backend/internal/repository"
	"socket_chat_backend/types"
)

type Authorization interface {
	CreateUser(user types.NewUserData) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
