package server

import (
	"url-shortner/internal/app/handler"
	"url-shortner/internal/config"

	"github.com/gin-gonic/gin"
)

func NewAppRoutes(config config.Configuration, handler handler.BaseHandler) *gin.Engine {
	route := gin.Default()

	routerApiGroup := route.Group("/api/v1")

	routerApiGroup.POST("/url-short", handler.URLShortner)
	routerApiGroup.GET("/{id}", handler.Redirect)

	return route
}
