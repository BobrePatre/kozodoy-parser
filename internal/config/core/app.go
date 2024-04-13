package core

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config"
	"github.com/go-playground/validator/v10"
)

type AppConfig struct {
	MODE string `env:"MODE" json:"mode" validate:"required"`
}

func NewAppConfig(validate *validator.Validate) (*AppConfig, error) {
	var cfg struct {
		Config AppConfig `env-prefix:"APP_" json:"app"`
	}
	if err := config.Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}
