package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	handlers := InitializeServer()

	router := gin.Default()
	router.GET("/r/:code", handlers.GetByCode.Execute)

	api := router.Group("/api")
	api.POST("/shortened-urls", handlers.Post.Execute)
	api.GET("/shortened-urls", handlers.GetAll.Execute)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
