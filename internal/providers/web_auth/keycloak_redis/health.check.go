package keycloak_redis

import (
	"context"
	"github.com/BobrePatre/kozodoy-parser/internal/providers/web_auth"
	"github.com/lestrrat-go/jwx/jwk"
)

var _ webAuthProvider.Provider = (*Provider)(nil)

func (p *Provider) CheckSsoConnection(ctx context.Context) error {

	_, err := jwk.Fetch(ctx, p.jwkOpts.JwkPublicUri)
	if err != nil {
		return err
	}

	return nil
}
