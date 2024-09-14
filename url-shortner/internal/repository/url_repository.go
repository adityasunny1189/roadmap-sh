package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/adityasunny1189/url-shortner/internal/core/domain"
	"github.com/adityasunny1189/url-shortner/internal/core/ports"
)

type urlShortenRepo struct {
	db *sql.DB
}

func (ur *urlShortenRepo) Delete(shortCode string) error {
	res, err := ur.db.Exec("DELETE FROM urls WHERE short_code = ?", shortCode)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DELETE_URL: no url found with shortCode %s", shortCode)
	}

	return nil
}

func (ur *urlShortenRepo) FindByOriginalUrl(url string) (domain.ShortURL, error) {
	var shortUrl domain.ShortURL
	var createdAt, updatedAt []byte

	log.Println("Repository: checking url if it is present or not: ", url)

	row := ur.db.QueryRow("SELECT * FROM urls WHERE original_url = ?", url)

	log.Println("Repository: row data: ", row)

	if err := row.Scan(&shortUrl.Id,
		&shortUrl.OriginalUrl, &shortUrl.ShortCode,
		&createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return shortUrl, errors.New("url not found for the given url")
		}
		return shortUrl, errors.New("error while retrieving url data")
	}

	createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return shortUrl, fmt.Errorf("error parsing created_at: %v", err)
	}

	shortUrl.CreatedAt = createdAtTime

	updatedAtTime, err := time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return shortUrl, fmt.Errorf("error parsing updated_at: %v", err)
	}

	shortUrl.UpdatedAt = updatedAtTime

	return shortUrl, nil
}

func (ur *urlShortenRepo) FindByShortCode(shortCode string) (domain.ShortURL, error) {
	var shortUrl domain.ShortURL
	var createdAt, updatedAt []byte

	row := ur.db.QueryRow("SELECT * FROM urls WHERE short_code = ?", shortCode)
	if err := row.Scan(&shortUrl.Id,
		&shortUrl.OriginalUrl, &shortUrl.ShortCode,
		&createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return shortUrl, errors.New("url not found for the given shortcode")
		}
		return shortUrl, errors.New("error while retrieving url data")
	}

	createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return shortUrl, fmt.Errorf("error parsing created_at: %v", err)
	}

	shortUrl.CreatedAt = createdAtTime

	updatedAtTime, err := time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return shortUrl, fmt.Errorf("error parsing updated_at: %v", err)
	}

	shortUrl.UpdatedAt = updatedAtTime

	return shortUrl, nil
}

func (ur *urlShortenRepo) Save(url string) (domain.ShortURL, error) {
	splitUrlData := strings.Split(url, ".")
	shortCode := splitUrlData[1][0:3] + "-code"

	log.Println("Repository: creating new url with shortcode: ", shortCode)

	res, err := ur.db.Exec("INSERT INTO urls (original_url, short_code) VALUES (?, ?)", url, shortCode)
	if err != nil {
		return domain.ShortURL{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return domain.ShortURL{}, err
	}

	log.Println("Repository: id of created url data: ", id)

	// Find the inserted data using id returned
	var shortUrl domain.ShortURL
	row := ur.db.QueryRow("SELECT * FROM urls WHERE id = ?", id)

	var createdAt, updatedAt []byte

	if err := row.Scan(&shortUrl.Id,
		&shortUrl.OriginalUrl, &shortUrl.ShortCode,
		&createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return shortUrl, errors.New("url not found for the given id")
		}
		return shortUrl, err
	}

	createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return shortUrl, fmt.Errorf("error parsing created_at: %v", err)
	}

	shortUrl.CreatedAt = createdAtTime

	updatedAtTime, err := time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return shortUrl, fmt.Errorf("error parsing updated_at: %v", err)
	}

	shortUrl.UpdatedAt = updatedAtTime

	return shortUrl, nil
}

func (ur *urlShortenRepo) Update(shortCode, updatedUrl string) (domain.ShortURL, error) {
	res, err := ur.db.Exec("UPDATE urls SET original_url = ? WHERE short_code = ?",
		updatedUrl, shortCode)
	if err != nil {
		return domain.ShortURL{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.ShortURL{}, err
	}

	if rowsAffected == 0 {
		return domain.ShortURL{}, fmt.Errorf("UPDATE_URL: no url data found with shortCode: %s",
			shortCode)
	}

	return ur.FindByShortCode(shortCode)
}

func NewURLShortenerRepository(db *sql.DB) ports.UrlShortnerRepository {
	return &urlShortenRepo{
		db: db,
	}
}
