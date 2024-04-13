package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/delivery"
	"log"
)

func (p *Provider) HTTPConfig() *delivery.HttpConfig {
	if p.httpConfig == nil {
		cfg, err := delivery.NewHTTPConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get delivery config: %s", err.Error())
		}
		p.httpConfig = cfg
	}

	return p.httpConfig
}
