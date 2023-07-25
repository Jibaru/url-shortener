package domain

import (
	"errors"
	"math/rand"
)

const (
	ExactCodeLengthAllowed = 5
)

var (
	ErrCanNotGenerateRandomCode = errors.New("can not generate random value")
	ErrInvalidCodeLength        = errors.New("invalid code length")
)

type ShortenedURLCode struct {
	value string
}

func NewShortenedURLCode(value string) (*ShortenedURLCode, error) {
	if len(value) != ExactCodeLengthAllowed {
		return nil, ErrInvalidCodeLength
	}

	return &ShortenedURLCode{
		value,
	}, nil
}

func NewRandomShortenedURLCode() (*ShortenedURLCode, error) {
	value, err := generateRandomString(5)
	if err != nil {
		return nil, ErrCanNotGenerateRandomCode
	}

	return &ShortenedURLCode{
		value: value,
	}, nil
}

func generateRandomString(length int) (string, error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	randomString := make([]byte, length)
	for {
		_, err := rand.Read(randomString)
		if err != nil {
			return "", err
		}

		for i, b := range randomString {
			randomString[i] = characters[int(b)%len(characters)]
		}

		if len(randomString) == length {
			break
		}
	}

	return string(randomString), nil
}

func (v *ShortenedURLCode) Equals(other ShortenedURLCode) bool {
	return v.value == other.value
}

func (v *ShortenedURLCode) Value() string {
	return v.value
}
