package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/providers/service_access"
	keycloakToken "github.com/BobrePatre/kozodoy-parser/internal/providers/service_access/keycloak_token"
)

func (p *Provider) ServiceAccessProvider() service_access.Provider {
	if p.servieAccessProvider == nil {
		p.servieAccessProvider = keycloakToken.NewProvider(p.AuthConfig())
	}
	return p.servieAccessProvider
}
