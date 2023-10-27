package main

import (
	"github.com/gin-gonic/gin"
	"poc-push/config"
	"poc-push/entities"
	"poc-push/router"
)

func main() {
	config.OpenDatabaseConnection()
	defer config.CloseDatabaseConnection()
	config.DB.AutoMigrate(&entities.Subscription{})

	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	router.Routes(app)
	app.Run()
}
