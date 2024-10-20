package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
)

type Service struct {
	Home        *src.Home
	Transaction *route.TransactionHandler
}

type Serve struct {
	DB      *gorm.DB
	Router  *mux.Router
	Service Service
}

func (s *Serve) initialize(DBDriver, DBTransaction, DBPort, DBHost, DBName string) {
	var err error
	//Set migration table
	registry := app.SetMigrationTable()

	if DBDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBTransaction, DBTransaction, DBName)
		s.DB, err = gorm.Open(DBDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DBDriver)
		}
	}

	s.DB.Debug().AutoMigrate(registry...)

	s.Router = mux.NewRouter()

	s.initializeRoutes()
}
