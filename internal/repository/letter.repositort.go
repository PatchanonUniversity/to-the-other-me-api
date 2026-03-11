package repository

import (
	"context"
	"database/sql"
	"fmt"
	"to-the-other-me/internal/model"
)

type LetterRepository struct {
	DB *sql.DB
}

func NewLetterRepository(db *sql.DB) *LetterRepository {
	return &LetterRepository{DB: db}
}

func (repository *LetterRepository) CreateLetter(ctx context.Context, l *model.Letter) error {
	query := `
		INSERT INTO letters (email, name, content, to_sent_date)
		VALUES ($1, $2, $3, $4)
	`
	_, err := repository.DB.ExecContext(ctx, query, l.Email, l.Name, l.Content, l.ToSentDate)
	if err != nil {
		return fmt.Errorf("failed to create letter: %w", err)
	}
	return err
}

func (repository *LetterRepository) GetLetters(ctx context.Context) ([]model.Letter, error) {

	query := `
	SELECT id, email, name, content, to_sent_date
	FROM letters
	`

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch letters: %w", err)
	}
	defer rows.Close()

	letters := []model.Letter{}

	for rows.Next() {

		var letter model.Letter

		err := rows.Scan(
			&letter.ID,
			&letter.Email,
			&letter.Name,
			&letter.Content,
			&letter.ToSentDate,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan letter: %w", err)
		}

		letters = append(letters, letter)
	}

	return letters, nil
}

