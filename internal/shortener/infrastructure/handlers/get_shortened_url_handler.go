package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/shortener/application"
)

type GetShortenedURLRequest struct {
	Code string `uri:"code" binding:"required"`
}

type GetShortenedURLHandler struct {
	useCase *application.GetShortenedURLByCode
}

func NewGetShortenedURLHandler(useCase *application.GetShortenedURLByCode) *GetShortenedURLHandler {
	return &GetShortenedURLHandler{useCase}
}

func (h *GetShortenedURLHandler) Execute(ctx *gin.Context) {
	var request GetShortenedURLRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid code"})
		return
	}

	shortenedURL, err := h.useCase.Execute(request.Code)
	if err == application.ErrShortenedURLNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	originalURL := shortenedURL.URL()

	ctx.Redirect(http.StatusFound, originalURL.Value())
}
