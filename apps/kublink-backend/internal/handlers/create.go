package handlers

import (
	"encoding/json"
	"fmt"
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

	var req requests.CreateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requests.ErrorResponse{Error: err.Error()})
		return
	}

	newId := h.counter.Add(1)
	id := strconv.FormatInt(newId, 10)

	_, err := url.Parse(req.OriginalUrl)
	if err != nil {
		fmt.Println("error: not a valid url")
		return
	}

	newUrl := types.URLs{
		Id:          id,
		OriginalUrl: req.OriginalUrl,
		Alias:       "none",
		CreatedAt:   time.Now(),
		ExpireAt:    time.Now().AddDate(0, 0, 2),
	}

	h.Store.Set(id, newUrl)

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("POST /%s shortUrl created\n", id)

	shortUrlResponse := requests.ShortUrlCreatedResponse{
		ShortUrl: h.BaseURL + "/" + id,
	}
	err = json.NewEncoder(w).Encode(shortUrlResponse)
	if err != nil {
		return
	}
}
