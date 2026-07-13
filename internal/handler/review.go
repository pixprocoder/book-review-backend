package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pixprocoder/book-review-backend/internal/model"
	"github.com/pixprocoder/book-review-backend/internal/response"
	"github.com/pixprocoder/book-review-backend/internal/service"
)

type ReviewHandler struct {
	svc *service.ReviewService
}

func NewReviewHandler(svc *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{svc: svc}
}

func (h *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	bookID := r.PathValue("id")

	var req []model.CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if len(req) == 0 {
		response.Error(w, http.StatusBadRequest, "review data required")
		return
	}

	reviews, err := h.svc.Create(r.Context(), bookID, &req[0])
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "review posted successfully", reviews)
}

func (h *ReviewHandler) GetByBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PathValue("id")

	reviews, err := h.svc.GetByBookID(r.Context(), bookID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "reviews fetched successfully", reviews)
}
