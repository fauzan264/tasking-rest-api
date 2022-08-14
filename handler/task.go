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

	if len(id) != 0 {
		_, err := uuid.Parse(id)
		if err != nil {
			response := helper.APIResponse("Error to get tasks, ID not valid", http.StatusBadRequest, "error", task.FormatTasks(nil))
			c.JSON(http.StatusBadRequest, response)
			return
		}
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

func (h *taskHandler) GetTask(c *gin.Context) {
	var input task.GetTaskDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	taskDetail, err := h.service.GetTaskById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of task not id", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Task detail", http.StatusOK, "success", task.FormatDetailTask(taskDetail))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) CreateTask(c *gin.Context) {
	var input task.CreateTaskInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newtask, err := h.service.CreateTask(input)
	if err != nil {
		response := helper.APIResponse("Failed to create task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create task", http.StatusOK, "success", task.FormatTask(newtask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdateData(c *gin.Context) {
	var inputId task.GetTaskDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse("Failed to update task, but id not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData task.CreateTaskInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorData := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errorData}

		response := helper.APIResponse("Failed to update task, wrong data", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTask, err := h.service.UpdateData(inputId, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update task", http.StatusOK, "success", task.FormatTask(updateTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdateDataStatus(c *gin.Context) {
	var inputId task.GetTaskDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse("Failed to update task status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData task.CreateTaskStatusInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorData := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errorData}

		response := helper.APIResponse("Failed to update task status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTask, err := h.service.UpdateDataStatus(inputId, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update task status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update task", http.StatusOK, "success", task.FormatTask(updateTask))
	c.JSON(http.StatusOK, response)
}
