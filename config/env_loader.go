package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("D:\\Private Study\\IOTSmartMeasureKWH\\config\\.env")
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err)) // Include the actual error message
	} else {
		fmt.Println("Environment variables loaded successfully.")
	}
}
