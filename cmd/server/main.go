package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pixprocoder/book-review-backend/internal/config"
	"github.com/pixprocoder/book-review-backend/internal/handler"
	"github.com/pixprocoder/book-review-backend/internal/middleware"
	"github.com/pixprocoder/book-review-backend/internal/repository"
	"github.com/pixprocoder/book-review-backend/internal/repository/db"
	"github.com/pixprocoder/book-review-backend/internal/service"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := repository.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("connected to database")

	queries := db.New(pool)

	bookSvc := service.NewBookService(queries)
	reviewSvc := service.NewReviewService(queries)

	bookH := handler.NewBookHandler(bookSvc)
	reviewH := handler.NewReviewHandler(reviewSvc)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.RequestID)
	r.Use(middleware.CORS(cfg.AllowedOrigins))

	r.Get("/health", handler.HealthCheck)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/books", bookH.GetAll)
		r.Get("/books/{id}", bookH.GetByID)
		r.Post("/books", bookH.Create)
		r.Put("/books/{id}", bookH.Update)
		r.Delete("/books/{id}", bookH.Delete)

		r.Post("/books/{id}/review", reviewH.Create)
		r.Get("/books/{id}/reviews", reviewH.GetByBook)
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("server starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("server stopped")
}
