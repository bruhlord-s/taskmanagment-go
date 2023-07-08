package model

type Workspace struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
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
