package security

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config"
	"github.com/go-playground/validator/v10"
	"time"
)

type CorsConfig struct {
	AllowAllOrigins           bool          `env:"ALLOW_ALL_ORIGINS" json:"allowAllOrigins" env-default:"true"`
	AllowOrigins              []string      `env:"ALLOW_ORIGINS" json:"allowOrigins" env-default:"*"`
	AllowMethods              []string      `env:"ALLOW_METHODS" json:"allowMethos" env-default:"*"`
	AllowHeaders              []string      `env:"ALLOW_HEADERS" json:"allowHeaders" env-default:"*"`
	AllowPrivateNetwork       bool          `env:"ALLOW_PRIVATE_NETWORK" json:"allowPrivateNetwork" env-default:"true"`
	AllowCredentials          bool          `env:"ALLOW_CREDENTIALS" json:"allowCredentials" env-default:"true"`
	ExposeHeaders             []string      `env:"EXPOSE_HEADERS" json:"exposeHeaders"`
	MaxAge                    time.Duration `env:"MAX_AGE" json:"maxAge" env-default:"12h"`
	AllowWildcard             bool          `env:"ALLOW_WILDCARD" json:"allowWildcard" env-default:"true"`
	AllowBrowserExtensions    bool          `env:"ALLOW_BROWSER_EXTENSIONS" json:"allowBrowserExtensions" env-default:"true"`
	CustomSchemas             []string      `env:"CUSTOM_SCHEMAS" json:"customSchemas" env-default:"*"`
	AllowWebSockets           bool          `env:"ALLOW_WEBSOCKETS" json:"allowWebSockets" env-default:"false"`
	AllowFiles                bool          `env:"ALLOW_FILES" json:"allowFiles" env-default:"false"`
	OptionsResponseStatusCode int           `env:"OPTIONS_RESPONSE_STATUS_CODE" json:"optionsResponseStatusCode" env-default:"200"`
}

func NewCorsConfig(validate *validator.Validate) (*CorsConfig, error) {
	var cfg struct {
		Config CorsConfig `env-prefix:"CORS_" json:"cors"`
	}
	if err := config.Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}
