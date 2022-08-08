package main

import (
	"log"
	"net/http"
	"tasking-rest-api/task"

	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:root@tcp(127.0.0.1:3306)/tasking_app?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection to database is good.")

	// var tasks []task.Task

	// db.Find(&tasks)

	// length := len(tasks)

	// fmt.Println(length)

	// for _, task := range tasks {
	// 	fmt.Println(task.Assign)
	// 	fmt.Println(task.Deadline)
	// }

	router := gin.Default()
	router.GET("handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:root@tcp(127.0.0.1:3306)/tasking_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var tasks []task.Task

	db.Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}