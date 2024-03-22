package client

import (
	"github.com/go-pg/pg/v10"

	"reseller/config"
)

func NewPostgresClient(postgresqlConfig config.Postgres) *pg.DB {
	return pg.Connect(&pg.Options{
		User:     postgresqlConfig.DbUser,
		Password: postgresqlConfig.DbPassword,
		Database: postgresqlConfig.Database,
		Addr:     postgresqlConfig.Endpoint,
	})
	// TODO check it  out
	//defer db.Close()

}
