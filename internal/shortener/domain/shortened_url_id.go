package domain

type ShortenedURLID struct {
	value string
}

func NewShortenedURLID(value string) *ShortenedURLID {
	return &ShortenedURLID{
		value,
	}
}

func (v *ShortenedURLID) Value() string {
	return v.value
}
