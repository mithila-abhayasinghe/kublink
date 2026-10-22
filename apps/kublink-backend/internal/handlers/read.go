package handlers

import (
	"fmt"
	"kublink-backend/internal/store"
	"net/http"
)

func GetLongUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("shortUrl")
	url, ok := store.URLStore[shortUrl]
	longUrl := url.OriginalUrl
	if !ok {
		fmt.Printf("/%s shortUrl not found\n", shortUrl)
		http.NotFound(w, r)
		return
	}
	fmt.Printf("GET /%s shortUrl redirection\n", shortUrl)
	http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)

}
