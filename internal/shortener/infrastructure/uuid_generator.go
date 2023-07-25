package infrastructure

import (
	"github.com/google/uuid"
	"url-shortener/internal/shortener/domain"
)

type UUIDGenerator struct{}

func NewUUIDGenerator() domain.IDGenerator {
	return &UUIDGenerator{}
}

func (g *UUIDGenerator) Generate() string {
	return uuid.NewString()
}
