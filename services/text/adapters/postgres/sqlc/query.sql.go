// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
)

const countTexts = `-- name: CountTexts :one
SELECT count(*) FROM text
`

func (q *Queries) CountTexts(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countTexts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createText = `-- name: CreateText :one
INSERT INTO text (data) VALUES ($1) RETURNING id, data, created_at, updated_at
`

func (q *Queries) CreateText(ctx context.Context, data string) (Text, error) {
	row := q.db.QueryRow(ctx, createText, data)
	var i Text
	err := row.Scan(
		&i.ID,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRandText = `-- name: GetRandText :one
SELECT id, data, created_at, updated_at FROM text ORDER BY RANDOM() LIMIT 1
`

func (q *Queries) GetRandText(ctx context.Context) (Text, error) {
	row := q.db.QueryRow(ctx, getRandText)
	var i Text
	err := row.Scan(
		&i.ID,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
