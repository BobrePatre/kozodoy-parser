package diProvider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"log"
)

func (p *DiProvider) CorsConfig() *config.CorsConfig {

	if p.corsConfig == nil {
		cfg, err := config.NewCorsConfig(p.Validate())
		if err != nil {
			log.Fatalf("cannot read cors config, err = %s", err)
		}
		p.corsConfig = cfg
	}

	return p.corsConfig
}
