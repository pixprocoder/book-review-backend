package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pixprocoder/book-review-backend/internal/model"
	"github.com/pixprocoder/book-review-backend/internal/response"
	"github.com/pixprocoder/book-review-backend/internal/service"
)

type BookHandler struct {
	svc *service.BookService
}

func NewBookHandler(svc *service.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := h.svc.Create(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, "book created successfully", book)
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books, err := h.svc.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "books fetched successfully", books)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	book, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "book fetched successfully", book)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.svc.Delete(r.Context(), id); err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "book deleted successfully", nil)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req model.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := h.svc.Update(r.Context(), id, &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "book updated successfully", book)
}
