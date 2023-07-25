package repositories

import (
	"url-shortener/internal/shortener/domain"
)

type MemoryShortenedURLRepository struct {
	values []*domain.ShortenedURL
}

func NewMemoryShortenedURLRepository() domain.ShortenedURLRepository {
	return &MemoryShortenedURLRepository{}
}

func (r *MemoryShortenedURLRepository) Save(url *domain.ShortenedURL) {
	r.values = append(r.values, url)
}

func (r *MemoryShortenedURLRepository) ExistsByCode(code *domain.ShortenedURLCode) bool {
	for _, value := range r.values {
		if value.HasCode(*code) {
			return true
		}
	}

	return false
}

func (r *MemoryShortenedURLRepository) GetAll() []*domain.ShortenedURL {
	return r.values
}

func (r *MemoryShortenedURLRepository) GetByCode(code *domain.ShortenedURLCode) *domain.ShortenedURL {
	for _, value := range r.values {
		if value.HasCode(*code) {
			return value
		}
	}

	return nil
}
