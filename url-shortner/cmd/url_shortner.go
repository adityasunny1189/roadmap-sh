package urlshortner

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/url-shortner/internal/config"
	"github.com/adityasunny1189/url-shortner/internal/core/services"
	"github.com/adityasunny1189/url-shortner/internal/repository"
	"github.com/adityasunny1189/url-shortner/internal/storage/cache"
	"github.com/adityasunny1189/url-shortner/internal/storage/database"
	"github.com/adityasunny1189/url-shortner/internal/transport/http/rest"
	"github.com/gorilla/mux"
)

func UrlShortnerServer() {
	r := mux.NewRouter()

	cfg := config.Load()
	db := database.Load(cfg)
	counterCache := cache.NewShortUrlAccessCountMemoryCache()
	urlShortenerRepository := repository.NewURLShortenerRepository(db)
	urlShortenerService := services.NewURLShortenerService(urlShortenerRepository, counterCache)
	handler := rest.NewHandler(urlShortenerService)

	r.HandleFunc("/shorten", handler.ShortenUrlHandler).Methods("POST")
	r.HandleFunc("/shorten/{shortCode}", handler.GetShortenUrl).Methods("GET")
	r.HandleFunc("/shorten/{shortCode}", handler.UpdateShortenUrl).Methods("PUT")
	r.HandleFunc("/shorten/{shortCode}", handler.DeleteShortenUrl).Methods("DELETE")
	r.HandleFunc("/shorten/{shortCode}/stats", handler.GetStatistics).Methods("GET")

	log.Println("Starting server on port: 8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Println(err)
	}
}
