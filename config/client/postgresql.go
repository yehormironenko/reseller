package client

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/yehormironenko/reseller/config"
)

func NewPostgresClient(postgresqlConfig config.Postgres) *bun.DB {
	sqldb := sql.OpenDB(createConnector(postgresqlConfig))

	return bun.NewDB(sqldb, pgdialect.New())
}

func createConnector(postgresqlConfig config.Postgres) *pgdriver.Connector {
	return pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(postgresqlConfig.Endpoint),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser(postgresqlConfig.DbUser),
		pgdriver.WithPassword(postgresqlConfig.DbPassword),
		pgdriver.WithDatabase(postgresqlConfig.Database),
		pgdriver.WithApplicationName("reseller"),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)
}
