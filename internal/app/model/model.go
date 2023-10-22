package model

import "time"

type URLRequestResponse struct {
	URL string `json:"url"`
}

type URLShortenResponse struct {
	ShortURL  string     `json:"short_url"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Status    *bool      `json:"url_exists,omitempty"`
}

type URLStatsResponse struct {
	URLlist []*URLlist `json:"url_list,omitempty"`
}

type URLlist struct {
	URL   string `json:"url"`
	Count int    `json:"reuqested_count"`
}
