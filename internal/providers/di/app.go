package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/core"
	"log"
)

func (p *Provider) AppConfig() core.AppConfig {
	if p.appConfig == nil {
		cfg, err := core.NewAppConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get app config: %s", err.Error())
		}
		p.appConfig = cfg
	}

	return *p.appConfig
}
