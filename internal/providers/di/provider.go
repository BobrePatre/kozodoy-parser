package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/core"
	"github.com/BobrePatre/kozodoy-parser/internal/config/datasources"
	"github.com/BobrePatre/kozodoy-parser/internal/config/delivery"
	"github.com/BobrePatre/kozodoy-parser/internal/config/security"
	parserHandler "github.com/BobrePatre/kozodoy-parser/internal/delivery/http/handlers/parser"
	webAuthProvider "github.com/BobrePatre/kozodoy-parser/internal/providers/web_auth"
	parserRepository "github.com/BobrePatre/kozodoy-parser/internal/repository/parser"
	parserService "github.com/BobrePatre/kozodoy-parser/internal/service/parser"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Provider struct {
	redisClient *redis.Client
	redisConfig *datasources.RedisConfig

	sqlDatabase      *sqlx.DB
	postgresqlConfig *datasources.PostgresqlConfig

	validate *validator.Validate

	corsConfig *security.CorsConfig
	httpConfig *delivery.HttpConfig
	appConfig  *core.AppConfig

	webAuthProvider webAuthProvider.Provider
	webAuthConfig   *security.WebAuthConfig

	httpAuthMiddlewareConstructor       webAuthProvider.HttpMiddlewareConstructor
	grpcUnaryAuthInterceptorConstructor webAuthProvider.GrpcUnaryInterceptorConstructor

	parserHandler    *parserHandler.Handler
	parserService    *parserService.Service
	parserRepository *parserRepository.Repository
}

func NewDiProvider() *Provider {
	return &Provider{}
}
