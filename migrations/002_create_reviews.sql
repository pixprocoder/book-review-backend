-- +goose Up
CREATE TABLE reviews (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id    UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    user_email TEXT NOT NULL,
    user_name  TEXT NOT NULL DEFAULT '',
    rating     INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment    TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_reviews_book_id ON reviews(book_id);
CREATE INDEX idx_reviews_user_email ON reviews(user_email);

-- +goose Down
DROP TABLE IF EXISTS reviews;
