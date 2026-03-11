package service

import (
	"context"
	"time"
	"to-the-other-me/internal/model"
	"to-the-other-me/internal/repository"
)

type LetterService struct {
	Repository *repository.LetterRepository
}

func (service *LetterService) CreateLetter(ctx context.Context, req model.SaveLetterRequest) error {
	sendDate := time.Now().AddDate(5, 0, 0)
	letter := &model.Letter{
		Email:      req.Email,
		Name:       req.Name,
		Content:    req.Content,
		ToSentDate: sendDate,
	}

	return service.Repository.CreateLetter(ctx, letter)
}

func (service *LetterService) GetLetters(ctx context.Context) ([]model.Letter, error) {
	return service.Repository.GetLetters(ctx)
}

