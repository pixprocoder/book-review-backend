package model

import "time"

type Review struct {
	ID        string    `json:"_id"`
	BookID    string    `json:"bookId"`
	UserEmail string    `json:"userEmail"`
	UserName  string    `json:"userName"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateReviewRequest struct {
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	UserEmail string `json:"userEmail"`
	UserName  string `json:"userName"`
}
