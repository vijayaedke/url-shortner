package model

import "time"

type URLShortenRequest struct {
	URL string
}

type URLShortenResponse struct {
	ShortURL  string
	CreatedAt time.Time
}
