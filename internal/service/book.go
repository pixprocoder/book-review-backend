package service

import (
	"context"
	"fmt"

	"github.com/pixprocoder/book-review-backend/internal/model"
	"github.com/pixprocoder/book-review-backend/internal/repository/db"
)

type BookService struct {
	q *db.Queries
}

func NewBookService(q *db.Queries) *BookService {
	return &BookService{q: q}
}

func (s *BookService) Create(ctx context.Context, req *model.CreateBookRequest) (*model.Book, error) {
	book, err := s.q.CreateBook(ctx, db.CreateBookParams{
		Title:           req.Title,
		Author:          req.Author,
		Genre:           req.Genre,
		PublicationDate: req.PublicationDate,
		ImageUrl:        req.Image,
	})
	if err != nil {
		return nil, fmt.Errorf("create book: %w", err)
	}

	return toBook(&book), nil
}

func (s *BookService) GetAll(ctx context.Context) ([]model.Book, error) {
	books, err := s.q.ListBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("list books: %w", err)
	}

	result := make([]model.Book, 0, len(books))
	for _, b := range books {
		result = append(result, *toBook(&b))
	}
	return result, nil
}

func (s *BookService) GetByID(ctx context.Context, id string) (*model.Book, error) {
	uid, err := parseUUIDToPgType(id)
	if err != nil {
		return nil, fmt.Errorf("invalid book id: %w", err)
	}

	book, err := s.q.GetBookByID(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}

	return toBook(&book), nil
}

func (s *BookService) Delete(ctx context.Context, id string) error {
	uid, err := parseUUIDToPgType(id)
	if err != nil {
		return fmt.Errorf("invalid book id: %w", err)
	}

	if err := s.q.DeleteBook(ctx, uid); err != nil {
		return fmt.Errorf("delete book: %w", err)
	}

	return nil
}

func (s *BookService) Update(ctx context.Context, id string, req *model.CreateBookRequest) (*model.Book, error) {
	uid, err := parseUUIDToPgType(id)
	if err != nil {
		return nil, fmt.Errorf("invalid book id: %w", err)
	}

	book, err := s.q.UpdateBook(ctx, db.UpdateBookParams{
		ID:              uid,
		Title:           req.Title,
		Author:          req.Author,
		Genre:           req.Genre,
		PublicationDate: req.PublicationDate,
		ImageUrl:        req.Image,
	})
	if err != nil {
		return nil, fmt.Errorf("update book: %w", err)
	}

	return toBook(&book), nil
}

func toBook(b *db.Book) *model.Book {
	return &model.Book{
			ID:              pgtypeUUIDToString(b.ID),
		Title:           b.Title,
		Author:          b.Author,
		Genre:           b.Genre,
		PublicationDate: b.PublicationDate,
		Image:           b.ImageUrl,
		CreatedAt:       b.CreatedAt.Time,
		UpdatedAt:       b.UpdatedAt.Time,
	}
}
