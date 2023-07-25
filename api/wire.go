//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"url-shortener/config"
	"url-shortener/internal/shortener/application"
	"url-shortener/internal/shortener/infrastructure"
	"url-shortener/internal/shortener/infrastructure/handlers"
	"url-shortener/internal/shortener/infrastructure/repositories"
)

func InitializeServer() *handlers.Handlers {
	wire.Build(
		config.NewDBConfig,
		infrastructure.NewUUIDGenerator,
		infrastructure.NewDB,
		repositories.NewMySQLShortenedURLRepository,
		application.NewStoreShortenedURL,
		application.NewGetShortenedURLS,
		application.NewGetShortenedURLByCode,
		handlers.NewPostShortenedURLHandler,
		handlers.NewGetShortenedURLSHandler,
		handlers.NewGetShortenedURLHandler,
		handlers.NewHandlers,
	)

	return &handlers.Handlers{}
}
