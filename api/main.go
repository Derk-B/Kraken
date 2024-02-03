package main

import (
	"kraken/api-server/models"
	"kraken/api-server/server"
	"os"
)

func main() {
	// Get env variables for connection string
	//envFile, err := godotenv.Read("../.env")
	//
	//if err != nil {
	//	fmt.Println("Could not load env file")
	//}
	//PSQL_PORT := envFile["PSQL_PORT"]
	//PSQL_USERNAME := envFile["PSQL_USERNAME"]
	//PSQL_PASSWORD := envFile["PSQL_PASSWORD"]
	//PSQL_DATABASE := envFile["PSQL_DATABASE"]
	//PSQL_HOSTNAME := envFile["PSQL_HOSTNAME"]

	PSQL_HOSTNAME := os.Getenv("POSTGRES_HOSTNAME")
	PSQL_PORT := os.Getenv("POSTGRES_PORT")
	PSQL_USERNAME := os.Getenv("POSTGRES_USER")
	PSQL_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	PSQL_DATABASE := os.Getenv("POSTGRES_DB")

	models.Connect(PSQL_USERNAME, PSQL_PASSWORD, PSQL_PORT, PSQL_HOSTNAME, PSQL_DATABASE)
	s := server.New()
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
