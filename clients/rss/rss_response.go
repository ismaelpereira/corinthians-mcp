package clients

import "time"

type NewsItem struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Published   time.Time `json:"published"`
}
