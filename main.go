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

	// input := task.CreateTaskInput{}
	// input.Task = "Belajar Python programming"
	// input.Assign = "Ahmad"

	// deadlineTask, _ := time.Parse("2006-01-02", "2022-08-07")
	// input.Deadline = deadlineTask

	// fmt.Println(input)
	// _, err = taskService.CreateTask(input)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.GET("/tasks", taskHandler.GetTasks)
	api.POST("/tasks", taskHandler.CreateTask)
	router.Run()
}
