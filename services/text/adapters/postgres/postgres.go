package postgres

import (
	"context"
	"fmt"
	"telegraminput/services/text/entities"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn *pgx.Conn
}

func New(db *pgx.Conn) *Postgres {
	return &Postgres{conn: db}
}

func (r *Postgres) Rand(ctx context.Context) (entities.RandText, error) {
	query := `SELECT data FROM text ORDER BY RANDOM() LIMIT 1;`
	randMsg := ""

	err := r.conn.QueryRow(ctx, query).Scan(&randMsg)
	if err != nil {
		return entities.RandText{},
			fmt.Errorf("failed to fetch text from DB: %w", err)
	}

	text, err := entities.NewRandText(randMsg)
	if err != nil {
		return entities.RandText{},
			fmt.Errorf("failed to create new rand text: %w", err)
	}

	return text, nil
}

func (r *Postgres) Count(ctx context.Context) (uint, error) {
	var count uint

	query := `SELECT count(*) FROM text`

	err := r.conn.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("failed to count messages in DB: %w", err)
	}

	return count, nil
}

func (r *Postgres) Add(ctx context.Context, message string) error {
	query := "INSERT INTO text (data) VALUES ($1)"

	_, err := r.conn.Exec(ctx, query, message)
	if err != nil {
		return fmt.Errorf("failed to insert text to db: %w", err)
	}

	return nil
}
