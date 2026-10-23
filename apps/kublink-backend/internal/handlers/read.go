package handlers

import (
	"net/http"
)

// GetLongUrl godoc
//
//	@Summary		Redirect to long URL
//	@Description	Looks up a short URL identifier and redirects the client to the original URL
//	@Tags			urls
//	@Param			shortUrl	path		string	true	"Short URL identifier"
//	@Success		307			{string}	string	"Temporary Redirect to original URL"
//	@Failure		404			{string}	string	"Not Found"
//	@Router			/{shortUrl} [get]
func (h *URLHandler) GetLongUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("shortUrl")
	url, ok := h.Store.Get(shortUrl)
	if !ok {
		h.Logger.Warn("short url not found", "short_url", shortUrl)
		http.NotFound(w, r)
		return
	}
	longUrl := url.OriginalUrl
	h.Logger.Info("shorturl redirection", "short_url", shortUrl)
	http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)

}
