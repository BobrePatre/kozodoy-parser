package app

import (
	"context"
	"github.com/BobrePatre/kozodoy-parser/internal/constants"
	"github.com/BobrePatre/kozodoy-parser/internal/delivery/http/middlewares"
	"github.com/BobrePatre/kozodoy-parser/internal/delivery/http/routes/parser"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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

	_ = a.diProvider.CorsConfig()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:       []string{"*"},
		ExposedHeaders:       []string{"*"},
		AllowCredentials:     true,
		AllowPrivateNetwork:  true,
		OptionsPassthrough:   true,
		OptionsSuccessStatus: 200,
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
