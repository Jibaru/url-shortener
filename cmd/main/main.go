package main

import (
	"github.com/joho/godotenv"
	"log"
	"url-shortener/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	api.StartServer()
}
