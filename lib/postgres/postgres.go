package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Conn *pgx.Conn
}

func New(ctx context.Context, dbURL string) (Postgres, error) {
	const timeout = 3 * time.Second

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return Postgres{}, fmt.Errorf("failed to connect to db: %w", err)
	}

	return Postgres{conn}, nil
}

func (p Postgres) Stop(ctx context.Context) {
	p.Conn.Close(ctx)
}
