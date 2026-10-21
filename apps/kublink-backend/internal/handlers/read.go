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
	fmt.Printf("GET /%s shorturl redirection\n", shortUrl)

	longUrl, ok := urls[shortUrl]
	if !ok {
		fmt.Println("url not found")
	}

	w.Header().Set("Location", longUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)

}
