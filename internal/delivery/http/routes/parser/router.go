package parser

import (
	webAuthProvider "github.com/BobrePatre/kozodoy-parser/internal/providers/web_auth"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Parse(ctx *gin.Context)
}

type Router struct {
	router         *gin.RouterGroup
	authMiddleware webAuthProvider.HttpMiddleware
	handler        Handler
}

func NewRouter(router *gin.RouterGroup, authMiddleware webAuthProvider.HttpMiddleware, handler Handler) *Router {
	return &Router{
		router:         router,
		authMiddleware: authMiddleware,
		handler:        handler,
	}
}

func (r *Router) Register() {
	routerGroup := r.router.Group("/parser")
	{
		routerGroup.POST("/parse", r.authMiddleware("admin", "manager"), r.handler.Parse)
	}
}
