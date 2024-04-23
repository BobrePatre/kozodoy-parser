package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/datasources"
	"log/slog"
	"os"
)

func NetworkDatacourceConfig(p *Provider) *datasources.NetworkConfig {
	if p.networkConfig == nil {
		cfg, err := datasources.NewNetworkConfig(p.Validate())
		if err != nil {
			slog.Error("Error creating network datasources.NetworkConfig", "error", err)
			os.Exit(1)
		}
		p.networkConfig = cfg
	}
	return p.networkConfig
}
