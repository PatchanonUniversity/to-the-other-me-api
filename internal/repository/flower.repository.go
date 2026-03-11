package repository

import (
	"database/sql"
	"to-the-other-me/internal/model"
)

type GeminiRepository struct {
	DB *sql.DB
}

func NewFlowerRepository(db *sql.DB) *GeminiRepository {
	return &GeminiRepository{DB: db}
}

func (repository *GeminiRepository) Save(letter *model.GeminiRequest) error {
	query := `INSERT INTO Geminis (user_name, flower, meaning) VALUES ($1, $2, $3)`
	_, err := repository.DB.Exec(query, letter.UserName, letter.UserDream, letter.UserChasing)
	return err
}