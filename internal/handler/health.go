package handler

import (
	"net/http"

	"github.com/pixprocoder/book-review-backend/internal/response"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "server is running", map[string]string{
		"status": "healthy",
	})
}
