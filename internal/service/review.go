package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pixprocoder/book-review-backend/internal/model"
	"github.com/pixprocoder/book-review-backend/internal/repository/db"
)

type ReviewService struct {
	q *db.Queries
}

func NewReviewService(q *db.Queries) *ReviewService {
	return &ReviewService{q: q}
}

func (s *ReviewService) Create(ctx context.Context, bookID string, req *model.CreateReviewRequest) (*model.Review, error) {
	bid, err := parseUUIDToPgType(bookID)
	if err != nil {
		return nil, fmt.Errorf("invalid book id: %w", err)
	}

	review, err := s.q.CreateReview(ctx, db.CreateReviewParams{
		BookID:    bid,
		UserEmail: req.UserEmail,
		UserName:  req.UserName,
		Rating:    int32(req.Rating),
		Comment:   req.Comment,
	})
	if err != nil {
		return nil, fmt.Errorf("create review: %w", err)
	}

	return toReview(&review), nil
}

func (s *ReviewService) GetByBookID(ctx context.Context, bookID string) ([]model.Review, error) {
	bid, err := parseUUIDToPgType(bookID)
	if err != nil {
		return nil, fmt.Errorf("invalid book id: %w", err)
	}

	reviews, err := s.q.GetReviewsByBook(ctx, bid)
	if err != nil {
		return nil, fmt.Errorf("get reviews: %w", err)
	}

	result := make([]model.Review, 0, len(reviews))
	for _, r := range reviews {
		result = append(result, *toReview(&r))
	}
	return result, nil
}

func toReview(r *db.Review) *model.Review {
	return &model.Review{
		ID:        pgtypeUUIDToString(r.ID),
		BookID:    pgtypeUUIDToString(r.BookID),
		UserEmail: r.UserEmail,
		UserName:  r.UserName,
		Rating:    int(r.Rating),
		Comment:   r.Comment,
		CreatedAt: r.CreatedAt.Time,
	}
}

func parseUUIDToPgType(s string) (pgtype.UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("parse uuid: %w", err)
	}
	return pgtype.UUID{Bytes: parsed, Valid: true}, nil
}

func pgtypeUUIDToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	return uuid.UUID(u.Bytes).String()
}
