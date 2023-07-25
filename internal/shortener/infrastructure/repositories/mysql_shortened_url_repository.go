package repositories

import (
	"database/sql"
	"log"
	"url-shortener/internal/shortener/domain"
)

type MySQLShortenedURLRepository struct {
	db *sql.DB
}

func NewMySQLShortenedURLRepository(db *sql.DB) domain.ShortenedURLRepository {
	return &MySQLShortenedURLRepository{
		db: db,
	}
}

func (r *MySQLShortenedURLRepository) Save(url *domain.ShortenedURL) {
	id := url.ID()
	original := url.URL()
	code := url.Code()

	_, err := r.db.Exec("INSERT INTO shortened_urls (id, url, code) VALUES (?, ?, ?)",
		id.Value(), original.Value(), code.Value())
	if err != nil {
		log.Println(err)
	}
}

func (r *MySQLShortenedURLRepository) ExistsByCode(code *domain.ShortenedURLCode) bool {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM shortened_urls WHERE code = ?", code.Value()).Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (r *MySQLShortenedURLRepository) GetAll() []*domain.ShortenedURL {
	var urls []*domain.ShortenedURL
	rows, err := r.db.Query("SELECT id, url, code FROM shortened_urls")
	if err != nil {
		log.Println(err)
		return urls
	}
	defer rows.Close()

	for rows.Next() {
		var id, url, code string
		err := rows.Scan(&id, &url, &code)
		if err != nil {
			log.Println(err)
			return urls
		}

		original, err := domain.NewOriginal(url)
		if err != nil {
			log.Println(err)
			return urls
		}

		urlCode, err := domain.NewShortenedURLCode(code)
		if err != nil {
			log.Println(err)
			return urls
		}

		shortenedURL := domain.NewShortenedURL(
			*domain.NewShortenedURLID(id),
			*original,
			*urlCode,
		)
		urls = append(urls, shortenedURL)
	}

	return urls
}

func (r *MySQLShortenedURLRepository) GetByCode(code *domain.ShortenedURLCode) *domain.ShortenedURL {
	var id, url string
	err := r.db.QueryRow("SELECT id, url FROM shortened_urls WHERE code = ?", code.Value()).Scan(&id, &url)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return nil
		}
		log.Println(err)
		return nil
	}

	original, err := domain.NewOriginal(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	return domain.NewShortenedURL(
		*domain.NewShortenedURLID(id),
		*original,
		*code,
	)
}
