package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func Connect(PSQL_USERNAME, PSQL_PASSWORD, PSQL_PORT, PSQL_HOSTNAME, PSQL_DATABASE string) {
	dsn := "postgres://" + PSQL_USERNAME + ":" + PSQL_PASSWORD + "@" + PSQL_HOSTNAME + ":" + PSQL_PORT + "/" + PSQL_DATABASE + "?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	//defer sqldb.Close()
	//defer context.Background().Done()

	DB = bun.NewDB(sqldb, pgdialect.New())

	// Create table if not exist
	_, err := DB.NewCreateTable().Model((*User)(nil)).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	_, err = DB.NewCreateTable().Model((*Todo)(nil)).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
