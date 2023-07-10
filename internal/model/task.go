package model

type Task struct {
	Id            int    `json:"-"`
	BoardId       int    `json:"board_id"`
	Number        int    `json:"number"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	TimeEstimated int    `json:"time_estimated"`
	AssigneeId    int    `json:"assignee_id"`
	AuthorId      int    `json:"author_id"`
}

type TaskAttachment struct {
	Id     int    `json:"-"`
	File   string `json:"file"`
	TaskId int    `json:"task_id"`
}

type TaskTimeTrack struct {
	Id        int `json:"-"`
	TaskId    int `json:"task_id"`
	UserId    int `json:"user_id"`
	TimeSpent int `json:"time_spent"`
}

type TaskComment struct {
	Id     int    `json:"-"`
	TaskId int    `json:"task_id"`
	UserId int    `json:"user_id"`
	Body   string `json:"body"`
}
