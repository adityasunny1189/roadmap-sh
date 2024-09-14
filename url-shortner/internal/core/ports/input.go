package ports

import "github.com/adityasunny1189/url-shortner/internal/core/domain"

type UrlShortnerRepository interface {
	Save(url string) (domain.ShortURL, error)
	Update(shortCode, updatedUrl string) (domain.ShortURL, error)
	Delete(shortCode string) error
	FindByShortCode(shortCode string) (domain.ShortURL, error)
	FindByOriginalUrl(url string) (domain.ShortURL, error)
}

type UrlShortnerCache interface {
	Increment(shortCode string)
	Get(shortCode string) int
	Delete(shortCode string)
}
