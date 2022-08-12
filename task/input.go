package task

type CreateTaskInput struct {
	Task     string `json:"task" binding:"required"`
	Assign   string `json:"assign" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
}
