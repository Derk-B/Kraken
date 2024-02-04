package main

import (
	"fmt"
	"kraken/api-server/models"
	"kraken/api-server/server"
	"os"
	"time"
)

func main() {

	PSQL_HOSTNAME := os.Getenv("POSTGRES_HOSTNAME")
	PSQL_PORT := os.Getenv("POSTGRES_PORT")
	PSQL_USERNAME := os.Getenv("POSTGRES_USER")
	PSQL_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	PSQL_DATABASE := os.Getenv("POSTGRES_DB")

	for {
		err := models.Connect(PSQL_USERNAME, PSQL_PASSWORD, PSQL_PORT, PSQL_HOSTNAME, PSQL_DATABASE)
		if err != nil {
			time.Sleep(1 * time.Second)
			fmt.Print("Waiting for all components to be ready...")
			continue
		} else {
			break
		}
	}

	s := server.New()
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
