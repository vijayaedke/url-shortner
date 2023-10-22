package server

import (
	"url-shortner/internal/app/handler"
	"url-shortner/internal/config"

	"github.com/gin-gonic/gin"
)

func NewAppRoutes(config config.Configuration, handler handler.BaseHandler) *gin.Engine {
	route := gin.Default()
	route.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Next()
	})
	route.UseRawPath = true
	route.UnescapePathValues = false

	routerApiGroup := route.Group("/api/v1")

	URLShortApiGroup := routerApiGroup.Group("/url-short")
	URLShortApiGroup.POST("/", handler.URLShortner)
	URLShortApiGroup.GET("/", handler.Redirect)

	return route
}
