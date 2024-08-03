package main

import (
	"github.com/joho/godotenv"
	entry "go1x_gin.template/entry"
	"os"
)

func main() {
	godotenv.Load()

	router := entry.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
