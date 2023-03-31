package main

import (
	"fmt"

	"go-restapi-mysql/app/Models"
	"go-restapi-mysql/app/Models/Activity"
	"go-restapi-mysql/app/Models/Todo"
	Config "go-restapi-mysql/config"
	api "go-restapi-mysql/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Book{})
	Config.DB.AutoMigrate(&Activity.Activity{})
	Config.DB.AutoMigrate(&Todo.Todo{})

	router := api.SetupRouter()

	//running
	router.Run(":8000")
}
