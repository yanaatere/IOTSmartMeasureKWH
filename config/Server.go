package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	} else {
		fmt.Println("We are getting the env values")
	}

}
