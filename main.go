package main

import (
	"fmt"
	"log"
	"tasking-rest-api/task"

	"github.com/google/uuid"
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
	id := uuid.Must(uuid.MustParse("a15601a8-e41d-4281-969a-f5c991c44346"), err)
	tasks, err := taskRepository.FindByTaskID(id)

	fmt.Println("debug")
	for _, task := range tasks {
		fmt.Println(task.Assign)
	}

	// router := gin.Default()
	// api := router.Group("/api/v1")

	// api.GET("/", taskHandler.)
	// router.Run()
}
