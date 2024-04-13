package diProvider

import (
	"context"
	"github.com/BobrePatre/kozodoy-parser/internal/config/security"
	httpMiddlewares "github.com/BobrePatre/kozodoy-parser/internal/delivery/http/middlewares"
	webAuthProvider "github.com/BobrePatre/kozodoy-parser/internal/providers/web_auth"
	keycloakAuthProvider "github.com/BobrePatre/kozodoy-parser/internal/providers/web_auth/keycloak_redis"
	"log"
	"log/slog"
	"os"
	"time"
)

func (p *Provider) HttpAuthMiddlewareConstructor() webAuthProvider.HttpMiddlewareConstructor {
	if p.httpAuthMiddlewareConstructor == nil {
		p.httpAuthMiddlewareConstructor = httpMiddlewares.AuthMiddleware
	}
	return p.httpAuthMiddlewareConstructor

}

func (p *Provider) WebAuthProvider() webAuthProvider.Provider {
	if p.webAuthProvider == nil {
		p.webAuthProvider = keycloakAuthProvider.NewProvider(p.RedisClient(), webAuthProvider.JwkOptions{
			JwkPublicUri:      p.AuthConfig().PublicJwkUri,
			RefreshJwkTimeout: p.AuthConfig().RefreshJwkTimeout,
		}, p.Validate(), p.AuthConfig().ClientId)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := p.webAuthProvider.CheckSsoConnection(ctx)
		if err != nil {
			slog.Error("failed to fetch jwks from sso", "error", err.Error())
			os.Exit(1)
		}
	}

	return p.webAuthProvider
}

func (p *Provider) AuthConfig() security.WebAuthConfig {
	if p.webAuthConfig == nil {
		cfg, err := security.NewAuthConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}
		p.webAuthConfig = cfg
	}
	return *p.webAuthConfig
}
