package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"

	"url-shortner/internal/app/model"
)

func (s *URLService) GetURLStats(ctx context.Context) (*model.URLStatsResponse, error) {
	keySet, err := s.redisClient.GetAllURLStore()
	if err != nil {
		return nil, err
	}

	urlListResponse := make([]*model.URLlist, 0, len(keySet))
	hashSet := make(map[string]int, len(keySet))
	for _, key := range keySet {
		if len(key) > 7 {
			decodeKey, _ := base64.StdEncoding.DecodeString(key)
			parsedURL, err := url.Parse(string(decodeKey))
			if err != nil {
				fmt.Println("failed to parse URL", err)
			} else {
				hashSet[parsedURL.Host]++
			}
		}
	}

	for key, val := range hashSet {
		urlListResponse = append(urlListResponse, &model.URLlist{
			URL:   key,
			Count: val,
		})
	}

	sort.SliceStable(urlListResponse, func(i, j int) bool {
		return urlListResponse[i].Count > urlListResponse[j].Count
	})

	if len(urlListResponse) >= 3 {
		return &model.URLStatsResponse{
			URLlist: urlListResponse[:3],
		}, nil
	}

	return &model.URLStatsResponse{
		URLlist: urlListResponse,
	}, nil
}
