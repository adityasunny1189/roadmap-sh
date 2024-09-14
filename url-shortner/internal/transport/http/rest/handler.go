package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adityasunny1189/url-shortner/internal/core/ports"
	"github.com/adityasunny1189/url-shortner/internal/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	urlShortenerService ports.UrlShortnerService
}

func NewHandler(urlShortenerService ports.UrlShortnerService) *Handler {
	return &Handler{
		urlShortenerService: urlShortenerService,
	}
}

func (h *Handler) ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request
	var shortenUrlReq models.CreateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&shortenUrlReq); err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	log.Println("Handler: Shorten url request: ", shortenUrlReq)

	// call service
	newUrl, err := h.urlShortenerService.ShortenURL(shortenUrlReq)
	if err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	log.Println("Handler: new url generated: ", newUrl)

	// return response
	sendJsonResponse(w, 201, newUrl, nil)
}

func (h *Handler) GetShortenUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// validate short code

	// call service to get the url
	shortUrl, err := h.urlShortenerService.GetURL(shortCode)
	if err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	sendJsonResponse(w, 200, shortUrl, nil)
}

func (h *Handler) UpdateShortenUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// validate short code

	// parse the request body
	var updateUrlReq models.UpdateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&updateUrlReq); err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	// call service to get the url
	updatedUrl, err := h.urlShortenerService.UpdateURL(shortCode, updateUrlReq)
	if err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	// send response
	sendJsonResponse(w, 200, updatedUrl, nil)
}

func (h *Handler) DeleteShortenUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// validate short code

	// call service to delete the url
	err := h.urlShortenerService.DeleteURL(shortCode)
	if err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	// send response
	sendJsonResponse(w, 204, nil, nil)
}

func (h *Handler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// validate short code

	// call service to get the statistics of url
	statResp, err := h.urlShortenerService.GetURLStats(shortCode)
	if err != nil {
		// handle error
		sendErrorResponse(w, models.BOTTOM_OVERLAY, err.Error())
		return
	}

	// send response
	sendJsonResponse(w, 200, statResp, nil)
}
