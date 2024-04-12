package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vnsonvo/ecom-rest-api/cmd/api"
	"github.com/vnsonvo/ecom-rest-api/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbConfig := os.Getenv("DBCONFIG")

	db, err := database.NewPostgresSQLStorage(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	server := api.NewAPIServer(port, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
