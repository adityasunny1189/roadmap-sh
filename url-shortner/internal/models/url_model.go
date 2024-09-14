package models

import "github.com/adityasunny1189/url-shortner/internal/core/domain"

type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

type UpdateShortUrlRequest struct {
	NewUrl string `json:"url"`
}

type GetUrlStatisticsResponse struct {
	ShortUrl    domain.ShortURL `json:"data"`
	AccessCount int             `json:"accessCount"`
}
