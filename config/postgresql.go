package config

import (
	"fmt"
	"strings"
)

type Postgres struct {
	Endpoint   string `koanf:"endpoint"`
	Database   string `koanf:"database"`
	Tables     Tables `koanf:"tables"`
	DbUser     string `koanf:"user"`
	DbPassword string `koanf:"password"`
}

type Tables struct {
	Books string `koanf:"books"`
}

func (d Postgres) Validate() error {
	var missing []string
	if d.Tables.Books == "" {
		missing = append(missing, "users")
	}
	if d.Database == "" {
		missing = append(missing, "database")
	}
	if d.DbUser == "" {
		missing = append(missing, "user")
	}
	if d.DbPassword == "" {
		missing = append(missing, "password")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing PostgreSQL config: %s", strings.Join(missing, ", "))
	}
	return nil
}
