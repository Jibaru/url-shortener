package application

import "url-shortener/internal/shortener/domain"

type GetShortenedURLS struct {
	repository domain.ShortenedURLRepository
}

func NewGetShortenedURLS(repository domain.ShortenedURLRepository) *GetShortenedURLS {
	return &GetShortenedURLS{repository: repository}
}

func (uc *GetShortenedURLS) Execute() []*domain.ShortenedURL {
	return uc.repository.GetAll()
}
