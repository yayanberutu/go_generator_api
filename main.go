package main

import (
	"crud-api/config"
	"crud-api/models"
	"crud-api/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Admin{})
	r := routes.SetupRouter()
	//running
	r.Run(Config.Port)
}
