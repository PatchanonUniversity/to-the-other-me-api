package routes

import (
	"net/http"

	"to-the-other-me/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	geminiHandler *handler.GeminiHandler,
	letterHandler *handler.LetterHandler,
) {

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	api := router.Group("/api/v1")

	letterRouter := api.Group("/letters")
	letterRouter.POST("", letterHandler.CreateLetter)
	letterRouter.GET("", letterHandler.GetLetters)

	flowerRouter := api.Group("/flowers")
	flowerRouter.POST("", geminiHandler.HandleGemini)


}