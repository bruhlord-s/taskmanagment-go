package service

import (
	"github.com/bruhlord-s/openboard-go/pkg/model"
	"github.com/bruhlord-s/openboard-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
