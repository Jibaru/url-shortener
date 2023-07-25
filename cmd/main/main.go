package main

import (
	"github.com/joho/godotenv"
	"log"
	"url-shortener/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	api.StartServer()
}
