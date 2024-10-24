package postgres

import (
	"context"
	"fmt"

	"github.com/timsofteng/jeka/services/text/adapters/postgres/sqlc"
	"github.com/timsofteng/jeka/services/text/entities"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn    *pgx.Conn
	queries *sqlc.Queries
}

func New(db *pgx.Conn) *Postgres {
	q := sqlc.New(db)

	return &Postgres{conn: db, queries: q}
}

func (r *Postgres) Rand(ctx context.Context) (entities.RandText, error) {
	textRow, err := r.queries.GetRandText(ctx)
	if err != nil {
		return entities.RandText{},
			fmt.Errorf("failed to fetch text from DB: %w", err)
	}

	text, err := entities.NewRandText(textRow.Data)
	if err != nil {
		return entities.RandText{},
			fmt.Errorf("failed to create new rand text: %w", err)
	}

	return text, nil
}

func (r *Postgres) Count(ctx context.Context) (uint, error) {
	count, err := r.queries.CountTexts(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to count messages in DB: %w", err)
	}

	return uint(count), nil
}

func (r *Postgres) Add(ctx context.Context, message string) error {
	_, err := r.queries.CreateText(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to insert text to db: %w", err)
	}

	return nil
}
