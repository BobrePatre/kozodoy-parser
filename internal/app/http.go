package app

import (
	"context"
	"github.com/BobrePatre/kozodoy-parser/internal/constants"
	"github.com/BobrePatre/kozodoy-parser/internal/delivery/http/middlewares"
	"github.com/BobrePatre/kozodoy-parser/internal/delivery/http/routes/parser"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

var _ = (*App)(nil)

func (a *App) initHTTPServer(_ context.Context) error {

	switch a.diProvider.AppConfig().MODE {
	case constants.EnvDevelopment:
		gin.SetMode(gin.DebugMode)
	case constants.EnvProduction:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.SlogLoggerMiddleware())

	corsCfg := a.diProvider.CorsConfig()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:           corsCfg.AllowAllOrigins,
		AllowOrigins:              corsCfg.AllowOrigins,
		AllowMethods:              corsCfg.AllowMethods,
		AllowHeaders:              corsCfg.AllowHeaders,
		MaxAge:                    corsCfg.MaxAge,
		AllowCredentials:          corsCfg.AllowCredentials,
		AllowWildcard:             corsCfg.AllowWildcard,
		ExposeHeaders:             corsCfg.ExposeHeaders,
		AllowBrowserExtensions:    corsCfg.AllowBrowserExtensions,
		AllowWebSockets:           corsCfg.AllowWebSockets,
		AllowFiles:                corsCfg.AllowFiles,
		AllowPrivateNetwork:       corsCfg.AllowPrivateNetwork,
		OptionsResponseStatusCode: corsCfg.OptionsResponseStatusCode,
		CustomSchemas:             corsCfg.CustomSchemas,
	}))

	authMiddlewareConstructor := a.diProvider.HttpAuthMiddlewareConstructor()
	v1RouterGroup := router.Group("/api/v1")

	parser.NewRouter(
		v1RouterGroup,
		authMiddlewareConstructor(a.diProvider.WebAuthProvider()),
		a.diProvider.ParserHandler(),
	).Register()

	a.httpServer = &http.Server{
		Addr:    a.diProvider.HTTPConfig().Address(),
		Handler: router,
	}
	return nil
}

func (a *App) runHTTPServer() error {
	slog.Info(
		startingMsg,
		httpServerTag,
		slog.String(addressMsg, a.diProvider.HTTPConfig().Address()),
	)
	return a.httpServer.ListenAndServe()
}
