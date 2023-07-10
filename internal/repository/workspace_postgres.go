package repository

import (
	"fmt"
	"strings"

	"github.com/bruhlord-s/openboard-go/internal/model"
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

func (r *WorkspacePostgres) Update(userId, workspaceId int, input model.UpdateWorkspaceInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *&input.Name)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *&input.Description)
		argId++
	}

	if input.Avatar != nil {
		setValues = append(setValues, fmt.Sprintf("avatar=$%d", argId))
		args = append(args, *&input.Avatar)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s w SET %s FROM %s wu 
			WHERE w.id = wu.workspace_id AND wu.workspace_id = $%d AND wu.user_id = $%d`,
		workspacesTable,
		setQuery,
		workspaceUsersTable,
		argId,
		argId+1,
	)
	args = append(args, workspaceId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *WorkspacePostgres) Delete(userId, workspaceId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s w USING %s wu 
			WHERE w.id = wu.workspace_id AND wu.user_id = $1 AND wu.workspace_id = $2`,
		workspacesTable,
		workspaceUsersTable,
	)
	_, err := r.db.Exec(query, userId, workspaceId)

	return err
}
