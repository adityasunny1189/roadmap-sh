package ports

import (
	"github.com/adityasunny1189/url-shortner/internal/core/domain"
	"github.com/adityasunny1189/url-shortner/internal/models"
)

type UrlShortnerService interface {
	ShortenURL(shortenUrlReq models.CreateShortUrlRequest) (domain.ShortURL, error)
	GetURL(shortCode string) (domain.ShortURL, error)
	UpdateURL(shortCode string, updateUrlReq models.UpdateShortUrlRequest) (domain.ShortURL, error)
	DeleteURL(shortCode string) (error)
	GetURLStats(shortCode string) (models.GetUrlStatisticsResponse, error)
}
