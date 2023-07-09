package repository

import (
	"fmt"

	"github.com/bruhlord-s/openboard-go/pkg/model"
	"github.com/jmoiron/sqlx"
)

type WorkspacePostgres struct {
	db *sqlx.DB
}

func NewWorkspacePostgres(db *sqlx.DB) *WorkspacePostgres {
	return &WorkspacePostgres{db: db}
}

func (r *WorkspacePostgres) Create(userId int, workspace model.Workspace) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createWorkspaceQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", workspacesTable)
	row := tx.QueryRow(createWorkspaceQuery, workspace.Name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createWorkspaceUsersQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, workspace_id, role) VALUES ($1, $2, $3)",
		workspaceUsersTable,
	)
	_, err = tx.Exec(createWorkspaceUsersQuery, userId, id, 1)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *WorkspacePostgres) GetAll(userId int) ([]model.Workspace, error) {
	var workspaces []model.Workspace

	query := fmt.Sprintf(
		"SELECT w.* FROM %s w INNER JOIN %s wu ON wu.workspace_id = w.id WHERE wu.user_id = $1",
		workspacesTable,
		workspaceUsersTable,
	)
	err := r.db.Select(&workspaces, query, userId)

	return workspaces, err
}

func (r *WorkspacePostgres) GetById(userId, workspaceId int) (model.Workspace, error) {
	var workspace model.Workspace

	query := fmt.Sprintf(
		`SELECT w.* FROM %s w INNER JOIN %s wu ON wu.workspace_id = w.id
			WHERE wu.user_id = $1 AND w.id = $2`,
		workspacesTable,
		workspaceUsersTable,
	)
	err := r.db.Get(&workspace, query, userId, workspaceId)

	return workspace, err
}
