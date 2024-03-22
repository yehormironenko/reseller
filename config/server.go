package config

import "time"

type Server struct {
	Endpoint        string        `koanf:"endpoint"`
	ShutdownTimeout time.Duration `koanf:"shutdownTimeout"`
}
