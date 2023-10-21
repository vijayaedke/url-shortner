package server

import (
	"fmt"
	"strconv"

	"url-shortner/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Configuration
	router *gin.Engine
}

func NewURLShortenerServer(config config.Configuration, router *gin.Engine) Server {
	return Server{
		router: router,
		config: config,
	}
}

func (s *Server) StartServer() {
	r := s.router
	r.Run(s.getServerHostConfig())
}

func (s *Server) getServerHostConfig() string {
	return fmt.Sprintf("%s:%s", s.config.GetServerHost(), strconv.Itoa(s.config.GetServerPort()))
}
