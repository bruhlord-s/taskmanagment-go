package service

import (
	"github.com/bruhlord-s/openboard-go/pkg/model"
	"github.com/bruhlord-s/openboard-go/pkg/repository"
)

type WorkspaceService struct {
	repo repository.Workspace
}

func NewWorkspaceService(repo repository.Workspace) *WorkspaceService {
	return &WorkspaceService{repo: repo}
}

func (s *WorkspaceService) Create(userId int, workspace model.Workspace) (int, error) {
	return s.repo.Create(userId, workspace)
}

func (s *WorkspaceService) GetAll(userId int) ([]model.Workspace, error) {
	return s.repo.GetAll(userId)
}

func (s *WorkspaceService) GetById(userId, workspaceId int) (model.Workspace, error) {
	return s.repo.GetById(userId, workspaceId)
}
