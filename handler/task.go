package handler

import (
	"net/http"
	"tasking-rest-api/helper"
	"tasking-rest-api/task"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type taskHandler struct {
	service task.Service
}

func NewTaskHandler(service task.Service) *taskHandler {
	return &taskHandler{service}
}

func (h *taskHandler) GetTasks(c *gin.Context) {
	id := c.Query("id")

	_, err := uuid.Parse(id)
	if err != nil {
		response := helper.APIResponse("Error to get tasks, ID not valid", http.StatusBadRequest, "error", task.FormatTasks(nil))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tasks, err := h.service.GetTasks(id)

	if err != nil {
		response := helper.APIResponse("Error to get tasks", http.StatusBadRequest, "error", task.FormatTasks(nil))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of tasks", http.StatusOK, "success", task.FormatTasks(tasks))
	c.JSON(http.StatusOK, response)
}
