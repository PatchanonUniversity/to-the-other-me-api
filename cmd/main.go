package main

import (
	"log"
	"os"

	"to-the-other-me/internal/config"
	"to-the-other-me/internal/handler"
	"to-the-other-me/internal/repository"
	"to-the-other-me/internal/service"
	"to-the-other-me/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	db := config.InitDB()
	defer db.Close()

	letterRepository := repository.NewLetterRepository(db)
	letterService := &service.LetterService{Repository: letterRepository}
	letterHandler := &handler.LetterHandler{Service: letterService}

	geminiService := &service.GeminiService{}
	geminiHandler := &handler.GeminiHandler{Service: geminiService}

	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, geminiHandler, letterHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}