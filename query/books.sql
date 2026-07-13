-- name: CreateBook :one
INSERT INTO books (title, author, genre, publication_date, image_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetBookByID :one
SELECT * FROM books
WHERE id = $1
LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY created_at DESC;

-- name: UpdateBook :one
UPDATE books
SET title = $2,
    author = $3,
    genre = $4,
    publication_date = $5,
    image_url = $6,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
