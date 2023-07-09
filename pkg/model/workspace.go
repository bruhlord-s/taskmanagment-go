package model

import "errors"

type Workspace struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required" db:"name"`
	Description string `json:"description" db:"description"`
	Avatar      string `json:"avatar" db:"avatar"`
}

type WorkspaceUser struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	WorkspaceId int `json:"workspace_id"`
	Role        int `json:"role"`
}

type WorkspaceInvite struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	WorkspaceId int `json:"workspace_id"`
}

type UpdateWorkspaceInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Avatar      *string `json:"avatar"`
}

func (i UpdateWorkspaceInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.Avatar == nil {
		return errors.New("update struct has no values")
	}

	return nil
}
