package task

type GetTaskDetailInput struct {
	Id string `uri:"id" binding:"required"`
}

type CreateTaskInput struct {
	Task     string `json:"task" binding:"required"`
	Assign   string `json:"assign" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
}

type CreateTaskStatusInput struct {
	Status int `json:"status" binding:"required"`
}
