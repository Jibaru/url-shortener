package resources

import (
	"encoding/json"
	"url-shortener/internal/shortener/domain"
)

type ShortenedURLCollection struct {
	Data []ShortenedURLResource `json:"data"`
}

func NewShortenedURLCollection(shortenedURLS []*domain.ShortenedURL) *ShortenedURLCollection {
	var resources []ShortenedURLResource
	for _, shortenedURL := range shortenedURLS {
		resources = append(resources, NewShortenedURLResource(shortenedURL))
	}

	return &ShortenedURLCollection{resources}
}

func (r *ShortenedURLCollection) MarshalJSON() ([]byte, error) {
	if len(r.Data) == 0 {
		return []byte("[]"), nil
	}

	return json.Marshal(r.Data)
}
