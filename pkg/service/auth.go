package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/bruhlord-s/openboard-go/pkg/model"
	"github.com/bruhlord-s/openboard-go/pkg/repository"
)

const salt = "sidiuahnduiwu"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePassordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePassordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
