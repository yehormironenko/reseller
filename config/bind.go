package config

import (
	_ "embed"
	"log"

	"github.com/knadh/koanf"
)

//go:embed app-config.yaml
var Yaml string

type Config struct {
	Postgres Postgres
	Server   Server
}

func Bind(kfg *koanf.Koanf) Config {
	appConfig := Config{}
	for path, field := range map[string]interface{}{
		"app.postgres": &appConfig.Postgres,
		"app.server":   &appConfig.Server,
	} {
		if err := kfg.Unmarshal(path, field); err != nil {
			log.Panic("config binding failed")
		}
	}

	if err := appConfig.Postgres.Validate(); err != nil {
		panic(err)
	}

	return appConfig
}
