package main

import (
	"fmt"
	"kraken/api-server/models"
	"kraken/api-server/server"

	"github.com/joho/godotenv"
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
	PSQL_HOSTNAME := envFile["PSQL_HOSTNAME"]

	models.Connect(PSQL_USERNAME, PSQL_PASSWORD, PSQL_PORT, PSQL_HOSTNAME, PSQL_DATABASE)
	s := server.New()
	s.Run()
}
