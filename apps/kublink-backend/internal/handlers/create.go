package handlers

import (
	"encoding/json"
	"kublink-backend/internal/requests"
	"kublink-backend/internal/types"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// CreateShortUrl godoc
//
//	@Summary		Create a short URL
//	@Description	Takes a long URL and generates a unique short code shortcut
//	@Tags			urls
//	@Accept			json
//	@Produce		json
//	@Param			request	body		requests.CreateShortUrlRequest	true	"Original URL Payload"
//	@Success		201		{object}	requests.ShortUrlCreatedResponse
//	@Failure		400		{object}	requests.ErrorResponse	"Bad Request"
//	@Router			/api/v1/urls [post]
func (h *URLHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(requests.ErrorResponse{Error: "Content-Type header must be application/json"})
		h.Logger.Warn("invalid content type", "content_type", r.Header.Get("Content-Type"))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1000000)

	var req requests.CreateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requests.ErrorResponse{Error: err.Error()})
		h.Logger.Warn("short url creation failed", "err", err.Error())
		return
	}

	parsedURL, err := url.ParseRequestURI(req.OriginalUrl)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requests.ErrorResponse{Error: err.Error()})
		h.Logger.Warn("original url validation failed", "err", err.Error())
		return
	}

	newId := h.counter.Add(1)
	id := strconv.FormatInt(newId, 10)
	newUrl := types.URLs{
		Id:          id,
		OriginalUrl: parsedURL.String(),
		Alias:       "",
		CreatedAt:   time.Now(),
		ExpireAt:    time.Now().AddDate(0, 0, 2),
	}

	err = h.Store.Create(r.Context(), newUrl)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requests.ErrorResponse{Error: err.Error()})
		h.Logger.Warn("shorturl creation failed", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	shortUrlResponse := requests.ShortUrlCreatedResponse{
		ShortUrl: h.BaseURL + "/" + id,
	}
	json.NewEncoder(w).Encode(shortUrlResponse)
	h.Logger.Info("shorturl creation successful", "short_url", id)
}
