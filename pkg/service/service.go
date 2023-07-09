package service

import (
	"github.com/bruhlord-s/openboard-go/pkg/model"
	"github.com/bruhlord-s/openboard-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Workspace interface {
	Create(userId int, workspace model.Workspace) (int, error)
	GetAll(userId int) ([]model.Workspace, error)
	GetById(userId, workspaceId int) (model.Workspace, error)
	Update(userId, workspaceId int, input model.UpdateWorkspaceInput) error
	Delete(userId, workspaceid int) error
}

type Service struct {
	Authorization
	Workspace
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Workspace:     NewWorkspaceService(repos.Workspace),
	}
}
