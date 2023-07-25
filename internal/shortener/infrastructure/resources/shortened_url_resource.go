package resources

import "url-shortener/internal/shortener/domain"

type ShortenedURLResource struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	URL  string `json:"url"`
}

func NewShortenedURLResource(shortenedURL *domain.ShortenedURL) ShortenedURLResource {
	id := shortenedURL.ID()
	code := shortenedURL.Code()
	url := shortenedURL.URL()

	return ShortenedURLResource{
		ID:   id.Value(),
		Code: code.Value(),
		URL:  url.Value(),
	}
}
