package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vnsonvo/ecom-rest-api/cmd/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")

	server := api.NewAPIServer(port, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
