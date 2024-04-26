package delivery

import (
	"fmt"
	"github.com/BobrePatre/kozodoy-parser/internal/config"
	"github.com/go-playground/validator/v10"
	"net"
)

type HttpConfig struct {
	Port int    `json:"port" env:"PORT" env-default:"8080"`
	Host string `json:"host" env:"HOST" env-default:"localhost"`
}

func NewHTTPConfig(validate *validator.Validate) (*HttpConfig, error) {

	var cfg struct {
		Config HttpConfig `env-prefix:"HTTP_" json:"http"`
	}
	if err := config.Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}

func (cfg *HttpConfig) Address() string {
	return net.JoinHostPort(cfg.Host, fmt.Sprint(cfg.Port))
}
