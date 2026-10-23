package handlers

import (
	"encoding/json"
	"fmt"
	"kublink-backend/internal/requests"
	"kublink-backend/internal/store"
	"kublink-backend/internal/types"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var counter int = 1

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	var req requests.CreateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := strconv.Itoa(counter)

	_, err := url.Parse(req.OriginalUrl)
	if err != nil {
		fmt.Println("error: not a valid url")
	}

	url := types.URLs{
		Id:          id,
		OriginalUrl: req.OriginalUrl,
		Alias:       "none",
		CreatedAt:   time.Now(),
		ExpireAt:    time.Now().AddDate(0, 0, 2),
	}

	store.URLStore[id] = url
	counter++

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("POST /%s shortUrl created\n", id)

}
