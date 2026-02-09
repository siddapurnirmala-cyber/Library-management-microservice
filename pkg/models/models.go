package models

import "time"

type Member struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	JoinedAt time.Time `json:"joined_at"`
}

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublishedYear   int    `json:"published_year"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
}

type Borrow struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	BookID     int        `json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // Pointer for nullable time
	Status     string     `json:"status"`
}
