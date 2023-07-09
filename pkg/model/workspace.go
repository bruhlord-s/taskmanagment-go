package model

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
