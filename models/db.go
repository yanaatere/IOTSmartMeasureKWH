package models

import (
	"fmt"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"gorm.io/gorm"
)

var Database *gorm.DB

func OpenDatabaseConnection() {
	var err error
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Africa/Douala", host, username, password, databaseName, port)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}
}

func OpenDatabaseConnectionUsingURL() {
	var err error

	// Fetch the connection URL from the environment variables
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal("POSTGRES_URL environment variable is not set")
	}

	// Open the database connection using the URL directly
	Database, err = gorm.Open(postgres.Open(postgresURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}
}
func AutoMigrateModels() {
	Database.AutoMigrate(&tenantentity.Tenant{})
}
