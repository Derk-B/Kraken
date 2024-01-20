package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"kraken/api-server/models"
	"kraken/api-server/server"
)

func main() {
	// Get env variables for connection string
	envFile, err := godotenv.Read("../.env")

	if err != nil {
		fmt.Println("Could not load env file")
	}

	// Create database connection
	PSQL_PORT := envFile["PSQL_PORT"]
	PSQL_USERNAME := envFile["PSQL_USERNAME"]
	PSQL_PASSWORD := envFile["PSQL_PASSWORD"]
	PSQL_DATABASE := envFile["PSQL_DATABASE"]

	models.Connect(PSQL_USERNAME, PSQL_PASSWORD, PSQL_PORT, PSQL_DATABASE)
	server := endpoint.New()
	server.Run()
}
