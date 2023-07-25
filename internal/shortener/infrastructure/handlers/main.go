package handlers

type Handlers struct {
	Post      *PostShortenedURLHandler
	GetAll    *GetShortenedURLSHandler
	GetByCode *GetShortenedURLHandler
}

func NewHandlers(
	Post *PostShortenedURLHandler,
	GetAll *GetShortenedURLSHandler,
	GetByCode *GetShortenedURLHandler,
) *Handlers {
	return &Handlers{Post, GetAll, GetByCode}
}
