-- name: GetRandText :one
SELECT * FROM text ORDER BY RANDOM() LIMIT 1;

-- name: CreateText :one
INSERT INTO text (data) VALUES ($1) RETURNING *;

-- name: CountTexts :one
SELECT count(*) FROM text;
