package server

import (
	"url-shortner/internal/app/handler"
	"url-shortner/internal/config"

	"github.com/gin-gonic/gin"
)

func NewAppRoutes(config config.Configuration, handler handler.BaseHandler) *gin.Engine {
	route := gin.Default()
	route.Use()
	routerApiGroup := route.Group("/api/v1")

	routerApiGroup.POST("/url-short", handler.URLShortner)
	routerApiGroup.GET("/{id}", handler.Redirect)

	return route
}

// func apiTimeout(c *gin.Context) gin.IRoutes {
// 	timeout := time.Second * 20
// 	ctx, cancel := context.WithTimeout(c, timeout)
// 	defer func() {
// 		cancel()
// 		if ctx.Err() == context.DeadlineExceeded {
// 			c.Request.Response.Header.Add("Gateway time out", http.StatusGatewayTimeout)
// 		}
// 	}()
// }
