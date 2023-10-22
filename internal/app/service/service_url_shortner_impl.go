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

	getShortURL := s.utils.BaseConversion(request.URL)
	key := base64.StdEncoding.EncodeToString([]byte(request.URL))

	setShortURL, err := s.redisClient.SetURLStore(key, getShortURL)
	if err != nil {
		return nil, err
	}

	if setShortURL == "" {
		// if URL is not set already, do cross mapping for short URL with long URL key
		if _, err = s.redisClient.SetURLStore(getShortURL, key); err != nil {
			return nil, err
		}

		currentTime := time.Now()
		return &model.URLShortenResponse{
			ShortURL:  fmt.Sprintf("%s/%s", BITLY_URL, getShortURL),
			CreatedAt: &currentTime,
		}, nil
	}

	urlExists := true
	return &model.URLShortenResponse{
		ShortURL: fmt.Sprintf("%s/%s", BITLY_URL, setShortURL),
		Status:   &urlExists,
	}, nil

	// read file utility
	// response, err := s.utils.ConvertLongURLToShortURL(ctx, request.URL)
	// if err != nil {
	// 	return nil, err
	// }

	// return response, nil
}

func (s *URLService) Redirect(ctx context.Context, requestURL string) (*model.URLRequestResponse, error) {
	if requestURL == "" {
		return nil, fmt.Errorf("no query param found")
	}

	getURL, err := s.redisClient.GetURLStore(requestURL)
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
