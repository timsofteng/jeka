package text

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Repo struct {
	conn *pgx.Conn
}

func NewDB(dbURL string) (*Repo, error) {
	const timeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed connect to db: %w", err)
	}
	defer conn.Close(ctx)

	return &Repo{conn: conn}, nil
}

func (r *Repo) Rand(ctx context.Context) (string, error) {
	query := `SELECT data FROM text ORDER BY RANDOM() LIMIT 1;`
	randMsg := ""

	err := r.conn.QueryRow(ctx, query).Scan(&randMsg)
	if err != nil {
		return "", fmt.Errorf("failed to fetch text from DB: %w", err)
	}

	return randMsg, nil
}

func (r *Repo) Count(ctx context.Context) (int, error) {
	query := `SELECT count(*) FROM text`
	count := 0

	err := r.conn.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("failed to count messages in DB: %w", err)
	}

	return count, nil
}

func (r *Repo) Add(ctx context.Context, message string) error {
	query := "INSERT INTO text (data) VALUES ($1)"

	_, err := r.conn.Exec(ctx, query, message)
	if err != nil {
		return fmt.Errorf("failed to insert text to db: %w", err)
	}

	return nil
}
