package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/shortener/application"
	"url-shortener/internal/shortener/infrastructure/resources"
)

type GetShortenedURLSHandler struct {
	useCase *application.GetShortenedURLS
}

func NewGetShortenedURLSHandler(useCase *application.GetShortenedURLS) *GetShortenedURLSHandler {
	return &GetShortenedURLSHandler{useCase}
}

func (h *GetShortenedURLSHandler) Execute(ctx *gin.Context) {
	shortenedURLS := h.useCase.Execute()

	ctx.JSON(http.StatusOK, gin.H{
		"data": resources.NewShortenedURLCollection(shortenedURLS),
	})
}
