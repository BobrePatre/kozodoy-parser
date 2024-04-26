package delivery

import (
	"fmt"
	"github.com/BobrePatre/kozodoy-parser/internal/config"
	"github.com/go-playground/validator/v10"
	"net"
)

type GrpcConfig struct {
	Port string `env:"PORT" env-default:"50051" json:"grpc"`
	Host string `env:"HOST" env-default:"localhost" json:"host"`
}

func NewGrpcConfig(validate *validator.Validate) (*GrpcConfig, error) {
	var cfg struct {
		Config GrpcConfig `env-prefix:"GRPC_" json:"grpc"`
	}
	if err := config.Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}

func (cfg *GrpcConfig) Address() string {
	return net.JoinHostPort("localhost", fmt.Sprint(cfg.Port))
}
