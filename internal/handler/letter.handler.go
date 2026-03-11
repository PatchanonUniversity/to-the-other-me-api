package handler

import (
	"to-the-other-me/internal/model"
	"to-the-other-me/internal/service"

	"github.com/gin-gonic/gin"
)

type LetterHandler struct {
	Service *service.LetterService
}

func (handler *LetterHandler) CreateLetter(ctx *gin.Context) {

	var letter model.SaveLetterRequest

	if err := ctx.ShouldBindJSON(&letter); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := handler.Service.CreateLetter(ctx.Request.Context(), letter)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "letter created successfully",
	})
}

func (handler *LetterHandler) GetLetters(ctx *gin.Context) {

	letters, err := handler.Service.GetLetters(ctx.Request.Context())

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, letters)
}

