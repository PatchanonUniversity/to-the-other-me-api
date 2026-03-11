package routes

import (
	"to-the-other-me/internal/handler"
	"to-the-other-me/internal/middleware"

	"github.com/gin-gonic/gin"

	"os"

	"github.com/gin-contrib/cors"
)


func RegisterRoutes(
	router *gin.Engine,
	geminiHandler *handler.GeminiHandler,
	letterHandler *handler.LetterHandler,
) {
	frontendURL := os.Getenv("FRONTEND_URL")
	router.Use(cors.New(cors.Config{
	AllowOrigins: []string{
		frontendURL,
	},
	AllowMethods: []string{"GET", "POST","OPTIONS"},
	AllowHeaders: []string{"Content-Type","Authorization"},
}))
	router.Use(middleware.RateLimit())


	api := router.Group("/api/v1")

	letterRouter := api.Group("/letters")
	letterRouter.POST("", letterHandler.CreateLetter)
	letterRouter.GET("", letterHandler.GetLetters)

	flowerRouter := api.Group("/flowers")
	flowerRouter.POST("", geminiHandler.HandleGemini)


}