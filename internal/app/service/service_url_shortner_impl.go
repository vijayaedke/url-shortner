package service

import (
	"context"
	"fmt"

	"url-shortner/internal/app/model"
)

const keyPrefix = "urlshorten:"

func (s *URLService) URLShortner(ctx context.Context, request *model.URLShortenRequest) (*model.URLShortenResponse, error) {
	if request == nil || (request != nil && request.URL == "") {
		return nil, fmt.Errorf("no request url found")
	}

	response, err := s.utils.ConvertLongURLToShortURL(request.URL)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *URLService) Redirect(ctx context.Context) error {
	return nil
}
