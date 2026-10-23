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
	"sync/atomic"
	"time"
)

var counter atomic.Int64

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	counter.Store(1)

	var req requests.CreateShortUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := strconv.FormatInt(counter.Load(), 10)

	_, err := url.Parse(req.OriginalUrl)
	if err != nil {
		fmt.Println("error: not a valid url")
		return
	}

	url := types.URLs{
		Id:          id,
		OriginalUrl: req.OriginalUrl,
		Alias:       "none",
		CreatedAt:   time.Now(),
		ExpireAt:    time.Now().AddDate(0, 0, 2),
	}

	store.URLStore[id] = url
	counter.Add(1)

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("POST /%s shortUrl created\n", id)

}
