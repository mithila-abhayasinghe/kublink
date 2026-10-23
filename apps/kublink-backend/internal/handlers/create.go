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

func (h *URLHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	var req requests.CreateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		ShortUrl: h.CurrentURL + "/" + id,
	}
	err = json.NewEncoder(w).Encode(shortUrlResponse)
	if err != nil {
		return
	}
}
