package requests

type CreateShortUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}
