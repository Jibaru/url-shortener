package domain

type ShortenedURLRepository interface {
	Save(url *ShortenedURL)
	ExistsByCode(code *ShortenedURLCode) bool
	GetAll() []*ShortenedURL
	GetByCode(code *ShortenedURLCode) *ShortenedURL
}
