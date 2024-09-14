package cache

import (
	"sync"

	"github.com/adityasunny1189/url-shortner/internal/core/ports"
)

type ShortUrlAccessCountMemoryCache struct {
	data  map[string]int
	mutex *sync.RWMutex
}

func NewShortUrlAccessCountMemoryCache() ports.UrlShortnerCache {
	return &ShortUrlAccessCountMemoryCache{
		data:  make(map[string]int),
		mutex: &sync.RWMutex{},
	}
}

func (s *ShortUrlAccessCountMemoryCache) Increment(shortCode string) {
	s.mutex.Lock()
	s.data[shortCode]++
	s.mutex.Unlock()
}

func (s *ShortUrlAccessCountMemoryCache) Get(shortCode string) int {
	s.mutex.RLock()
	val := s.data[shortCode]
	s.mutex.RUnlock()
	return val
}

func (s *ShortUrlAccessCountMemoryCache) Delete(shortCode string) {
	s.mutex.Lock()
	delete(s.data, shortCode)
	s.mutex.Unlock()
}
