package requests

type CreateShortUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}

type ShortUrlCreatedResponse struct {
	ShortUrl string `json:"short_url"`
}
