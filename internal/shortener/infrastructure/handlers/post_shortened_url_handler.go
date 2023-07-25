package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/shortener/application"
	"url-shortener/internal/shortener/infrastructure/resources"
)

type PostShortenedURLRequest struct {
	URL string `json:"url"`
}

type PostShortenedURLHandler struct {
	useCase *application.StoreShortenedURL
}

func NewPostShortenedURLHandler(useCase *application.StoreShortenedURL) *PostShortenedURLHandler {
	return &PostShortenedURLHandler{
		useCase: useCase,
	}
}

func (h *PostShortenedURLHandler) Execute(ctx *gin.Context) {
	var request PostShortenedURLRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}

	shortenedURL, err := h.useCase.Execute(request.URL)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": resources.NewShortenedURLResource(shortenedURL),
	})
}
