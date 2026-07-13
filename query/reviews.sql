-- name: CreateReview :one
INSERT INTO reviews (book_id, user_email, user_name, rating, comment)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetReviewsByBook :many
SELECT * FROM reviews
WHERE book_id = $1
ORDER BY created_at DESC;

-- name: DeleteReviewsByBook :exec
DELETE FROM reviews
WHERE book_id = $1;
