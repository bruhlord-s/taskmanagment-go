package repository

import (
	"github.com/bruhlord-s/openboard-go/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Workspace interface {
	Create(userId int, workspace model.Workspace) (int, error)
	GetAll(userId int) ([]model.Workspace, error)
	GetById(userId, workspaceId int) (model.Workspace, error)
	Update(userId, workspaceId int, input model.UpdateWorkspaceInput) error
	Delete(userId, workspaceId int) error
}

type Repository struct {
	Authorization
	Workspace
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Workspace:     NewWorkspacePostgres(db),
	}
}
