package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/security"
	"log"
)

func (p *Provider) CorsConfig() *security.CorsConfig {

	if p.corsConfig == nil {
		cfg, err := security.NewCorsConfig(p.Validate())
		if err != nil {
			log.Fatalf("cannot read cors config, err = %s", err)
		}
		p.corsConfig = cfg
	}

	return p.corsConfig
}
