package datasources

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config"
	"github.com/go-playground/validator/v10"
)

type NetworkConfig struct {
	Host            string `env:"HOST" json:"host" validate:"required"`
	Port            int    `env:"PORT" json:"port" env-default:"5432"`
	User            string `env:"USER" json:"user" env-default:"postgres"`
	Password        string `env:"PASSWORD" json:"password" env-default:"postgres"`
	Database        string `env:"DATABASE" json:"database" env-default:"postgres"`
	CoreBackendHost string `env:"CORE_BACKEND_HOST" json:"coreBackendHost" env-default:"localhost:2000"`
}

func NewNetworkConfig(validate *validator.Validate) (*NetworkConfig, error) {
	var cfg struct {
		Config NetworkConfig `json:"network" env-prefix:"NETWORK_"`
	}
	if err := config.Load(&cfg, validate); err != nil {
		return nil, err
	}
	return &cfg.Config, nil
}
