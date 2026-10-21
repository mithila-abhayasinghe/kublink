package handlers

import (
	"fmt"
	"net/http"
)

var urls = map[string]string{
	"gg":        "https://www.google.com",
	"portfolio": "https://www.mithila.cc",
}

func GetLongUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("shortUrl")
	longUrl, ok := urls[shortUrl]
	if !ok {
		fmt.Printf("/%s shortUrl not found\n", shortUrl)
		http.NotFound(w, r)
		return
	}
	fmt.Printf("GET /%s shortUrl redirection\n", shortUrl)
	http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)

}
