package service

import (
	"context"

	"url-shortner/internal/app/model"
	"url-shortner/internal/app/utils"
	"url-shortner/internal/config"

	"url-shortner/pkg/redis"
)

type Service interface {
	URLShortner(ctx context.Context, request *model.URLRequestResponse) (*model.URLShortenResponse, error)
	Redirect(ctx context.Context, request string) (*model.URLRequestResponse, error)
	GetURLStats(ctx context.Context) (*model.URLStatsResponse, error)
}

type URLService struct {
	config      config.Configuration
	utils       utils.Utility
	redisClient redis.RedisService
}

func NewURLShortnerService(config config.Configuration, utils utils.Utility, redisClient redis.RedisService) Service {
	return &URLService{
		config:      config,
		utils:       utils,
		redisClient: redisClient,
	}
}
