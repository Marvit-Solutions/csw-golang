package main

import (
	"log"
	"os"

	"csw-golang/cmd/app"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	log.Println("Starting application...")
	r := app.StartApp()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start application:", err)
	}
}
