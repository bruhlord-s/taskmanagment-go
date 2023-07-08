package model

type Board struct {
	Id          int    `json:"-"`
	WorkspaceId int    `json:"workspace_id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
}
