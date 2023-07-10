package service

import (
	"github.com/bruhlord-s/openboard-go/internal/model"
	"github.com/bruhlord-s/openboard-go/internal/repository"
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

func (s *WorkspaceService) Update(userId, workspaceId int, input model.UpdateWorkspaceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, workspaceId, input)
}

func (s *WorkspaceService) Delete(userId, workspaceId int) error {
	return s.repo.Delete(userId, workspaceId)
}
