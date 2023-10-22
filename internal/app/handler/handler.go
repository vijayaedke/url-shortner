package handler

import (
	"url-shortner/internal/app/service"
	"url-shortner/internal/config"

	"github.com/gin-gonic/gin"
)

type BaseHandler interface {
	URLShortner(c *gin.Context)
	Redirect(c *gin.Context)
	GetURLStats(c *gin.Context)
}

type Handler struct {
	config  config.Configuration
	service service.Service
}

func NewBaseHandler(config config.Configuration, service service.Service) BaseHandler {
	return &Handler{
		config:  config,
		service: service,
	}
}
