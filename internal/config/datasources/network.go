package datasources

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config"
	"github.com/go-playground/validator/v10"
)

type NetworkConfig struct {
	CoreBackendHost string `env:"CORE_BACKEND_HOST" json:"coreBackendHost" env-default:"http://localhost:2000"`
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
