package main

import (
	"github.com/retere/IOTSmartMeasureKWH/config"
	"github.com/retere/IOTSmartMeasureKWH/models"
	"github.com/retere/IOTSmartMeasureKWH/routes"
)

func main() {
	config.LoadEnv()
	models.OpenDatabaseConnectionUsingURL()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	router.Run(":8080")
}
