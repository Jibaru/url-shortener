package application

import (
	"errors"
	"url-shortener/internal/shortener/domain"
)

var (
	ErrShortenedURLNotFound = errors.New("shortened url not found")
)

type GetShortenedURLByCode struct {
	repository domain.ShortenedURLRepository
}

func NewGetShortenedURLByCode(repository domain.ShortenedURLRepository) *GetShortenedURLByCode {
	return &GetShortenedURLByCode{repository}
}

func (uc *GetShortenedURLByCode) Execute(code string) (*domain.ShortenedURL, error) {
	shortenedURLCode, err := domain.NewShortenedURLCode(code)
	if err != nil {
		return nil, err
	}

	shortenedURL := uc.repository.GetByCode(shortenedURLCode)
	if shortenedURL == nil {
		return nil, ErrShortenedURLNotFound
	}

	return shortenedURL, nil
}
