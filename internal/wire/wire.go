//go:build wireinject
// +build wireinject

package wire

import (
	"url-shortner/internal/app/handler"
	"url-shortner/internal/app/server"
	"url-shortner/internal/app/service"
	"url-shortner/internal/app/utils"
	"url-shortner/internal/config"
	"url-shortner/pkg/redis"

	"github.com/google/wire"
)

func InitializeServer() (server.Server, func(), error) {
	panic(wire.Build(
		server.NewURLShortenerServer,
		server.NewAppRoutes,
		service.NewURLShortnerService,
		redis.NewRedisClient,
		handler.NewBaseHandler,
		config.NewConfig,
		utils.NewUtilityService,
	))
}
