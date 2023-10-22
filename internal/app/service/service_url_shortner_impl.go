package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"url-shortner/internal/app/model"
)

const BITLY_URL = "https://bitly.com"

func (s *URLService) URLShortner(ctx context.Context, request *model.URLRequestResponse) (*model.URLShortenResponse, error) {
	if request == nil || (request != nil && request.URL == "") {
		return nil, fmt.Errorf("no request url found")
	}

	shortURL := s.utils.BaseConversion(request.URL)
	key := base64.StdEncoding.EncodeToString([]byte(request.URL))

	shortURL, err := s.redisClient.SetURLStore(key, shortURL)
	if err != nil {
		return nil, err
	}
	if _, err = s.redisClient.SetURLStore(shortURL, key); err != nil {
		return nil, err
	}

	if shortURL != "" {
		return &model.URLShortenResponse{
			ShortURL:  fmt.Sprintf("%s/%s", BITLY_URL, shortURL),
			CreatedAt: time.Now(),
		}, nil
	}

	return nil, nil

	// read file utility
	// response, err := s.utils.ConvertLongURLToShortURL(ctx, request.URL)
	// if err != nil {
	// 	return nil, err
	// }

	// return response, nil
}

func (s *URLService) Redirect(ctx context.Context, request *model.URLRequestResponse) (*model.URLRequestResponse, error) {
	getURL, err := s.redisClient.GetURLStore(request.URL)
	if err != nil {
		return nil, err
	}

	decodeURL, err := base64.StdEncoding.DecodeString(getURL)
	if err != nil {
		return nil, err
	}

	return &model.URLRequestResponse{
		URL: string(decodeURL),
	}, nil
}
