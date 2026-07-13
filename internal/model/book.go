package model

import "time"

type Book struct {
	ID              string    `json:"_id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Genre           string    `json:"genre"`
	PublicationDate string    `json:"publicationDate"`
	Image           string    `json:"image"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateBookRequest struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	PublicationDate string `json:"publicationDate"`
	Image           string `json:"image"`
}
