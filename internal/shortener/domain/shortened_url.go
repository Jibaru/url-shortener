package domain

type ShortenedURL struct {
	id   ShortenedURLID
	url  ShortenedURLOriginal
	code ShortenedURLCode
}

func NewShortenedURL(
	id ShortenedURLID,
	url ShortenedURLOriginal,
	code ShortenedURLCode,
) *ShortenedURL {
	return &ShortenedURL{
		id,
		url,
		code,
	}
}

func (m *ShortenedURL) HasCode(code ShortenedURLCode) bool {
	return m.code.Equals(code)
}

func (m *ShortenedURL) ID() ShortenedURLID {
	return m.id
}

func (m *ShortenedURL) URL() ShortenedURLOriginal {
	return m.url
}

func (m *ShortenedURL) Code() ShortenedURLCode {
	return m.code
}
