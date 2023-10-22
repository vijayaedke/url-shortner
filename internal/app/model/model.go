package model

import "time"

type URLRequestResponse struct {
	URL string `json:"url"`
}

type URLShortenResponse struct {
	ShortURL  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
