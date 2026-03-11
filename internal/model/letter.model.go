package model

import "time"

type Letter struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	ToSentDate time.Time `json:"toSentDate"`
	Status     string    `json:"status"`
}

type SaveLetterRequest struct {
	Email   string `json:"email" binding:"required"`
	Name    string `json:"name"`
	Content string `json:"content" binding:"required"`
}