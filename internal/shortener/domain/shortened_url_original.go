package domain

import (
	"errors"
	"net/url"
)

var (
	ErrInvalidURL = errors.New("invalid url")
)

type ShortenedURLOriginal struct {
	value string
}

func NewOriginal(value string) (*ShortenedURLOriginal, error) {
	if !isValidURL(value) {
		return nil, ErrInvalidURL
	}

	return &ShortenedURLOriginal{
		value,
	}, nil
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

func (v *ShortenedURLOriginal) Value() string {
	return v.value
}
