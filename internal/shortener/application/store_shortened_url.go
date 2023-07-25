package application

import (
	"url-shortener/internal/shortener/domain"
)

type StoreShortenedURL struct {
	idGenerator domain.IDGenerator
	repository  domain.ShortenedURLRepository
}

func NewStoreShortenedURL(
	idGenerator domain.IDGenerator,
	repository domain.ShortenedURLRepository,
) *StoreShortenedURL {
	return &StoreShortenedURL{
		idGenerator: idGenerator,
		repository:  repository,
	}
}

func (uc *StoreShortenedURL) Execute(
	url string,
) (*domain.ShortenedURL, error) {
	originalURL, err := domain.NewOriginal(url)
	if err != nil {
		return nil, err
	}

	code, err := domain.NewRandomShortenedURLCode()
	if err != nil {
		return nil, err
	}

	for uc.repository.ExistsByCode(code) {
		code, err = domain.NewRandomShortenedURLCode()
		if err != nil {
			return nil, err
		}
	}

	shortenedURL := domain.NewShortenedURL(
		*domain.NewShortenedURLID(uc.idGenerator.Generate()),
		*originalURL,
		*code,
	)

	uc.repository.Save(shortenedURL)

	return shortenedURL, nil
}
