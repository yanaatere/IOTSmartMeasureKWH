package main

import (
	"github.com/retere/IOTSmartMeasureKWH/config"
	"github.com/retere/IOTSmartMeasureKWH/models"
	"github.com/retere/IOTSmartMeasureKWH/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Println("Starting Service ...")

	config.LoadEnv()
	log.Println("Environment variables loaded ...")

	models.OpenDatabaseConnectionUsingURL()
	log.Println("Database connection established ...")

	models.AutoMigrateModels()
	log.Println("Database Models Migrated successfully")

	router := routes.SetupRoutes()
	log.Println("Routes Configured")

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}
	log.Println("Trusted proxies configured")

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down service ...")
	log.Println("Service Stopped")
}
