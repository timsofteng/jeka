package postgres

import (
	"context"
	"fmt"
	"log"
	"telegraminput/services/voice/entities"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn *pgx.Conn
}

func New(db *pgx.Conn) *Postgres {
	return &Postgres{conn: db}
}

func (r *Postgres) Rand(ctx context.Context) (entities.RandVoice, error) {
	voiceID := ""
	query := `SELECT id FROM voice ORDER BY RANDOM() LIMIT 1;`

	err := r.conn.QueryRow(ctx, query).Scan(&voiceID)
	if err != nil {
		return entities.RandVoice{},
			fmt.Errorf("failed to get voice id from db: %w", err)
	}

	randVoice, err := entities.NewRandVoice(voiceID)
	if err != nil {
		return entities.RandVoice{},
			fmt.Errorf("failed to create new voice entity: %w", err)
	}

	return randVoice, nil
}

func (r *Postgres) Count(ctx context.Context) (uint, error) {
	var count uint

	query := `SELECT count(*) FROM voice`

	err := r.conn.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("filed to get voices count from db: %w", err)
	}

	return count, nil
}

func (r *Postgres) Add(ctx context.Context, voiceID string) error {
	query := "INSERT INTO voice (id) VALUES ($1)"

	_, err := r.conn.Exec(ctx, query, voiceID)
	if err != nil {
		return fmt.Errorf("failed to add voice to db: %w", err)
	}

	log.Print("voice_id added to database: ", voiceID)

	return nil
}
