package domain

import (
	"time"
)

type ShortURL struct {
	Id          int64
	OriginalUrl string
	ShortCode   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
