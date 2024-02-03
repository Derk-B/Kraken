package models

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func Connect(PSQL_USERNAME, PSQL_PASSWORD, PSQL_PORT, PSQL_HOSTNAME, PSQL_DATABASE string) error {
	dsn := "postgres://" + PSQL_USERNAME + ":" + PSQL_PASSWORD + "@" + PSQL_HOSTNAME + ":" + PSQL_PORT + "/" + PSQL_DATABASE + "?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqldb, pgdialect.New())

	// Create table if not exist
	_, err := DB.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}
	_, err = DB.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
