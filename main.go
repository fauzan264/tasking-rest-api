package main

import (
	"fmt"
	"log"
	"tasking-rest-api/task"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/tasking_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good.")

	var tasks []task.Task

	db.Find(&tasks)

	length := len(tasks)

	fmt.Println(length)

	for _, task := range tasks {
		fmt.Println(task.Assign)
		fmt.Println(task.Deadline)
	}
}