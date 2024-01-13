package main

import (
	"context"
	"database/sql"
	"fmt"
	api "kraken/api-server/endpoint"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

	dsn := "postgres://" + PSQL_USERNAME + ":" + PSQL_PASSWORD + "@localhost:" + PSQL_PORT + "/" + PSQL_DATABASE + "?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	defer sqldb.Close()
	defer context.Background().Done()

	db := bun.NewDB(sqldb, pgdialect.New())

	// Run router
	r := gin.Default()

	api.Run(r, db)

	err = r.Run("0.0.0.0:8888")
	if err != nil {
		return
	}
}
