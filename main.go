package main

import (
	"log"
	"tasking-rest-api/handler"
	"tasking-rest-api/task"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/tasking_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.GET("/tasks", taskHandler.GetTasks)
	api.GET("/task/:id", taskHandler.GetTask)
	api.POST("/tasks", taskHandler.CreateTask)
	api.PUT("/task/:id", taskHandler.UpdateData)
	api.PUT("/status/:id", taskHandler.UpdateDataStatus)

	router.Run()
}
