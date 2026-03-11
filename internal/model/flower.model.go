package model

import "time"

type GeminiRequest struct {
	UserName        string `json:"userName"`
	UserHobby       string `json:"userHobby"`
	UserDream       string `json:"userDream"`
	UserFeeling     string `json:"userFeeling"`
	UserChasing     string `json:"userChasing"`
	UserLife        string `json:"userLife"`
	UserExpectation string `json:"userExpectation"`
}


type AIResponse struct {
	Name    string `json:"name"`
	Meaning string `json:"meaning"`
}


type Gemini struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Flower    string    `json:"flower"`
	Meaning   string    `json:"meaning"`
	CreatedAt time.Time `json:"created_at"`
}