package services

import (
	"errors"
	"log"

	"github.com/adityasunny1189/url-shortner/internal/core/domain"
	"github.com/adityasunny1189/url-shortner/internal/core/ports"
	"github.com/adityasunny1189/url-shortner/internal/models"
)

type shortenService struct {
	urlShortenerCache ports.UrlShortnerCache
	urlShortenerRepo  ports.UrlShortnerRepository
}

func NewURLShortenerService(urlShortenerRepo ports.UrlShortnerRepository,
	urlShortenerCache ports.UrlShortnerCache) ports.UrlShortnerService {
	return &shortenService{
		urlShortenerRepo:  urlShortenerRepo,
		urlShortenerCache: urlShortenerCache,
	}
}

func (ss *shortenService) ShortenURL(shortenUrlReq models.CreateShortUrlRequest) (domain.ShortURL, error) {
	// check if the url is already present or not
	shortUrl, err := ss.urlShortenerRepo.FindByOriginalUrl(shortenUrlReq.Url)
	if err != nil {
		log.Println("Service: url not found, creating new data: ", shortenUrlReq)

		// call the repository to save the data and return back the new short url data
		shortUrl, err = ss.urlShortenerRepo.Save(shortenUrlReq.Url)
		if err != nil {
			return domain.ShortURL{}, err
		}

		// return the response back to serrvice layer
		return shortUrl, nil
	}

	log.Println("Service: url already present: ", shortUrl)

	return shortUrl, nil
}

func (ss *shortenService) GetURL(shortCode string) (domain.ShortURL, error) {
	// check if url is present
	urlData, err := ss.urlShortenerRepo.FindByShortCode(shortCode)
	if err != nil {
		return domain.ShortURL{}, err
	}

	// check if url shortcode mathes the requested shortCode
	if urlData.ShortCode != shortCode {
		return urlData, errors.New("url not found")
	}

	// increase the analytical counter from cache
	ss.urlShortenerCache.Increment(shortCode)

	// if found return response else return error
	return urlData, nil
}

func (ss *shortenService) UpdateURL(shortCode string,
	updateUrlReq models.UpdateShortUrlRequest) (domain.ShortURL, error) {
	// check if url is present
	urlData, err := ss.urlShortenerRepo.FindByShortCode(shortCode)
	if err != nil {
		return domain.ShortURL{}, err
	}

	// check if url shortcode mathes the requested shortCode
	if urlData.ShortCode != shortCode {
		return urlData, errors.New("url not found")
	}

	// call the repository to update the url
	urlData, err = ss.urlShortenerRepo.Update(shortCode, updateUrlReq.NewUrl)
	if err != nil {
		return urlData, err
	}

	return urlData, nil
}

func (ss *shortenService) DeleteURL(shortCode string) error {
	urlData, err := ss.urlShortenerRepo.FindByShortCode(shortCode)
	if err != nil {
		return err
	}

	// check if url shortcode mathes the requested shortCode
	if urlData.ShortCode != shortCode {
		return errors.New("url not found")
	}

	// call the repository to delete the url
	err = ss.urlShortenerRepo.Delete(shortCode)
	if err != nil {
		return err
	}

	// remove the shortCode from analytical cache
	ss.urlShortenerCache.Delete(shortCode)

	return nil
}

func (ss *shortenService) GetURLStats(shortCode string) (models.GetUrlStatisticsResponse, error) {
	urlData, err := ss.urlShortenerRepo.FindByShortCode(shortCode)
	if err != nil {
		return models.GetUrlStatisticsResponse{}, err
	}

	// check if url shortcode mathes the requested shortCode
	if urlData.ShortCode != shortCode {
		return models.GetUrlStatisticsResponse{}, errors.New("url not found")
	}

	// read the analytical data from in memory cache
	urlStatResp := models.GetUrlStatisticsResponse{
		ShortUrl:    urlData,
		AccessCount: ss.urlShortenerCache.Get(shortCode),
	}

	return urlStatResp, nil
}
