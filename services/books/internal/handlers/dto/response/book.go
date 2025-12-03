package response

import "time"

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	ISBN          string    `json:"isbn"`
	PublishedYear int       `json:"publishedYear"`
	Genre         string    `json:"genre"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsAvailable   bool      `json:"isAvailable"`
}
